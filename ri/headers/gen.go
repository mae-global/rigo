/* RtToken generator */
package main

/* This tool will read ri.h and generate all the known tokens */

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"
)

const (
	Version = 0
)

func main() {
	target := flag.String("target", "/opt/pixar/RenderManProServer-20.8/include/ri.h", "ri.h to generate from")
	pname := flag.String("package", "ri", "package name")

	flag.Parse()

	fmt.Printf("RtToken Generator Version %d\n", Version)
	fmt.Printf("Parsing [%s]...\n", *target)

	file, err := ioutil.ReadFile(*target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading [%s] -- %v\n", *target, err)
		os.Exit(1)
	}

	lines := strings.Split(string(file), "\n")

	ri := &Ri{}
	ri.Tokens = make([]string, 0)

	hashes := make(map[string][]int, 0)
	hashesStarted := false

	/* add 'version' */
	hashes["version"] = Hash("version")

	/* very crude parsing here */
	for i, line := range lines {
		if strings.Contains(line, "#define RI_VERSION") {
			parts := strings.Split(line, " ")
			if len(parts) == 3 {
				ver, err := strconv.Atoi(parts[2])
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error parsing RI_VERSION -- %v\n", err)
					os.Exit(1)
				}
				ri.Version = ver
				fmt.Printf("[%05d] RI_VERSION = %d\n", i, ri.Version)
			}
			continue
		} /* RI_VERSION */

		if strings.Contains(line, "RI_EXPORT RtToken") {
			parts := strings.Split(line, " ")
			if len(parts) == 3 {
				token := strings.TrimPrefix(strings.Replace(parts[2], ";", "", 1), "RI_")
				ri.Tokens = append(ri.Tokens, token)
				if len(token) > ri.longest {
					ri.longest = len(token)
				}
			}
		}
		
		if strings.Contains(line, "RI_EXPORT") {
				//	fmt.Printf("[%05d] parsing %s\n",i,line)

				parts := strings.Split(line, " ")
				/* ignore first 2 */
				name := ""
				for _, p := range parts {
					if strings.Contains(p, "Ri") {
						name = p
						break
					}
				}
				if name == "" {
					continue
				}

				nameparts := strings.Split(name, "(")
				name = strings.TrimPrefix(nameparts[0], "Ri")
				name = strings.Replace(name, ";", "", -1)

				if name == "ArchiveBegin" { /* skip until we find the the actual Ri C API */
					hashesStarted = true
				}

				if hashesStarted {

					m := Hash(name)
					hashes[name] = m

					fmt.Printf("[%05d] Func = [%s] bloom@%v\n", i, name, m)
				}
			}
		}
	

	/* remove all the ...V functions */
	toremove := make([]string, 0)

	for key, _ := range hashes {
		if key[len(key)-1] == 'V' {
			toremove = append(toremove, key)
		}
	}

	for _, key := range toremove {
		delete(hashes, key)
	}

	/* find the best fit for the bloom filter */
	size := 128
	var bloom []int
	var filter *BloomFilter

	for {
		bloom = make([]int, size)

		for _, h := range hashes {

			for i := 0; i < len(h); i++ {
				bloom[h[i]%size]++
			}
		}

		filter = &BloomFilter{bloom, len(hashes)}
		fmt.Printf("\n============== [%05d] BloomFilter\n%s\n", size, filter.Print())

		_, _, sparse := filter.Stats()

		if !filter.IsMember("Alice", "Fred", "Eve") && filter.IsMember("Sphere", "Translate") {
			if sparse >= 0.8 {
				break
			}
		} else {
			fmt.Printf(">>> failed membership test\n")
		}

		size = size * 2
	}

	keys, bits := filter.Raw()
	ri.FilterData = bits
	ri.FilterKeys = keys
	ri.ArgumentsData = ArgsIndex

	ri.FilterKeysData = make([]string, 0)
	for key, _ := range hashes {
		ri.FilterKeysData = append(ri.FilterKeysData, key)
	}

	sort.Sort(ByAlpha(ri.FilterKeysData))

	if len(ri.FilterKeysData) != len(ArgsIndex) {
		fmt.Fprintf(os.Stderr, "Mismatch keys (%d) with args index (%d)\n", len(ri.FilterKeysData), len(ArgsIndex))
		os.Exit(1)
	}

	for i := 0; i < len(ri.FilterKeysData); i++ {
		fmt.Printf("[%s], %d args\n", ri.FilterKeysData[i], ArgsIndex[i])
	}

	fmt.Printf("\n\nParsed %d lines\n", len(lines))

	o := &T{}
	o.Version = Version
	o.Date = time.Now().UTC()
	o.Source = *target
	o.Name = *pname
	o.Ri = ri

	t, err := template.New("temp").Funcs(template.FuncMap{"Token": ri.Token, "Value": ri.Value}).Parse(Template)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing template -- %v\n", err)
		os.Exit(1)
	}

	buf := bytes.NewBuffer(nil)

	if err := t.Execute(buf, o); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing template -- %v\n", err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile("out.txt", buf.Bytes(), 0664); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file -- %v\n", err)
		os.Exit(1)
	}

}

type ByAlpha []string

func (b ByAlpha) Len() int           { return len(b) }
func (b ByAlpha) Less(i, j int) bool { return (b[i] < b[j]) }
func (b ByAlpha) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type T struct {
	Version int
	Date    time.Time
	Source  string
	Name    string
	Ri      *Ri
}

type Ri struct {
	Version        int
	Tokens         []string
	longest        int
	FilterData     []int
	FilterKeys     int
	FilterKeysData []string
	ArgumentsData  []int
}

func (ri *Ri) Token(t string) string {
	//return fmt.Sprintf("\t%s",strings.Title(strings.ToLower(t)))
	return fmt.Sprintf("\t%s", t)
}

func (ri *Ri) Value(v string) string {
	return strings.ToLower(v)
}

var (
	/* maps to each (normal) Ri call the number of arguments it should have */
	ArgsIndex = []int{
		1, 0, 2, 1, 1, 1, 0, 0, 4, 1, 7, 1, 2, 1, 2, 6, 1, 3,
		1, 3, 1, 1, 1, 4, 4, 4, 2, 1, 3, 1, 4, 3, 1, 4, 3, 1,
		1, 0, 1, 0, 1, 0, 0, 1, 3, 0, 1, 2, 1, 3, 1, 1, 0, 2,
		2, 1, 0, 1, 10, 3, 0, 1, 0, 2, 1, 2, 1, 2, 1, 3, 7,
		5, 2, 7, 1, 1, 0, 10, 0, 0, 1, 1, 1, 1, 4, 1, 5, 2,
		1, 3, 1, 2, 1, 1, 4, 3, 1, 4, 2, 1, 5, 2, 1, 2, 0, 0,
		0, 4, 3, 1, 4, 2, 1, 1, 2, 1, 7, 1, 0, 4, 9, 1, 1, 8,
		5, 0, 0, 4, 3, 10, 2, 1, 1, 1, 3, 2, 0, 0, 0, 1, 
	}
	/* TODO: currently this is manually generated, should replace */
)

const Template = `/* machine generated 
 * build tool version {{.Version}}
 * generated on {{.Date }}
 * source {{.Source}}
 * RiVersion {{.Ri.Version}} 	
 */

package {{.Name}}

const Version RtFloat = 3.04

const (
{{range $token := .Ri.Tokens}}{{$token | Token}} RtToken = "{{$token | Value}}"
{{end}}
)

var (
	bloomFilterData = []int{ {{range $i := .Ri.FilterData}}{{$i}},{{end}} } 
	bloomFilterKeysData = []string{ {{range $i := .Ri.FilterKeysData}}"{{$i}}",{{end}} }
	RiArgumentsData = []int{ {{range $i := .Ri.ArgumentsData}}{{$i}},{{end}} }
)

const (
	bloomFilterKeys int = {{.Ri.FilterKeys}}
)

func RiBloomFilter() *BloomFilter {
	bits := make([]int,len(bloomFilterData))
	copy(bits,bloomFilterData)
	return &BloomFilter{bits,bloomFilterKeys}
}



/* EOF */
`

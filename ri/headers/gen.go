/* RtToken generator */
package main

/* This tool will read ri.h and generate all the known tokens */

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
	"os"
	"text/template"
	"flag"
	"time"
	"bytes"
)

const (
	Version = 0
)

func main() {
	target := flag.String("target","/opt/pixar/RenderManProServer-20.8/include/ri.h","ri.h to generate from")
	pname := flag.String("package","ri","package name")
	

	flag.Parse()

	fmt.Printf("RtToken Generator Version %d\n",Version)
	fmt.Printf("Parsing [%s]...\n",*target)
	
	file,err := ioutil.ReadFile(*target)
	if err != nil {
		fmt.Fprintf(os.Stderr,"Error reading [%s] -- %v\n",*target,err)
		os.Exit(1)
	}

	lines := strings.Split(string(file),"\n")

	ri := &Ri{}
	ri.Tokens = make([]string,0)

	/* very crude parsing here */
	for i,line := range lines {
		if strings.Contains(line,"#define RI_VERSION") {
			parts := strings.Split(line," ")
			if len(parts) == 3 {
				ver,err := strconv.Atoi(parts[2])
				if err != nil {
					fmt.Fprintf(os.Stderr,"Error parsing RI_VERSION -- %v\n",err)
					os.Exit(1)
				}
				ri.Version = ver
				fmt.Printf("[%05d] RI_VERSION = %d\n",i,ri.Version)
			}
			continue		
		} /* RI_VERSION */

		if strings.Contains(line,"RI_EXPORT RtToken") {
			parts := strings.Split(line," ")
			if len(parts) == 3 {
				token := strings.TrimPrefix(strings.Replace(parts[2],";","",1),"RI_")
				ri.Tokens = append(ri.Tokens,token)
				if len(token) > ri.longest {
					ri.longest = len(token)
				}
			}
		}
	}

	fmt.Printf("Parsed %d lines\n",len(lines))

	o := &T{}
	o.Version = Version
	o.Date = time.Now().UTC()
	o.Source = *target
	o.Name = *pname
	o.Ri = ri


	t,err := template.New("temp").Funcs(template.FuncMap{"Token":ri.Token,"Value":ri.Value}).Parse(Template)
	if err != nil {
		fmt.Fprintf(os.Stderr,"Error parsing template -- %v\n",err)
		os.Exit(1)
	}

	buf := bytes.NewBuffer(nil)

	if err := t.Execute(buf,o); err != nil {
		fmt.Fprintf(os.Stderr,"Error executing template -- %v\n",err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile("out.txt",buf.Bytes(),0664); err != nil {
		fmt.Fprintf(os.Stderr,"Error writing file -- %v\n",err)
		os.Exit(1)
	}


}

type T struct {
	Version int
	Date time.Time
	Source string
	Name string
	Ri *Ri
}

type Ri struct {
	Version int
	Tokens []string
	longest int
}

func (ri *Ri) Token(t string) string {
	//return fmt.Sprintf("\t%s",strings.Title(strings.ToLower(t)))
	return fmt.Sprintf("\t%s",t)
}

func (ri *Ri) Value(v string) string {
	return strings.ToLower(v)
}

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

/* EOF */
`



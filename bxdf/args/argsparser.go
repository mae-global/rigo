/* An bxdf xml args parser, you should be able to 
 * parse all the attributes of an bxdf. */

package args

import (
	"encoding/xml"
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
	"os"
	"path"

	. "github.com/mae-global/rigo/ri"
	. "github.com/mae-global/rigo/bxdf"
)

const (
	DefaultArgsSubPath = "/lib/RIS/bxdf/Args/"
)

type ArgsTag struct {
	XMLName xml.Name `xml:"tag"`
	Value string `xml:"value,attr"`
}


type ArgsShaderType struct {
	XMLName xml.Name `xml:"shaderType"`
	Tag ArgsTag `xml:"tag"`
}

type ArgsTags struct {
	XMLName xml.Name `xml:"tags"`
	Tags []ArgsTag `xml:"tag"`
}

type ArgsHelp struct {
	XMLName xml.Name `xml:"help"`
	Value string `xml:",innerxml"`
	
}

type ArgsString struct {
	XMLName xml.Name `xml:"string"`

	Name string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type ArgsHintDict struct {
	XMLName xml.Name `xml:"hintdict"`
	Attributes []ArgsString `xml:"string"`
	
	Name string `xml:"name,attr"`
}

type ArgsParam struct {
	XMLName xml.Name `xml:"param"`
	Tags ArgsTags `xml:"tags"`
	Help ArgsHelp `xml:"help"`

	Label string `xml:"label,attr"`
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
	Default string `xml:"default,attr"`
	Min string `xml:"min,attr"`
	Max string `xml:"max,attr"`
	Widget string `xml:"widget,attr"`
	Connectable string `xml:"connectable,attr"`
}

type ArgsRfmdata struct {
	XMLName xml.Name `xml:"rfmdata"`
	
	NodeId string `xml:"nodeid,attr"`
	Classification string `xml:"classification,attr"`
}

type Args struct {
	XMLName xml.Name `xml:"args"`
	Shader ArgsShaderType `xml:"shaderType"`
	Params []ArgsParam `xml:"param"`
	Rfmdata ArgsRfmdata `xml:"rfmdata"`

	Format string `xml:"format,attr"`
}



/* This is the test parser */
func ParseArgsXML(data []byte) (*Args,error) {

	var a Args 
	if err := xml.Unmarshal(data,&a); err != nil {
		return nil,err
	}
	return &a,nil
}

func Parse(name string,data []byte) (Bxdfer,error) {

	args,err := ParseArgsXML(data)
	if err != nil {
		return nil,err
	}

	general := NewGeneralBxdf(RtToken(name),RtToken(args.Rfmdata.NodeId),RtString(args.Rfmdata.Classification))
	
	for _,param := range args.Params {

		var def Rter
		var min Rter
		var max Rter
		switch param.Type {
			case "float":
				def = RtFloat(0.0)
				min = RtFloat(0.0)
				max = RtFloat(0.0)				

				if len(param.Default) > 0 {
					if f,err := strconv.ParseFloat(param.Default,64);  err != nil {
						return nil,err
					}	else {
						def = RtFloat(f)
					}
				}
				if len(param.Min) > 0 {
					if f,err := strconv.ParseFloat(param.Min,64); err != nil {
						return nil,err
					} else {
						min = RtFloat(f)
					}
				}
				if len(param.Max) > 0 {
					if f,err := strconv.ParseFloat(param.Max,64); err != nil {
						return nil,err
					} else {
						max = RtFloat(f)
					}
				}
			break
			case "int":
				def = RtInt(0)
				min = RtInt(0)
				max = RtInt(0)

				if len(param.Default) > 0 {
					if i,err := strconv.Atoi(param.Default); err != nil {
						return nil,err
					} else {
						def = RtInt(i)
					}
				}
				if len(param.Min) > 0 {
					if i,err := strconv.Atoi(param.Min); err != nil { 
						return nil,err
					} else {
						min = RtInt(i)
					}
				}
				if len(param.Max) > 0 {
					if i,err := strconv.Atoi(param.Max); err != nil {
						return nil,err
					} else {
						max = RtInt(i)
					}
				}

			break
			case "color":
				def = RtColor{0,0,0}
				min = RtColor{0,0,0}
				max = RtColor{0,0,0}

				if len(param.Default) > 0 {
					def = Str2Color(param.Default)
				}
				if len(param.Min) > 0 {
					min = Str2Color(param.Min)
				}
				if len(param.Max) > 0 {
					max = Str2Color(param.Max)
				}				

			break
			case "normal":
				def = RtNormal{0,0,0}
				min = RtNormal{0,0,0}
				max = RtNormal{0,0,0}

				if len(param.Default) > 0 {
					def = Str2Normal(param.Default)
				}
				if len(param.Min) > 0 {
					min = Str2Normal(param.Min)
				}
				if len(param.Max) > 0 {
					max = Str2Normal(param.Max)
				}
			break		
			case "vector":
				def = RtVector{0,0,0}
				min = RtVector{0,0,0}
				max = RtVector{0,0,0}

				if len(param.Default) > 0 {
					def = Str2Vector(param.Default)
				}
				if len(param.Min) > 0 {
					min = Str2Vector(param.Min)
				}
				if len(param.Max) > 0 {
					max = Str2Vector(param.Max)
				}
			break
			case "string":
				def = RtString(param.Default)
				min = RtString(param.Min)
				max = RtString(param.Max)
			break
			case "struct": /* FIXME, don't know how to handle this type ? */
				continue
			break	
			default:
				return nil,fmt.Errorf("Unknown Type %s=[%s]",param.Name,param.Type)
			break
		}

		param := &Param{Label:RtString(param.Label),
													Name:RtToken(param.Name),
													Type:RtToken(param.Type),
													Widget:RtToken(param.Widget),
													Help:RtString(strings.TrimSpace(param.Help.Value)),
													Default:def,Min:min,Max:max,Value:def}

		if err := general.AddParam(param); err != nil {
			return nil,err
		}
	}

	
	return general,nil
}

/* ParseFile */
func ParseFile(root,name string) (Bxdfer,error) {
	
	if root == "" {
		root = "RMANTREE"
	}	
	
	rmantree := os.Getenv(root)
	if len(rmantree) == 0 {
		return nil,fmt.Errorf("is \"%s\" set?",root)
	}

	file,err := ioutil.ReadFile(rmantree + DefaultArgsSubPath + name + ".args")
	if err != nil {
		return nil,err
	}

	return Parse(name,file)
}

/* ParseFileAbs */
func ParseFileAbs(filepath string) (Bxdfer,error) {

	ext := path.Ext(filepath)
	name := strings.TrimSuffix(filepath,ext)

	content,err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil,err
	}

	return Parse(name,content)
}

/* ParseFiles */
func ParseFiles(root string,names... string) ([]Bxdfer,error) {

	if root == "" {
		root = "RMANTREE"
	}

	rmantree := os.Getenv(root)
	if len(rmantree) == 0 {
		return nil,fmt.Errorf("is \"%s\" set?",root)
	}

	out := make([]Bxdfer,0)

	for _,name := range names {
		file,err := ioutil.ReadFile(rmantree + DefaultArgsSubPath + name + ".args")
		if err != nil {
			return nil,err
		}

		b,err := Parse(name,file)
		if err != nil {
			return nil,err
		}
		out = append(out,b)
	}
	return out,nil
}

/* ParseDir */
func ParseDir(root string) ([]Bxdfer,error) {
	
	if root == "" {
		root = "RMANTREE"
	}

	rmantree := os.Getenv(root)
	if len(rmantree) == 0 {
		return nil,fmt.Errorf("is \"%s\" set?",root)
	}

	from := rmantree + DefaultArgsSubPath

	files, err := ioutil.ReadDir(from)
	if err != nil {
		return nil,err
	}

	list := make([]Bxdfer,0)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := path.Ext(file.Name())
		if ext != ".args" {
			continue
		}

		name := strings.TrimSuffix(file.Name(),ext)

		content,err := ioutil.ReadFile(from + file.Name())
		if err != nil {
			return nil,err
		}

		bxdf,err := Parse(name,content)
		if err != nil {
			return nil,err
		}			

		list = append(list,bxdf)		
	}
	
	return list,nil
}






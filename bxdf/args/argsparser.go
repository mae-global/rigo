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

	. "github.com/mae-global/rigo/ri"
	. "github.com/mae-global/rigo/bxdf"
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
func ParseArgs(data []byte) (*Args,error) {

	var a Args 
	if err := xml.Unmarshal(data,&a); err != nil {
		return nil,err
	}
	return &a,nil
}

func Parse(name string,data []byte) (Bxdfer,error) {

	args,err := ParseArgs(data)
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

func ParseArgsFile(name string) (Bxdfer,error) {

	rmantree := os.Getenv("RMANTREE")
	if len(rmantree) == 0 {
		return nil,fmt.Errorf("is RMANTREE set?")
	}

	file,err := ioutil.ReadFile(rmantree + "/lib/RIS/bxdf/Args/" + name + ".args")
	if err != nil {
		return nil,err
	}

	return Parse(name,file)
}








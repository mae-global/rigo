/* An bxdf xml args parser, you should be able to 
 * parse all the attributes of an bxdf. */

package args

import (
	"encoding/xml"
	"fmt"
	"strings"

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

	general := &GeneralBxdf{}
	general.name = RtToken(name)
	general.nodeid = RtToken(args.Rfmdata.NodeId)
	general.classification = RtString(args.Rfmdata.Classification)
	general.params = make([]Param,len(args.Params))
	general.values = make(map[RtToken]Rter,0)

	for i,param := range args.Params {

		var def Rter
		var min Rter
		var max Rter
		switch param.Type {
			case "float":
				def = RtFloat(1.0)
				min = RtFloat(0.0)
				max = RtFloat(0.0)
			break
			case "int":
				def = RtInt(1)
				min = RtInt(0)
				max = RtInt(0)
			break
			case "color":
				def = RtColor{1,1,1}
				min = RtColor{0,0,0}
				max = RtColor{1,1,1}
			break
			default:
				return nil,fmt.Errorf("Unknown Type %s=[%s]",param.Name,param.Type)
			break
		}


		general.params[i] = Param{Label:RtString(param.Label),
															Name:RtToken(param.Name),
															Type:RtToken(param.Type),
															Widget:RtToken(param.Widget),
															Help:RtString(strings.TrimSpace(param.Help.Value)),
															Default:def,Min:min,Max:max}

		general.values[RtToken(param.Name)] = def
	}

	
	return general,nil
}






/* An bxdf xml args parser, you should be able to 
 * parse all the attributes of an bxdf. */

package args

import (
	"encoding/xml"
)

type Tag struct {
	XMLName xml.Name `xml:"tag"`
	Value string `xml:"value,attr"`
}


type ShaderType struct {
	XMLName xml.Name `xml:"shaderType"`
	Tag Tag `xml:"tag"`
}

type Tags struct {
	XMLName xml.Name `xml:"tags"`
	Tags []Tag `xml:"tag"`
}

type Help struct {
	XMLName xml.Name `xml:"help"`
}

type String struct {
	XMLName xml.Name `xml:"string"`

	Name string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type HintDict struct {
	XMLName xml.Name `xml:"hintdict"`
	Attributes []String `xml:"string"`
	
	Name string `xml:"name,attr"`
}

type Param struct {
	XMLName xml.Name `xml:"param"`
	Tags Tags `xml:"tags"`
	Help Help `xml:"help"`

	Label string `xml:"label,attr"`
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
	Default string `xml:"default,attr"`
	Widget string `xml:"widget,attr"`
	Connectable string `xml:"connectable,attr"`
}

type Rfmdata struct {
	XMLName xml.Name `xml:"rfmdata"`
	
	NodeId string `xml:"nodeid,attr"`
	Classification string `xml:"classification,attr"`
}

type Args struct {
	XMLName xml.Name `xml:"args"`
	Shader ShaderType `xml:"shaderType"`
	Params []Param `xml:"param"`
	Rfmdata Rfmdata `xml:"rfmdata"`

	Format string `xml:"format,attr"`
}

/* This is the test parser */
func Parse(data []byte) (*Args,error) {

	var a Args 
	if err := xml.Unmarshal(data,&a); err != nil {
		return nil,err
	}
	return &a,nil
}





package ris

import (
	. "github.com/mae-global/rigo/ri"
	. "github.com/mae-global/rigo/ri/handles"
	"fmt"
)
/* https://renderman.pixar.com/resources/current/RenderMan/devExamples.html */

var (
	ErrInvalidType = fmt.Errorf("invalid type")
)


type Widget interface {
	Name() RtToken
	NameSpec() RtToken
	Label() RtString
	SetValue(value Rter) error
	GetValue() Rter
	Help() RtString
	Bounds() (Rter,Rter) /* min and max if set */
	Default() error

	Next() Widget
	Prev() Widget
}

type Shader interface {
	/* name, shaderhandle,args,params,values */
	Write() (RtName,RtShaderHandle,[]Rter,[]Rter,[]Rter)
	ShaderType() RtName
	Name() RtToken
	NodeId() RtToken
	Classifiation() RtString
	
	Widget(name RtToken) Widget 
	FirstWidget() Widget
	LastWidget() Widget

	Names() []RtToken
	NamesSpec() []RtToken
	Handle() RtShaderHandle

	SetValue(name RtToken,value Rter) error 
	Value(name RtToken) Rter	
}







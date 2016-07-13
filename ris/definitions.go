package ris

import (
	"fmt"
	. "github.com/mae-global/rigo/ri"
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
	Bounds() (Rter, Rter) /* min and max if set */
	Default() error

	Next() Widget
	Prev() Widget
}

type Shader interface {
	/* name, shaderhandle,args,params,values */
	Write() (RtName, RtShaderHandle, []Rter, []Rter, []Rter)
	ShaderType() RtName
	Name() RtToken
	NodeId() RtToken
	Classifiation() RtString
	Info() (RtInt,RtInt) /* param, output counts */

	Widget(name RtToken) Widget
	FirstWidget() Widget
	LastWidget() Widget

	Names() []RtToken
	NamesSpec() []RtToken
	Handle() RtShaderHandle

	SetValue(name RtToken, value Rter) error
	SetReferencedValue(name RtToken, value RtString) error
	
	Value(name RtToken) Rter
	ReferenceOutput(name RtToken) RtString
}

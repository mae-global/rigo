package bxdf

import (
	. "github.com/mae-global/rigo/ri"
)
/* https://renderman.pixar.com/resources/current/RenderMan/devExamples.html */

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

type Bxdfer interface {
	Write() (RtName,[]Rter,[]Rter)
	Name() RtToken
	NodeId() RtToken
	Classifiation() RtString
	
	Widget(name RtToken) Widget 
	FirstWidget() Widget
	LastWidget() Widget

	Names() []RtToken
	NamesSpec() []RtToken

	SetValue(name RtToken,value Rter) error 
	Value(name RtToken) Rter	
}







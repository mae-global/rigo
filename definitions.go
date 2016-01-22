/* rigo/definitions.go */
package ri

import (
	"fmt"
)

type Rter interface { /* TODO: add String() */
	Serialise() string
}

type RtName string

func (s RtName) Serialise() string {
	return string(s)
}

type RtBoolean bool
type RtInt int

func (i RtInt) Serialise() string {
	return fmt.Sprintf("%d",int(i))
}

type RtIntArray []RtInt

func (a RtIntArray) Serialise() string {
	out := ""
	for i,r := range a {
		out += fmt.Sprintf("%d",int(r))
		if i < (len(a) - 1) {
			out += " "
		}
	}
	return fmt.Sprintf("[%s]",out)
}


type RtFloat float64

func (f RtFloat) String() string {
	return f.Serialise()
}

func (f RtFloat) Serialise() string {
	return fmt.Sprintf("%f",float64(f))
}

type RtFloatArray []RtFloat

func (a RtFloatArray) Serialise() string {
	out := ""
	for i,r := range a {
		out += fmt.Sprintf("%f",float64(r))
		if i < (len(a) - 1) {
			out += " "
		}
	}
	return fmt.Sprintf("[%s]",out)
}

type RtToken string

func (s RtToken) String() string {
	return s.Serialise()
}

func (s RtToken) Serialise() string {
	return fmt.Sprintf("\"%s\"",string(s))
}

type RtColor []RtFloat

func (c RtColor) String() string {
	return c.Serialise()
}

func (c RtColor) Serialise() string {
	out := ""
	for i,r := range c {
		out += fmt.Sprintf("%f",float64(r))
		if i < (len(c) - 1) {
			out += " "
		}
	}
	return fmt.Sprintf("[%s]",out)
}


type RtPoint [3]RtFloat
type RtVector [3]RtFloat
type RtNormal [3]RtFloat
type RtHpoint [4]RtFloat
type RtMatrix [4][4]RtFloat
type RtBasis  [4][4]RtFloat
type RtBound  [6]RtFloat
type RtString string

func (s RtString) Serialise() string {
	return fmt.Sprintf("\"%s\"",string(s))
}


type RtLightHandle uint

func (l RtLightHandle) String() string {
	return l.Serialise()
}

func (l RtLightHandle) Serialise() string {
	return fmt.Sprintf("%d",uint(l))
} 

const (
	Perspective RtToken = "perspective"
	Orthographic RtToken = "orthographic"
)

const (
	BoxFilter RtToken = "box"
	TriangleFilter RtToken = "triangle"
	CatmullRomFilter RtToken = "catmull-rom"
	GaussianFilter RtToken = "gaussian"
	SincFilter RtToken = "sinc"

	RGBA RtToken = "RGBA"
)

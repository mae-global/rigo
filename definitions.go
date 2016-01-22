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

func (s RtBoolean) String() string {
	return s.Serialise()
}

func (s RtBoolean) Serialise() string {
	if bool(s) {
		return "1"
	}
	return "0"
}

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

func (p RtPoint) Serialise() string {
	return fmt.Sprintf("%f %f %f",p[0],p[1],p[2])
}

type RtPointArray []RtPoint

func (p RtPointArray) Serialise() string {
	out := ""
	for i := 0; i < len(p); i++ {
		out += p[i].Serialise()
		if i < len(p) - 1 {
			out += " "
		}
	}
	return fmt.Sprintf("[%s]",out)
}

type RtVector [3]RtFloat
type RtNormal [3]RtFloat
type RtHpoint [4]RtFloat
type RtMatrix [16]RtFloat

func (m RtMatrix) Serialise() string {
	out := ""
	for i := 0; i < 16; i++ {
		out += fmt.Sprintf("%f",m[i])
		if i < 15 {
			out += " "
		}
	}
	
	return fmt.Sprintf("[%s]",out)
}

type RtBasis  [16]RtFloat

func (b RtBasis) Serialise() string {
	out := ""
	for i := 0; i < 16; i++ {
		out += fmt.Sprintf("%f",b[i])
		if i < 15 {
			out += " "
		}	
	}
	return fmt.Sprintf("[%s]",out)
}



type RtBound  [6]RtFloat

func (b RtBound) Serialise() string {
	return fmt.Sprintf("[%f %f %f %f %f %f]",b[0],b[1],b[2],b[3],b[4],b[5])
}


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
	Bilinear RtToken = "bilinear"
	Bicubic RtToken = "bicubic"

	RGBA RtToken = "RGBA"
	P RtToken = "P"
	Pz RtToken = "Pz"
	Pw RtToken = "Pw"
	N RtToken = "N"
	Cs RtToken = "Cs"
	Os RtToken = "Os"
	
)

/* rigo/definitions.go */
package ri

import (
	"fmt"
)

type Rter interface { 
	String() string
	Serialise() string
}

/* RtName internal use for RIB command names */
type RtName string

func (s RtName) String() string {
	return s.Serialise()
}

func (s RtName) Serialise() string {
	return string(s)
}

/* RtBoolean boolean value */
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

/* RtInt integer value */
type RtInt int

func (i RtInt) String() string {
	return i.Serialise()
}

func (i RtInt) Serialise() string {
	return fmt.Sprintf("%d",int(i))
}

/* RtIntArray integer array */
type RtIntArray []RtInt

func (a RtIntArray) String() string {
	return a.Serialise()
}

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

/* RtFloat float64 value */
type RtFloat float64

func (f RtFloat) String() string {
	return f.Serialise()
}

func (f RtFloat) Serialise() string {
	return fmt.Sprintf("%f",float64(f))
}

/* RtFloatArray float64 array */
type RtFloatArray []RtFloat

func (a RtFloatArray) String() string {
	return a.Serialise()
}

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

/* RtToken */
type RtToken string

func (s RtToken) String() string {
	return s.Serialise()
}

func (s RtToken) Serialise() string {
	return fmt.Sprintf("\"%s\"",string(s))
}

/* RtToken array */
type RtTokenArray []RtToken

func (a RtTokenArray) String() string {
	return a.Serialise()
}

func (a RtTokenArray) Serialise() string {
	out := ""
	for i := 0; i < len(a); i++ {
		out += a[i].Serialise()
		if i < (len(a) - 1) {
			out += " "
		}
	}
	return fmt.Sprintf("[%s]",out)
}

/* RtColor implemented as an array */
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

/* RtPoint */
type RtPoint [3]RtFloat

func (p RtPoint) String() string {
	return p.Serialise()
}

func (p RtPoint) Serialise() string {
	return fmt.Sprintf("%f %f %f",p[0],p[1],p[2])
}

/* RtPointArray */
type RtPointArray []RtPoint

func (p RtPointArray) String() string {
	return p.Serialise()
}

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

/* RtVector */
type RtVector [3]RtFloat

func (v RtVector) String() string {
	return v.Serialise()
}

func (v RtVector) Serialise() string {
	return fmt.Sprintf("[%f %f %f]",v[0],v[1],v[2])
}

/* RtNormal */
type RtNormal [3]RtFloat

func (n RtNormal) String() string {
	return n.Serialise()
}

func (n RtNormal) Serialise() string {
	return fmt.Sprintf("[%f %f %f]",n[0],n[1],n[2])
}

/* RtHpoint */
type RtHpoint [4]RtFloat

func (h RtHpoint) String() string {
	return h.Serialise()
}

func (h RtHpoint) Serialise() string {
	return fmt.Sprintf("[%f %f %f %f]",h[0],h[1],h[2],h[3])
}

/* RtMatrix */
type RtMatrix [16]RtFloat

func (m RtMatrix) String() string {
	return m.Serialise()
}

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

/* RtBasis */
type RtBasis  [16]RtFloat

func (m RtBasis) String() string {
	return m.Serialise()
}

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


/* RtBound */
type RtBound  [6]RtFloat

func (b RtBound) String() string {
	return b.Serialise()
}

func (b RtBound) Serialise() string {
	return fmt.Sprintf("[%f %f %f %f %f %f]",b[0],b[1],b[2],b[3],b[4],b[5])
}

/* RtString */
type RtString string

func (s RtString) String() string {
	return s.Serialise()
}

func (s RtString) Serialise() string {
	return fmt.Sprintf("\"%s\"",string(s))
}

/* RtStringArray array of strings */
type RtStringArray []RtString

func (a RtStringArray) String() string {
	return a.Serialise()
}

func (a RtStringArray) Serialise() string {
	out := ""
	for i := 0; i < len(a); i++ {
		out += a[i].Serialise()
		if i < len(a) - 1 {
			out += " "
		}
	}
	return fmt.Sprintf("[%s]",out)
}

/* RtLightHandle */
type RtLightHandle uint

func (l RtLightHandle) String() string {
	return l.Serialise()
}

func (l RtLightHandle) Serialise() string {
	return fmt.Sprintf("%d",uint(l))
} 

/* RtObjectHandle */
type RtObjectHandle uint

func (l RtObjectHandle) String() string {
	return l.Serialise()
}

func (l RtObjectHandle) Serialise() string {
	return fmt.Sprintf("%d",uint(l))
} 

/* RtFilterFunc */
type RtFilterFunc string

func (s RtFilterFunc) String() string {
	return s.Serialise()
}

func (s RtFilterFunc) Serialise() string {
	return fmt.Sprintf("\"%s\"",string(s))
}

/* RtProcSubdivFunc subdivision function */
type RtProcSubdivFunc string

func (s RtProcSubdivFunc) String() string {
	return s.Serialise()
}

func (s RtProcSubdivFunc) Serialise() string {
	return fmt.Sprintf("\"%s\"",string(s))
}

/* RtProcFreeFunc */
type RtProcFreeFunc string

func (s RtProcFreeFunc) String() string {
	return s.Serialise()
}

func (s RtProcFreeFunc) Serialise() string {
	return fmt.Sprintf("\"%s\"",string(s))
}

/* RtArchiveCallbackFunc */
type RtArchiveCallbackFunc string

func (s RtArchiveCallbackFunc) String() string {
	return s.Serialise()
}

func (s RtArchiveCallbackFunc) Serialise() string {
	return fmt.Sprintf("\"%s\"",string(s))
}

const (
	BoxFilter RtFilterFunc = "box"
	TriangleFilter RtFilterFunc = "triangle"
	CatmullRomFilter RtFilterFunc = "catmull-rom"
	GaussianFilter RtFilterFunc = "gaussian"
	SincFilter RtFilterFunc = "sinc"

	ReadArchiveCallback RtArchiveCallbackFunc = "ReadArchive"

	Perspective RtToken = "perspective"
	Orthographic RtToken = "orthographic"
	Bilinear RtToken = "bilinear"
	Bicubic RtToken = "bicubic"
	RGBA RtToken = "RGBA"
	P RtToken = "P"
	Pz RtToken = "Pz"
	Pw RtToken = "Pw"
	N RtToken = "N"
	Cs RtToken = "Cs"
	Os RtToken = "Os"

	ProcDelayedReadArchive RtProcSubdivFunc = "DelayedReadArchive"
	ProcRunProgram RtProcSubdivFunc = "RunProgram"
	ProcDynamicLoad RtProcSubdivFunc = "DynamicLoad"
	
	ProcFree RtProcFreeFunc = "free"

)

var (
	ErrArrayTooBig = fmt.Errorf("insufficient memory to construct array")
	ErrBadArgument = fmt.Errorf("incorrect parameter value")
	ErrBadArray    = fmt.Errorf("invalid array specification")
	ErrBadBasis    = fmt.Errorf("undefined basis matrix name")
	ErrBadColor	   = fmt.Errorf("invalid color specification")
	ErrBadHandle   = fmt.Errorf("invalid light or object handle")
	ErrBadParamlist= fmt.Errorf("parameter list type mismatch")
	ErrBadRIPCode  = fmt.Errorf("invalid encoded RIB request code")
	ErrBadStringToken = fmt.Errorf("undefined encoded string token")
	ErrBadToken = fmt.Errorf("invalid binary token")
	ErrBadVersion = fmt.Errorf("protocol version number mismatch")
	ErrLimitCheck = fmt.Errorf("overflowing an internal limit")
	ErrOutOfMemory = fmt.Errorf("generic instance of insufficient memory")
	ErrProtocolBotch = fmt.Errorf("malformed binary encoding")
	ErrStringTooBig = fmt.Errorf("insufficient memory to read string")
	ErrSyntaxError = fmt.Errorf("general syntactic error")
	ErrUnregistered = fmt.Errorf("undefined RIB request")
)




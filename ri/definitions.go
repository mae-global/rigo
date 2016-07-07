/* rigo/definitions.go */
package ri

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ImplementationVersion = "3.2.1 November 2005"
	Major                 = 3
	Minor                 = 2
	Patch                 = 1
)

type Rter interface {
	String() string
	Serialise() string
	Type() string
	Equal(Rter) bool
}

type Rtattr interface {
	Break() (RtToken,Rter)
}

type RterReader interface {
	ReadFrom() (RtName, []Rter, []Rter, []Rter)
}

type RterWriter interface {
	WriteTo(RtName, []Rter, []Rter, []Rter) error
}


type RterArray []Rter 

func (s RterArray) Type() string {
	return "rter[]"
}

func (s RterArray) String() string {
	return ""
}

func (s RterArray) Serialise() string {
	return ""
}

func (s RterArray) Equal(o Rter) bool {
	return false
}


/* RtName internal use for RIB command names */
type RtName string

func (s RtName) Type() string {
	return "name"
}

func (s RtName) String() string {
	return s.Serialise()
}

func (s RtName) Serialise() string {
	return string(s)
}

func (s RtName) Equal(o Rter) bool {
	if name, ok := o.(RtName); ok {
		return (name == s)
	}
	return false
}

func (s RtName) Prefix(pre string) RtName {
	return RtName(pre + string(s))
}

func (s RtName) Trim(pre string) RtName {
	return RtName(strings.TrimPrefix(string(s), pre))
}

/* RtBoolean boolean value */
type RtBoolean bool

func (s RtBoolean) Type() string {
	return "boolean"
}

func (s RtBoolean) String() string {
	return s.Serialise()
}

func (s RtBoolean) Serialise() string {
	if bool(s) {
		return "1"
	}
	return "0"
}

func (s RtBoolean) Equal(o Rter) bool {
	if other, ok := o.(RtBoolean); ok {
		return (other == s)
	}
	return false
}

/* RtInt integer value */
type RtInt int

func (i RtInt) Type() string {
	return "int"
}

func (i RtInt) String() string {
	return i.Serialise()
}

func (i RtInt) Serialise() string {
	return fmt.Sprintf("%d", int(i))
}

func (i RtInt) Equal(o Rter) bool {
	if other, ok := o.(RtInt); ok {
		return (other == i)
	}
	return false
}

/* RtIntArray integer array */
type RtIntArray []RtInt

func (a RtIntArray) Type() string {
	return "int[]"
}

func (a RtIntArray) String() string {
	return a.Serialise()
}

func (a RtIntArray) Serialise() string {
	out := ""
	for i, r := range a {
		out += fmt.Sprintf("%d", int(r))
		if i < (len(a) - 1) {
			out += " "
		}
	}
	return fmt.Sprintf("[%s]", out)
}

func (a RtIntArray) Equal(o Rter) bool {
	if other, ok := o.(RtIntArray); ok {
		if len(other) != len(a) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if other[i] != a[i] {
				return false
			}
		}
		return true
	}
	return false
}

/* RtFloat float64 value */
type RtFloat float64

func (f RtFloat) Type() string {
	return "float"
}

func (f RtFloat) String() string {
	return f.Serialise()
}

func (f RtFloat) Serialise() string {
	return fmt.Sprintf("%s", reduce(f))
}

func (f RtFloat) Equal(o Rter) bool {
	if other, ok := o.(RtFloat); ok {
		return (other == f)
	}
	return false
}

/* RtFloatArray float64 array */
type RtFloatArray []RtFloat

func (a RtFloatArray) Type() string {
	return "float[]"
}

func (a RtFloatArray) String() string {
	return a.Serialise()
}

func (a RtFloatArray) Serialise() string {
	out := ""
	for i, r := range a {
		out += fmt.Sprintf("%s", reduce(r))
		if i < (len(a) - 1) {
			out += " "
		}
	}
	return fmt.Sprintf("[%s]", out)
}

func (a RtFloatArray) Equal(o Rter) bool {
	if other, ok := o.(RtFloatArray); ok {
		if len(other) != len(a) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if other[i] != a[i] {
				return false
			}
		}
		return true
	}
	return false
}

/* RtToken */
type RtToken string

func (s RtToken) Type() string {
	return "token"
}

func (s RtToken) String() string {
	return s.Serialise()
}

func (s RtToken) Serialise() string {
	return fmt.Sprintf("\"%s\"", string(s))
}

func (s RtToken) Equal(o Rter) bool {
	if other, ok := o.(RtToken); ok {
		return (other == s)
	}
	return false
}

/* RtToken array */
type RtTokenArray []RtToken

func (a RtTokenArray) Type() string {
	return "token[]"
}

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
	return fmt.Sprintf("[%s]", out)
}

func (a RtTokenArray) Equal(o Rter) bool {
	if other, ok := o.(RtTokenArray); ok {
		if len(other) != len(a) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if other[i] != a[i] {
				return false
			}
		}
		return true
	}
	return false
}

/* RtColor implemented as an array */
type RtColor []RtFloat

func (c RtColor) Type() string {
	return "color"
}

func (c RtColor) String() string {
	return c.Serialise()
}

func (c RtColor) Serialise() string {
	out := ""
	for i, r := range c {
		out += fmt.Sprintf("%s", reduce(r))
		if i < (len(c) - 1) {
			out += " "
		}
	}
	return fmt.Sprintf("[%s]", out)
}

func (c RtColor) Equal(o Rter) bool {
	if other, ok := o.(RtColor); ok {
		if len(other) != len(c) {
			return false
		}
		for i := 0; i < len(c); i++ {
			if other[i] != c[i] {
				return false
			}
		}
		return true
	}
	return false
}

func Str2Color(str string) RtColor {

	parts := strings.Split(strings.TrimSpace(str), " ")
	out := make([]RtFloat, 0)

	for _, part := range parts {
		if f, err := strconv.ParseFloat(part, 64); err != nil {
			/* eat error */
			continue
		} else {
			out = append(out, RtFloat(f))
		}
	}

	return RtColor(out)
}

/* RtPoint */
type RtPoint [3]RtFloat

func (p RtPoint) Type() string {
	return "point"
}

func (p RtPoint) String() string {
	return p.Serialise()
}

func (p RtPoint) Serialise() string {
	return fmt.Sprintf("%s %s %s", reduce(p[0]), reduce(p[1]), reduce(p[2]))
}

func (p RtPoint) Equal(o Rter) bool {
	if other, ok := o.(RtPoint); ok {
		return (other[0] == p[0] && other[1] == p[1] && other[2] == p[2])
	}
	return false
}

/* RtPointArray */
type RtPointArray []RtPoint

func (p RtPointArray) Type() string {
	return "point[]"
}

func (p RtPointArray) String() string {
	return p.Serialise()
}

func (p RtPointArray) Serialise() string {
	out := ""
	for i := 0; i < len(p); i++ {
		out += p[i].Serialise()
		if i < len(p)-1 {
			out += " "
		}
	}
	return fmt.Sprintf("[%s]", out)
}

func (p RtPointArray) Equal(o Rter) bool {
	if other, ok := o.(RtPointArray); ok {
		if len(other) != len(p) {
			return false
		}
		for i := 0; i < len(p); i++ {
			p0 := p[i]
			p1 := other[i]
			if p1[0] != p0[0] || p1[1] != p0[1] || p1[2] != p0[2] {
				return false
			}
		}
		return true
	}
	return false
}

/* RtInternal */
type RtInterval [2]RtFloat

func (i RtInterval) Type() string {
	return "interval"
}

func (i RtInterval) String() string {
	return i.Serialise()
}

func (i RtInterval) Serialise() string {
	return fmt.Sprintf("%s %s", reduce(i[0]), reduce(i[1]))
}

func (i RtInterval) Equal(o Rter) bool {
	if other, ok := o.(RtInterval); ok {
		return (other[0] == i[0] && other[1] == i[1])
	}
	return false
}

/* RtVector */
type RtVector [3]RtFloat

func (v RtVector) Type() string {
	return "vector"
}

func (v RtVector) String() string {
	return v.Serialise()
}

func (v RtVector) Serialise() string {
	return fmt.Sprintf("[%s %s %s]", reduce(v[0]), reduce(v[1]), reduce(v[2]))
}

func (v RtVector) Equal(o Rter) bool {
	if other, ok := o.(RtVector); ok {
		return (other[0] == v[0] && other[1] == v[1] && other[2] == v[2])
	}
	return false
}

func Str2Vector(str string) RtVector {

	parts := strings.Split(strings.TrimSpace(str), " ")

	if len(parts) != 3 {
		return RtVector{0, 0, 0}
	}
	out := RtVector{0, 0, 0}

	for i, part := range parts {
		if f, err := strconv.ParseFloat(part, 64); err != nil {
			/* eat error */
			continue
		} else {
			out[i] = RtFloat(f)
		}
	}

	return out
}

/* RtNormal */
type RtNormal [3]RtFloat

func (n RtNormal) Type() string {
	return "normal"
}

func (n RtNormal) String() string {
	return n.Serialise()
}

func (n RtNormal) Serialise() string {
	return fmt.Sprintf("[%s %s %s]", reduce(n[0]), reduce(n[1]), reduce(n[2]))
}

func (n RtNormal) Equal(o Rter) bool {
	if other, ok := o.(RtNormal); ok {
		return (other[0] == n[0] && other[1] == n[1] && other[2] == n[2])
	}
	return false
}

func Str2Normal(str string) RtNormal {

	parts := strings.Split(strings.TrimSpace(str), " ")

	if len(parts) != 3 {
		return RtNormal{0, 0, 0}
	}
	out := RtNormal{0, 0, 0}

	for i, part := range parts {
		if f, err := strconv.ParseFloat(part, 64); err != nil {
			/* eat error */
			continue
		} else {
			out[i] = RtFloat(f)
		}
	}

	return out
}

/* RtHpoint */
type RtHpoint [4]RtFloat

func (h RtHpoint) Type() string {
	return "hpoint"
}

func (h RtHpoint) String() string {
	return h.Serialise()
}

func (h RtHpoint) Serialise() string {
	return fmt.Sprintf("[%s %s %s %s]", reduce(h[0]), reduce(h[1]), reduce(h[2]), reduce(h[3]))
}

func (h RtHpoint) Equal(o Rter) bool {
	if other, ok := o.(RtHpoint); ok {
		return (other[0] == h[0] && other[1] == h[1] && other[2] == h[2] && other[3] == h[3])
	}
	return false
}

/* RtMatrix */
type RtMatrix [16]RtFloat

func (m RtMatrix) Type() string {
	return "matrix"
}

func (m RtMatrix) String() string {
	return m.Serialise()
}

func (m RtMatrix) Serialise() string {
	out := ""
	for i := 0; i < 16; i++ {
		out += fmt.Sprintf("%s", reduce(m[i]))
		if i < 15 {
			out += " "
		}
	}

	return fmt.Sprintf("[%s]", out)
}

func (m RtMatrix) Equal(o Rter) bool {
	if other, ok := o.(RtMatrix); ok {
		for i := 0; i < 16; i++ {
			if other[i] != m[i] {
				return false
			}
		}
		return true
	}
	return false
}

/* RtBasis */
type RtBasis [16]RtFloat

func (m RtBasis) Type() string {
	return "basis"
}

func (m RtBasis) String() string {
	return m.Serialise()
}

func (b RtBasis) Serialise() string {
	out := ""
	for i := 0; i < 16; i++ {
		out += fmt.Sprintf("%s", reduce(b[i]))
		if i < 15 {
			out += " "
		}
	}
	return fmt.Sprintf("[%s]", out)
}

func (b RtBasis) Equal(o Rter) bool {
	if other, ok := o.(RtBasis); ok {
		for i := 0; i < 16; i++ {
			if other[i] != b[i] {
				return false
			}
		}
		return true
	}
	return false
}

/* RtBound */
type RtBound [6]RtFloat

func (b RtBound) Type() string {
	return "bound"
}

func (b RtBound) String() string {
	return b.Serialise()
}

func (b RtBound) Serialise() string {
	return fmt.Sprintf("[%s %s %s %s %s %s]", reduce(b[0]), reduce(b[1]), reduce(b[2]), reduce(b[3]), reduce(b[4]), reduce(b[5]))
}

func (b RtBound) Equal(o Rter) bool {
	if other, ok := o.(RtBound); ok {
		for i := 0; i < 6; i++ {
			if other[i] != b[i] {
				return false
			}
		}
		return true
	}
	return false
}

/* RtString */
type RtString string

func (s RtString) Type() string {
	return "string"
}

func (s RtString) String() string {
	return s.Serialise()
}

func (s RtString) Serialise() string {
	return fmt.Sprintf("\"%s\"", string(s))
}

func (s RtString) Equal(o Rter) bool {
	if other, ok := o.(RtString); ok {
		return (other == s)
	}
	return false
}

/* RtStringArray array of strings */
type RtStringArray []RtString

func (a RtStringArray) Type() string {
	return "string[]"
}

func (a RtStringArray) String() string {
	return a.Serialise()
}

func (a RtStringArray) Serialise() string {
	out := ""
	for i := 0; i < len(a); i++ {
		out += a[i].Serialise()
		if i < len(a)-1 {
			out += " "
		}
	}
	return fmt.Sprintf("[%s]", out)
}

func (a RtStringArray) Equal(o Rter) bool {
	if other, ok := o.(RtStringArray); ok {
		if len(other) != len(a) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if other[i] != a[i] {
				return false
			}
		}
		return true
	}
	return false
}

/* RtFilterFunc */
type RtFilterFunc string

func (s RtFilterFunc) Type() string {
	return "filterfunc"
}

func (s RtFilterFunc) String() string {
	return s.Serialise()
}

func (s RtFilterFunc) Serialise() string {
	return fmt.Sprintf("\"%s\"", string(s))
}

func (s RtFilterFunc) Equal(o Rter) bool {
	if other, ok := o.(RtFilterFunc); ok {
		return (other == s)
	}
	return false
}


/* RtErrorHandler */
type RtErrorHandler string

func (s RtErrorHandler) Type() string {
	return "errorhandler"
}

func (s RtErrorHandler) String() string {
	return s.Serialise()
}

func (s RtErrorHandler) Serialise() string {
	return fmt.Sprintf("\"%s\"",string(s))
}

func (s RtErrorHandler) Equal(o Rter) bool {
	if other,ok := o.(RtErrorHandler); ok {
		return (other == s)
	}
	return false
}


/* RtProcSubdivFunc subdivision function */
type RtProcSubdivFunc string

func (s RtProcSubdivFunc) Type() string {
	return "procsubdivfunc"
}

func (s RtProcSubdivFunc) String() string {
	return s.Serialise()
}

func (s RtProcSubdivFunc) Serialise() string {
	return fmt.Sprintf("\"%s\"", string(s))
}

func (s RtProcSubdivFunc) Equal(o Rter) bool {
	if other, ok := o.(RtProcSubdivFunc); ok {
		return (other == s)
	}
	return false
}

/* RtProc2SubdivFunc */
type RtProc2SubdivFunc string

func (s RtProc2SubdivFunc) Type() string {
	return "proc2subdivfunc"
}

func (s RtProc2SubdivFunc) String() string {
	return s.Serialise()
}

func (s RtProc2SubdivFunc) Serialise() string {
	return fmt.Sprintf("\"%s\"", string(s))
}

func (s RtProc2SubdivFunc) Equal(o Rter) bool {
	if other, ok := o.(RtProc2SubdivFunc); ok {
		return (other == s)
	}
	return false
}

/* RtProc2BoundFunc */
type RtProc2BoundFunc string

func (s RtProc2BoundFunc) Type() string {
	return "proc2boundfunc"
}

func (s RtProc2BoundFunc) String() string {
	return s.Serialise()
}

func (s RtProc2BoundFunc) Serialise() string {
	return fmt.Sprintf("\"%s\"", string(s))
}

func (s RtProc2BoundFunc) Equal(o Rter) bool {
	if other, ok := o.(RtProc2BoundFunc); ok {
		return (other == s)
	}
	return false
}

/* RtProcFreeFunc */
type RtProcFreeFunc string

func (s RtProcFreeFunc) Type() string {
	return "procfreefunc"
}

func (s RtProcFreeFunc) String() string {
	return s.Serialise()
}

func (s RtProcFreeFunc) Serialise() string {
	return fmt.Sprintf("\"%s\"", string(s))
}

func (s RtProcFreeFunc) Equal(o Rter) bool {
	if other, ok := o.(RtProcFreeFunc); ok {
		return (other == s)
	}
	return false
}

/* RtArchiveCallbackFunc */
type RtArchiveCallbackFunc string

func (s RtArchiveCallbackFunc) Type() string {
	return "archivecallbackfunc"
}

func (s RtArchiveCallbackFunc) String() string {
	return s.Serialise()
}

func (s RtArchiveCallbackFunc) Serialise() string {
	return fmt.Sprintf("\"%s\"", string(s))
}

func (s RtArchiveCallbackFunc) Equal(o Rter) bool {
	if other, ok := o.(RtArchiveCallbackFunc); ok {
		return (other == s)
	}
	return false
}

/* RtAnnotation (TODO: move this to RtxAnnotation as it does not belong in the Ri spec.) */
type RtAnnotation string

func (s RtAnnotation) Type() string {
	return "annotation"
}

func (s RtAnnotation) String() string {
	return s.Serialise()
}

func (s RtAnnotation) Serialise() string {
	if len(s) == 0 {
		return ""
	}
	return "#" + string(s)
}

func (s RtAnnotation) Equal(o Rter) bool {
	if other, ok := o.(RtAnnotation); ok {
		return (other == s)
	}
	return false
}

const (
	PARAMETERLIST RtToken = "_PARAMETERLIST_"
	DEPTH         RtName  = "_DEPTH_"
	DEBUGBARRIER  RtName  = "-->"
)

const (
	BesselFilter              RtFilterFunc = "bessel"
	BlackmanHarrisFilter      RtFilterFunc = "blackman-harris"
	BoxFilter                 RtFilterFunc = "box"
	TriangleFilter            RtFilterFunc = "triangle"
	DiskFilter                RtFilterFunc = "disk"
	CatmullRomFilter          RtFilterFunc = "catmull-rom"
	GaussianFilter            RtFilterFunc = "gaussian"
	MitchellFilter            RtFilterFunc = "mitchell"
	LanczosFilter             RtFilterFunc = "lanczos"
	SeperableCatmullRomFilter RtFilterFunc = "seperable-catmull-rom"
	SincFilter                RtFilterFunc = "sinc"

	ErrorIgnore								RtErrorHandler = "ignore"
	ErrorPrint								RtErrorHandler = "print"
	ErrorAbort								RtErrorHandler = "abort"

	ReadArchiveCallback RtArchiveCallbackFunc = "ReadArchive"

	Uniform RtToken = "uniform"
	Vertex  RtToken = "vertex"
	Varying RtToken = "varying"

	ProcDelayedReadArchive RtProcSubdivFunc = "DelayedReadArchive"
	ProcRunProgram         RtProcSubdivFunc = "RunProgram"
	ProcDynamicLoad        RtProcSubdivFunc = "DynamicLoad"

	Proc2DelayedReadArchive RtProc2SubdivFunc = "DelayedReadArchive"
	Proc2DynamicLoad        RtProc2SubdivFunc = "DynamicLoad"

	SimpleBound RtProc2BoundFunc = "Bound"
	DSOBound    RtProc2BoundFunc = "DSOBound"

	ProcFree RtProcFreeFunc = "free"

	StructuralHint RtName = "##"
	RIBStructure   RtName = "##RenderMan RIB-Structure 1.1"
)

var (
	ErrArrayTooBig    = fmt.Errorf("insufficient memory to construct array")
	ErrBadArgument    = fmt.Errorf("incorrect parameter value")
	ErrBadArray       = fmt.Errorf("invalid array specification")
	ErrBadBasis       = fmt.Errorf("undefined basis matrix name")
	ErrBadColor       = fmt.Errorf("invalid color specification")
	ErrBadHandle      = fmt.Errorf("invalid light or object handle")
	ErrBadParamlist   = fmt.Errorf("parameter list type mismatch")
	ErrBadRIBCode     = fmt.Errorf("invalid encoded RIB request code")
	ErrBadStringToken = fmt.Errorf("undefined encoded string token")
	ErrBadToken       = fmt.Errorf("invalid binary token")
	ErrBadVersion     = fmt.Errorf("protocol version number mismatch")
	ErrLimitCheck     = fmt.Errorf("overflowing an internal limit")
	ErrOutOfMemory    = fmt.Errorf("generic instance of insufficient memory")
	ErrProtocolBotch  = fmt.Errorf("malformed binary encoding")
	ErrStringTooBig   = fmt.Errorf("insufficient memory to read string")
	ErrSyntaxError    = fmt.Errorf("general syntactic error")
	ErrUnregistered   = fmt.Errorf("undefined RIB request")
	ErrNotSupported   = fmt.Errorf("not supported at this time")
)

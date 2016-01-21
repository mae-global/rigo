/* rigo/definitions.go */
package ri

import (
	"fmt"
)

type RtBoolean bool
type RtInt int
type RtFloat float64

func (f RtFloat) String() string {
	return fmt.Sprintf("%f",float64(f))
}

type RtToken string
type RtColor [3]RtFloat
type RtPoint [3]RtFloat
type RtVector [3]RtFloat
type RtNormal [3]RtFloat
type RtHpoint [4]RtFloat
type RtMatrix [4][4]RtFloat
type RtBasis  [4][4]RtFloat
type RtBound  [6]RtFloat
type RtString string

func (s RtString) String() string {
	return fmt.Sprintf("\"%s\"",string(s))
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
)

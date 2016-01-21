/* rigo/definitions.go */
package ri



type RtBoolean bool
type RtInt int
type RtFloat float64
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

type RtFilterFunc interface {
	Filter(RtFloat,RtFloat,RtFloat,RtFloat)
}

type RtErrorHandler interface {
	Handle(RtInt,RtInt,string)
}

type RtProcSubdivFunc interface {
	Subdiv(interface{},RtFloat)
}

type RtProcFreeFunc interface {
	Free(interface{})
}

type RtArchiveCallback interface {
	Archive(RtToken,string,...interface{})
}

type RtObjectHandle interface{}
type RtLightHandle interface{}
type RtContextHandle interface{}





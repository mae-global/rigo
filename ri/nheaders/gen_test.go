package main

import (
	"fmt"

	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func concat(p []string) string {
	out := ""
	for _,l := range p {
		if len(out) > 0 {
			out += " "
		}
		out += l
	}
	return out
}



func Test_ParseCPrototype(t *testing.T) {

	Convey("Parse C Prototype",t,func() {

		prototypes := ParseCPrototype("RiWorldEnd()")
		So(len(prototypes),ShouldEqual,1)
		So(concat(prototypes),ShouldEqual,"WorldEnd")

		prototypes = ParseCPrototype("RiShader(RtToken name, RtToken handle, ...)")
		for i,l := range prototypes {
			fmt.Printf("[%03d] \"%s\"\n",i,l)
		}

		So(len(prototypes),ShouldEqual,4)
		So(concat(prototypes),ShouldEqual,"Shader token name token handle ...")



		prototypes = ParseCPrototype(`RiNuPatch(RtInt nu, RtInt uorder,                                RtFloat* uknot, RtFloat umin, RtFloat umax,                                RtInt nv, RtInt vorder, RtFloat* vknot,                                RtFloat vmin, RtFloat vmax, ...)`)

		for i,l := range prototypes {
			fmt.Printf("[%03d] \"%s\"\n",i,l)
		}
		So(concat(prototypes),ShouldEqual,`NuPatch int nu int uorder float_array uknot float umin float umax int nv int vorder float_array vknot float vmin float vmax ...`)



		prototypes = ParseCPrototype(`RiIlluminate(RtLightHandle light, RtBoolean onoff)`)

		for i,l := range prototypes {
			fmt.Printf("[%03d] \"%s\"\n",i,l)
		}
		So(concat(prototypes),ShouldEqual,`Illuminate lighthandle light boolean onoff`)


		prototypes = ParseCPrototype(`RiSystem(char * args)`)
	
		for i,l := range prototypes {
			fmt.Printf("[%03d] \"%s\"\n",i,l)
		}
		So(concat(prototypes),ShouldEqual,`System string args`)
		

	})
}

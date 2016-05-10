package rigo

import (
	"fmt"
	"testing"

	. "github.com/mae-global/rigo/ri"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_RIBExamples(t *testing.T) {

	Convey("RIB examples", t, func() {

		ri, pipe := DefaultPipeline(&Configuration{PrettyPrint: true})
		ri.Begin("output/ribexample0.rib") /* FIXME, should be in ParseString */

		So(ParseString(RIBExample0, ri), ShouldBeNil)

		So(ri.End(), ShouldBeNil) /* FIXME, should be in ParseString */

		/* output gathered stats */
		p := pipe.GetByName(PipeToStats{}.Name())
		So(p, ShouldNotBeNil)
		s, ok := p.(*PipeToStats)
		So(s, ShouldNotBeNil)
		So(ok, ShouldBeTrue)

		p = pipe.GetByName(PipeTimer{}.Name())
		So(p, ShouldNotBeNil)
		t, ok := p.(*PipeTimer)
		So(t, ShouldNotBeNil)
		So(ok, ShouldBeTrue)

		fmt.Printf("%s%s", s, t)
	})
}

const RIBExample0 = `##RenderMan RIB-Structure 1.1
version 3.04
Projection "perspective" "fov" 30.0
Color [1 0 0]
Sphere 1 -1 1 360
`

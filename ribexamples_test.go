package rigo

import (
	"fmt"
	"testing"
	"os"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_RIBExamples(t *testing.T) {

	Convey("RIB examples", t, func() {
		
		Convey("Parse RIB String",func() {

			ri, pipe := DefaultPipeline(&Configuration{PrettyPrint: true})
			ri.Begin("output/ribexample0.rib") 

			So(ri.ParseRIBString(RIBExample0), ShouldBeNil)

			So(ri.End(), ShouldBeNil) 

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
		
		Convey("Parse RIB File",func() {

			f,err := os.Open("output/ribexample0.rib")
			So(err,ShouldBeNil)
			So(f,ShouldNotBeNil)
			defer f.Close()

			ri, pipe := DefaultPipeline(&Configuration{PrettyPrint: true})
			ri.Begin("output/ribexample1.rib")

			So(ri.ParseRIB(f),ShouldBeNil)

			So(ri.End(),ShouldBeNil)

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
	})
}

const RIBExample0 = `##RenderMan RIB-Structure 1.1
version 3.04
Projection "perspective" "fov" 30.0
Color [1 0 0]
Sphere 1 -1 1 360
`



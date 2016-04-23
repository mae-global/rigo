package rigo

import (
	"fmt"
	"os/user"
	"testing"
	"time"

	. "github.com/mae-global/rigo/ri"
	. "github.com/mae-global/rigo/ri/handles"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_ExampleD14(t *testing.T) {

	Convey("Example D1.4", t, func() {

		sgroup := RtString("shadinggroup")
		frames := 2
		cuser, err := user.Current()
		So(err, ShouldBeNil)

			/* create a formal version of the RIB output to the pipe, RiBegin,RiEnd... instead of Begin,End... */
		ri,pipe := DefaultPipeline(&Configuration{Formal:true,PrettyPrint:true})
		ri.Begin("output/exampleD14.rib")
		ri.ArchiveRecord("structure", "Scene Bouncing Ball")
		ri.ArchiveRecord("structure", "Creator %s", Author)
		ri.ArchiveRecord("structure", "CreationDate %s", time.Now())
		ri.ArchiveRecord("structure", "For %s", cuser.Username)
		ri.ArchiveRecord("structure", "Frames %d", frames)
		ri.ArchiveRecord("structure", "Shaders PIXARmarble, PIXARwood, MyUserShader")
		ri.ArchiveRecord("structure", "CapabilitiesNeeded ShadingLanguage Displacements")
		ri.Declare("d", "uniform point")
		ri.Declare("squish", "uniform float")
		ri.Option("limits", RtToken("bucketsize"), RtIntArray{6, 6})
		ri.Option("limits", RtToken("gridsize"), RtIntArray{18})
		ri.Format(1024, 768, 1)
		ri.Projection(PERSPECTIVE)
		ri.Clipping(10, 1000.0)
		ri.FrameBegin(1)
		ri.ArchiveRecord("structure", "Shaders PIXARmarble, PIXARwood")
		ri.ArchiveRecord("structure", "CameraOrientation %.1f %.1f %.1f %.1f %.1f %.1f", 10., 10., 10., 0., 0., 0.)
		ri.Transform(RtMatrix{.707107, -.408248, 0.57735, 0, 0, .816497, -.57735, 0, -.707107, -.408248, -.57735, 0, 0, 0, 17.3205, 1})
		ri.WorldBegin()
		ri.AttributeBegin()
		ri.Attribute("identifier", RtString("name"), RtString("myball"))
		ri.Displacement("MyUserShader", RtString("squish"), RtInt(5))
		ri.AttributeBegin()
		ri.Attribute("identifier", sgroup, RtStringArray{"tophalf"})
		ri.Surface("plastic")
		ri.Sphere(.5, -.5, 0, 360)
		ri.AttributeEnd()
		ri.AttributeBegin()
		ri.Attribute("identifier", sgroup, RtStringArray{"bothalf"})
		ri.Surface("PIXARmarble")
		ri.Sphere(.5, 0, .5, 360)
		ri.AttributeEnd()
		ri.AttributeEnd()
		ri.AttributeBegin()
		ri.Attribute("identifier", RtString("name"), RtStringArray{"floor"})
		ri.Surface("PIXARwood", RtString("roughness"), RtFloatArray{.3}, RtToken("d"), RtIntArray{1})
		ri.Comment("geometry for floor")
		ri.Polygon(4, RtToken("P"), RtFloatArray{-100, 0, -100, -100, 0, 100, 100, 0, 100, 10, 0, -100})
		ri.AttributeEnd()
		ri.WorldEnd()
		ri.FrameEnd()

		ri.FrameBegin(2)
		ri.ArchiveRecord("structure", "Shaders PIXARwood, PIXARmarbles")
		ri.ArchiveRecord("structure", "CameraOrientation %.1f %.1f %.1f %.1f %.1f %.1f", 10., 20., 10., 0., 0., 0.)
		ri.Transform(RtMatrix{.707107, -.57735, -.408248, 0, 0, .57735, -.815447, 0, -.707107, -.57735, -.408248, 0, 0, 0, 24.4949, 1})
		ri.WorldBegin()
		ri.AttributeBegin()
		ri.Attribute("identifier", RtString("name"), RtStringArray{"myball"})
		ri.AttributeBegin()
		ri.Attribute("identifier", sgroup, RtStringArray{"tophalf"})
		ri.Surface("PIXARmarble")
		ri.ShadingRate(.1)
		ri.Sphere(.5, 0, .5, 360)
		ri.AttributeEnd()
		ri.AttributeBegin()
		ri.Attribute("identifier", sgroup, RtStringArray{"bothalf"})
		ri.Surface("plastic")
		ri.Sphere(.5, -.5, 0, 360)
		ri.AttributeEnd()
		ri.AttributeEnd()
		ri.AttributeBegin()
		ri.Attribute("identifier", RtString("name"), RtStringArray{"floor"})
		ri.Surface("PIXARwood", RtToken("roughness"), RtFloatArray{.3}, RtToken("d"), RtIntArray{1})
		ri.Comment("geometry for floor")
		ri.Polygon(4, RtToken("P"), RtFloatArray{-100, 0, -100, -100, 0, 100, 100, 0, 100, 10, 0, -100})
		ri.AttributeEnd()
		ri.WorldEnd()
		ri.FrameEnd()

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
}


func Test_ExampleD21(t *testing.T) {

	Convey("Example D.2.1 RIB Entity", t, func() {

		ri,pipe := EntityPipeline()
		ri.Begin("output/exampleD21.rib")
		ri.AttributeBegin("begin unit cube")
		ri.Attribute("identifier", RtToken("name"), RtToken("unitcube"))
		ri.Bound(RtBound{-.5, .5, -.5, .5, -.5, .5})
		ri.TransformBegin()

		ri.Comment("far face")
		ri.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})
		ri.Rotate(90, 0, 1, 0)

		ri.Comment("right face")
		ri.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})
		ri.Rotate(90, 0, 1, 0)

		ri.Comment("near face")
		ri.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})
		ri.Rotate(90, 0, 1, 0)

		ri.Comment("left face")
		ri.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})

		ri.TransformEnd()
		ri.TransformBegin()

		ri.Comment("bottom face")
		ri.Rotate(90, 1, 0, 0)
		ri.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})

		ri.TransformEnd()
		ri.TransformBegin()

		ri.Comment("top face")
		ri.Rotate(-90, 1, 0, 0)
		ri.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})

		ri.TransformEnd()
		ri.AttributeEnd("end unit cube")

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
}

func Test_SimpleExample(t *testing.T) {

	Convey("Simple Example",t,func() {
	
		pipe := DefaultFilePipe()
	
		/* use a custom unique generator with a prefix for the light handles */
		lights := NewPrefixLightUniqueGenerator("light_")
	
		ctx := NewContext(pipe,lights,nil,nil,&Configuration{PrettyPrint:true})
		ri := RI(ctx)
		ri.Begin("output/simple.rib")
		ri.Display("sphere.tif","file","rgb")
		ri.Format(320,240,1)
		ri.Projection(PERSPECTIVE,RtString("fov"),RtFloat(30))
		ri.Translate(0,0,6)
		ri.WorldBegin()
		ri.LightSource("ambientlight",RtString("intensity"),RtFloat(0.5))
		ri.LightSource("distantlight",RtString("intensity"),RtFloat(1.2),RtString("form"),RtIntArray{0,0,-6},RtString("to"),RtIntArray{0,0,0})
		ri.Color(RtColor{1,0,0})
		ri.Sphere(1,-1,1,360)
		ri.WorldEnd()

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
}

func Test_SimpleExampleWithConditionals(t *testing.T) {

	Convey("Simple Example with Conditionals",t,func() {
	
		pipe := DefaultFilePipe()
	
		/* use a custom unique generator with a prefix for the light handles */
		lights := NewPrefixLightUniqueGenerator("light_")
	
		ctx := NewContext(pipe,lights,nil,nil,&Configuration{PrettyPrint:true})
		ri := RI(ctx)
		ri.Begin("output/simpleconditionals.rib")
		ri.Display("sphere.tif","file","rgb")
		ri.Format(320,240,1)
		ri.Projection(PERSPECTIVE,RtString("fov"),RtFloat(30))
		ri.Translate(0,0,6)
		ri.WorldBegin()
		ri.LightSource("ambientlight",RtString("intensity"),RtFloat(0.5))
		ri.LightSource("distantlight",RtString("intensity"),RtFloat(1.2),RtString("form"),RtIntArray{0,0,-6},RtString("to"),RtIntArray{0,0,0})

		ri.Option("user",RtString("string renderpass"),RtString("red"))
		ri.IfBegin("$user:renderpass == 'red'")
		ri.Color(RtColor{1,0,0})
		ri.ElseIf("$user:renderpass == 'blue'")
		ri.Color(RtColor{0,0,1})
		ri.Else()
		ri.Color(RtColor{0,1,0})
		ri.IfEnd()
		

		ri.Sphere(1,-1,1,360)
		ri.WorldEnd()

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
}

const hello_world = `
package main

import "fmt"

func main() {
	fmt.Printf("Hello there $name\n")
}
`

func Test_Archive(t *testing.T) {

	Convey("Archive",t,func() {
	
		ri,pipe := DefaultPipeline(nil)
		ri.Begin("output/archive.rib")
		aw,err := ri.ArchiveBegin("test",RtToken("Content-Type"),RtString("application/go"))
		So(err,ShouldBeNil)
		So(aw,ShouldNotBeNil)

		/* attempt to open a new archive */
		aw1,err := ri.ArchiveBegin("test1")
		So(err,ShouldEqual,ErrNotSupported)
		So(aw1,ShouldBeNil)

		aw.Write([]byte(hello_world))
		So(ri.ArchiveEnd("test"),ShouldBeNil)

		ri.ArchiveInstance("test",RtString("string name"),RtString("Alice"))

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
}


/* go test -bench=. */
func Benchmark_SimpleExampleNumberHandlers(b *testing.B) {

	for i := 0; i < b.N; i++ {
		ri,_ := DefaultPipeline(nil)
		ri.Begin("simple.rib")
		ri.Display("sphere.tif","file","rgb")
		ri.Format(320,240,1)
		ri.Projection(PERSPECTIVE,RtToken("fov"),RtFloat(30))
		ri.Translate(0,0,6)
		ri.WorldBegin()
		ri.LightSource("ambientlight",RtToken("intensity"),RtFloat(0.5))
		ri.LightSource("distantlight",RtToken("intensity"),RtFloat(1.2),RtToken("form"),RtIntArray{0,0,-6},RtString("to"),RtIntArray{0,0,0})
		ri.Color(RtColor{1,0,0})
		ri.Sphere(1,-1,1,360)
		ri.WorldEnd()
		ri.End()
	}
}

func Benchmark_SimpleExampleUniqueHandlers(b *testing.B) {

	for i := 0; i < b.N; i++ {
		pipe := NullPipe()
		ctx := NewContext(pipe,NewLightUniqueGenerator(),nil,nil,nil)
		ri := RI(ctx)
		ri.Begin("simple.rib")
		ri.Display("sphere.tif","file","rgb")
		ri.Format(320,240,1)
		ri.Projection(PERSPECTIVE,RtString("fov"),RtFloat(30))
		ri.Translate(0,0,6)
		ri.WorldBegin()
		ri.LightSource("ambientlight",RtString("intensity"),RtFloat(0.5))
		ri.LightSource("distantlight",RtString("intensity"),RtFloat(1.2),RtString("form"),RtIntArray{0,0,-6},RtString("to"),RtIntArray{0,0,0})
		ri.Color(RtColor{1,0,0})
		ri.Sphere(1,-1,1,360)
		ri.WorldEnd()
		ri.End()
	}
}







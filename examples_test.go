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

		pipe := DefaultFilePipe()
		So(pipe, ShouldNotBeNil)

		/* create a formal version of the RIB output to the pipe, RiBegin,RiEnd... instead of Begin,End... */
		ctx := New(pipe, &Configuration{Formal: true,PrettyPrint: true})
		ctx.Begin("output/exampleD14.rib")
		ctx.ArchiveRecord("structure", "Scene Bouncing Ball")
		ctx.ArchiveRecord("structure", "Creator %s-%s", Author, Version)
		ctx.ArchiveRecord("structure", "CreationDate %s", time.Now())
		ctx.ArchiveRecord("structure", "For %s", cuser.Username)
		ctx.ArchiveRecord("structure", "Frames %d", frames)
		ctx.ArchiveRecord("structure", "Shaders PIXARmarble, PIXARwood, MyUserShader")
		ctx.ArchiveRecord("structure", "CapabilitiesNeeded ShadingLanguage Displacements")
		ctx.Declare("d", "uniform point")
		ctx.Declare("squish", "uniform float")
		ctx.Option("limits", RtToken("bucketsize"), RtIntArray{6, 6})
		ctx.Option("limits", RtToken("gridsize"), RtIntArray{18})
		ctx.Format(1024, 768, 1)
		ctx.Projection(Perspective)
		ctx.Clipping(10, 1000.0)
		ctx.FrameBegin(1)
		ctx.ArchiveRecord("structure", "Shaders PIXARmarble, PIXARwood")
		ctx.ArchiveRecord("structure", "CameraOrientation %.1f %.1f %.1f %.1f %.1f %.1f", 10., 10., 10., 0., 0., 0.)
		ctx.Transform(RtMatrix{.707107, -.408248, 0.57735, 0, 0, .816497, -.57735, 0, -.707107, -.408248, -.57735, 0, 0, 0, 17.3205, 1})
		ctx.WorldBegin()
		ctx.AttributeBegin()
		ctx.Attribute("identifier", RtString("name"), RtString("myball"))
		ctx.Displacement("MyUserShader", RtString("squish"), RtInt(5))
		ctx.AttributeBegin()
		ctx.Attribute("identifier", sgroup, RtStringArray{"tophalf"})
		ctx.Surface("plastic")
		ctx.Sphere(.5, -.5, 0, 360)
		ctx.AttributeEnd()
		ctx.AttributeBegin()
		ctx.Attribute("identifier", sgroup, RtStringArray{"bothalf"})
		ctx.Surface("PIXARmarble")
		ctx.Sphere(.5, 0, .5, 360)
		ctx.AttributeEnd()
		ctx.AttributeEnd()
		ctx.AttributeBegin()
		ctx.Attribute("identifier", RtString("name"), RtStringArray{"floor"})
		ctx.Surface("PIXARwood", RtString("roughness"), RtFloatArray{.3}, RtToken("d"), RtIntArray{1})
		ctx.Comment("geometry for floor")
		ctx.Polygon(4, RtToken("P"), RtFloatArray{-100, 0, -100, -100, 0, 100, 100, 0, 100, 10, 0, -100})
		ctx.AttributeEnd()
		ctx.WorldEnd()
		ctx.FrameEnd()

		ctx.FrameBegin(2)
		ctx.ArchiveRecord("structure", "Shaders PIXARwood, PIXARmarbles")
		ctx.ArchiveRecord("structure", "CameraOrientation %.1f %.1f %.1f %.1f %.1f %.1f", 10., 20., 10., 0., 0., 0.)
		ctx.Transform(RtMatrix{.707107, -.57735, -.408248, 0, 0, .57735, -.815447, 0, -.707107, -.57735, -.408248, 0, 0, 0, 24.4949, 1})
		ctx.WorldBegin()
		ctx.AttributeBegin()
		ctx.Attribute("identifier", RtString("name"), RtStringArray{"myball"})
		ctx.AttributeBegin()
		ctx.Attribute("identifier", sgroup, RtStringArray{"tophalf"})
		ctx.Surface("PIXARmarble")
		ctx.ShadingRate(.1)
		ctx.Sphere(.5, 0, .5, 360)
		ctx.AttributeEnd()
		ctx.AttributeBegin()
		ctx.Attribute("identifier", sgroup, RtStringArray{"bothalf"})
		ctx.Surface("plastic")
		ctx.Sphere(.5, -.5, 0, 360)
		ctx.AttributeEnd()
		ctx.AttributeEnd()
		ctx.AttributeBegin()
		ctx.Attribute("identifier", RtString("name"), RtStringArray{"floor"})
		ctx.Surface("PIXARwood", RtToken("roughness"), RtFloatArray{.3}, RtToken("d"), RtIntArray{1})
		ctx.Comment("geometry for floor")
		ctx.Polygon(4, RtToken("P"), RtFloatArray{-100, 0, -100, -100, 0, 100, 100, 0, 100, 10, 0, -100})
		ctx.AttributeEnd()
		ctx.WorldEnd()
		ctx.FrameEnd()

		So(ctx.End(), ShouldBeNil)

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

		pipe := DefaultFilePipe()

		ctx := NewEntity(pipe)
		ctx.Begin("output/exampleD21.rib")
		ctx.AttributeBegin("begin unit cube")
		ctx.Attribute("identifier", RtToken("name"), RtToken("unitcube"))
		ctx.Bound(RtBound{-.5, .5, -.5, .5, -.5, .5})
		ctx.TransformBegin()

		ctx.Comment("far face")
		ctx.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})
		ctx.Rotate(90, 0, 1, 0)

		ctx.Comment("right face")
		ctx.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})
		ctx.Rotate(90, 0, 1, 0)

		ctx.Comment("near face")
		ctx.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})
		ctx.Rotate(90, 0, 1, 0)

		ctx.Comment("left face")
		ctx.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})

		ctx.TransformEnd()
		ctx.TransformBegin()

		ctx.Comment("bottom face")
		ctx.Rotate(90, 1, 0, 0)
		ctx.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})

		ctx.TransformEnd()
		ctx.TransformBegin()

		ctx.Comment("top face")
		ctx.Rotate(-90, 1, 0, 0)
		ctx.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})

		ctx.TransformEnd()
		ctx.AttributeEnd("end unit cube")

		So(ctx.End(), ShouldBeNil)

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
	
		ctx := NewCustom(pipe,lights,nil,&Configuration{PrettyPrint:true})
		ctx.Begin("output/simple.rib")
		ctx.Display("sphere.tif","file","rgb")
		ctx.Format(320,240,1)
		ctx.Projection(Perspective,RtString("fov"),RtFloat(30))
		ctx.Translate(0,0,6)
		ctx.WorldBegin()
		ctx.LightSource("ambientlight",RtString("intensity"),RtFloat(0.5))
		ctx.LightSource("distantlight",RtString("intensity"),RtFloat(1.2),RtString("form"),RtIntArray{0,0,-6},RtString("to"),RtIntArray{0,0,0})
		ctx.Color(RtColor{1,0,0})
		ctx.Sphere(1,-1,1,360)
		ctx.WorldEnd()

		So(ctx.End(),ShouldBeNil)

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
	
		ctx := NewCustom(pipe,lights,nil,&Configuration{PrettyPrint:true})
		ctx.Begin("output/simpleconditionals.rib")
		ctx.Display("sphere.tif","file","rgb")
		ctx.Format(320,240,1)
		ctx.Projection(Perspective,RtString("fov"),RtFloat(30))
		ctx.Translate(0,0,6)
		ctx.WorldBegin()
		ctx.LightSource("ambientlight",RtString("intensity"),RtFloat(0.5))
		ctx.LightSource("distantlight",RtString("intensity"),RtFloat(1.2),RtString("form"),RtIntArray{0,0,-6},RtString("to"),RtIntArray{0,0,0})

		ctx.Option("user",RtString("string renderpass"),RtString("red"))
		ctx.IfBegin("$user:renderpass == 'red'")
		ctx.Color(RtColor{1,0,0})
		ctx.ElseIf("$user:renderpass == 'blue'")
		ctx.Color(RtColor{0,0,1})
		ctx.Else()
		ctx.Color(RtColor{0,1,0})
		ctx.IfEnd()
		

		ctx.Sphere(1,-1,1,360)
		ctx.WorldEnd()

		So(ctx.End(),ShouldBeNil)

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
		pipe := NullPipe()
		ctx := New(pipe,nil)
		ctx.Begin("simple.rib")
		ctx.Display("sphere.tif","file","rgb")
		ctx.Format(320,240,1)
		ctx.Projection(Perspective,RtString("fov"),RtFloat(30))
		ctx.Translate(0,0,6)
		ctx.WorldBegin()
		ctx.LightSource("ambientlight",RtString("intensity"),RtFloat(0.5))
		ctx.LightSource("distantlight",RtString("intensity"),RtFloat(1.2),RtString("form"),RtIntArray{0,0,-6},RtString("to"),RtIntArray{0,0,0})
		ctx.Color(RtColor{1,0,0})
		ctx.Sphere(1,-1,1,360)
		ctx.WorldEnd()
		ctx.End()
	}
}

func Benchmark_SimpleExampleUniqueHandlers(b *testing.B) {

	for i := 0; i < b.N; i++ {
		pipe := NullPipe()
		ctx := NewCustom(pipe,NewLightUniqueGenerator(),nil,nil)
		ctx.Begin("simple.rib")
		ctx.Display("sphere.tif","file","rgb")
		ctx.Format(320,240,1)
		ctx.Projection(Perspective,RtString("fov"),RtFloat(30))
		ctx.Translate(0,0,6)
		ctx.WorldBegin()
		ctx.LightSource("ambientlight",RtString("intensity"),RtFloat(0.5))
		ctx.LightSource("distantlight",RtString("intensity"),RtFloat(1.2),RtString("form"),RtIntArray{0,0,-6},RtString("to"),RtIntArray{0,0,0})
		ctx.Color(RtColor{1,0,0})
		ctx.Sphere(1,-1,1,360)
		ctx.WorldEnd()
		ctx.End()
	}
}







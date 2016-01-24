package ri

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"os/user"
	"testing"
	"time"
)

func Test_ExampleD14(t *testing.T) {

	Convey("Example D1.4", t, func() {

		sgroup := RtString("shadinggroup")
		frames := 2
		cuser, err := user.Current()
		So(err, ShouldBeNil)

		pipe := DefaultFilePipe()
		So(pipe, ShouldNotBeNil)

		ctx := New(pipe)
		ctx.Begin("output/exampleD14.rib")
		ctx.ArchiveRecord("structure", "Scene Bouncing Ball")
		ctx.ArchiveRecord("structure", "Creator %s-%s", Author, Version)
		ctx.ArchiveRecord("structure", "CreationDate %s", time.Now())
		ctx.ArchiveRecord("structure", "For %s", cuser.Username)
		ctx.ArchiveRecord("structure", "Frames %d", frames)
		ctx.ArchiveRecord("Structure", "Shaders PIXARmarble, PIXARwood, MyUserShader")
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
		if s, ok := p.(*PipeToStats); ok {
			fmt.Printf("%s", s)
		}
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

		p := pipe.GetByName(PipeToStats{}.Name())
		So(p, ShouldNotBeNil)
		if s, ok := p.(*PipeToStats); ok {
			fmt.Printf("%s", s)
		}
	})
}

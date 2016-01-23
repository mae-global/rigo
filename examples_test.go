package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func Test_ExampleD14(t *testing.T) {

	Convey("Example D1.4",t,func() {

		sgroup := RtString("shadinggroup")

		ctx := New(nil)
		ctx.Begin("output/exampleD14.rib")
		/* TODO */
		ctx.Declare("d","uniform point")
		ctx.Declare("squish","uniform float")
		ctx.Option("limits",RtToken("bucketsize"),RtIntArray{6,6})
		ctx.Option("limits",RtToken("gridsize"),RtIntArray{18})
		ctx.Format(1024,768,1)
		ctx.Projection(Perspective)
		ctx.Clipping(10,1000.0)
		ctx.FrameBegin(1)
			/* TODO */
			ctx.Transform(RtMatrix{.707107,-.408248,0.57735,0,0,.816497,-.57735,0,-.707107,-.408248,-.57735,0,0,0,17.3205,1})
			ctx.WorldBegin()
				ctx.AttributeBegin()
					ctx.Attribute("identifier",RtString("name"),RtString("myball"))
					ctx.Displacement("MyUserShader",RtString("squish"),RtInt(5))
					ctx.AttributeBegin()
						ctx.Attribute("identifier",sgroup,RtStringArray{"tophalf"})
						ctx.Surface("plastic")		
						ctx.Sphere(.5,-.5,0,360)
					ctx.AttributeEnd()
					ctx.AttributeBegin()
						ctx.Attribute("identifier",sgroup,RtStringArray{"bothalf"})
						ctx.Surface("PIXARmarble")
						ctx.Sphere(.5,0,.5,360)
					ctx.AttributeEnd()
				ctx.AttributeEnd()
				ctx.AttributeBegin()
					ctx.Attribute("identifier",RtString("name"),RtStringArray{"floor"})
					ctx.Surface("PIXARwood",RtString("roughness"),RtFloatArray{.3},RtToken("d"),RtIntArray{1})
					ctx.Comment("geometry for floor")
					ctx.Polygon(4,RtToken("P"),RtFloatArray{-100,0,-100,-100,0,100,100,0,100,10,0,-100})
				ctx.AttributeEnd()
			ctx.WorldEnd()	
		ctx.FrameEnd()		

		ctx.FrameBegin(2)
			/* TODO */
			ctx.Transform(RtMatrix{.707107,-.57735,-.408248,0,0,.57735,-.815447,0,-.707107,-.57735,-.408248,0,0,0,24.4949,1})
			ctx.WorldBegin()
				ctx.AttributeBegin()
					ctx.Attribute("identifier",RtString("name"),RtStringArray{"myball"})
					ctx.AttributeBegin()
						ctx.Attribute("identifier",sgroup,RtStringArray{"tophalf"})
						ctx.Surface("PIXARmarble")
						ctx.ShadingRate(.1)
						ctx.Sphere(.5,0,.5,360)
					ctx.AttributeEnd()
					ctx.AttributeBegin()
						ctx.Attribute("identifier",sgroup,RtStringArray{"bothalf"})
						ctx.Surface("plastic")
						ctx.Sphere(.5,-.5,0,360)
					ctx.AttributeEnd()
				ctx.AttributeEnd()
				ctx.AttributeBegin()
					ctx.Attribute("identifier",RtString("name"),RtStringArray{"floor"})
					ctx.Surface("PIXARwood",RtToken("roughness"),RtFloatArray{.3},RtToken("d"),RtIntArray{1})
					ctx.Comment("geometry for floor")
					ctx.Polygon(4,RtToken("P"),RtFloatArray{-100,0,-100,-100,0,100,100,0,100,10,0,-100})
				ctx.AttributeEnd()
			ctx.WorldEnd()	
		ctx.FrameEnd()		

		So(ctx.End(),ShouldBeNil)
	})
}

func Test_ExampleD21(t *testing.T) {

	Convey("Example D.2.1 RIB Entity",t,func() {

		ctx := NewEntity(nil)
		ctx.Begin("output/exampleD21.rib")
		/* TODO */
		ctx.AttributeBegin("begin unit cube")
			ctx.Attribute("identifier",RtToken("name"),RtToken("unitcube"))
			ctx.Bound(RtBound{-.5,.5,-.5,.5,-.5,.5})
			ctx.TransformBegin()
				ctx.Comment("far face")
				ctx.Polygon(4,RtToken("P"),RtFloatArray{.5,.5,.5, -.5,.5,.5, -.5,-.5,.5, .5,-.5,.5})
				ctx.Rotate(90,0,1,0)

				ctx.Comment("right face")
				ctx.Polygon(4,RtToken("P"),RtFloatArray{.5,.5,.5, -.5,.5,.5, -.5,-.5,.5, .5,-.5,.5})
				ctx.Rotate(90,0,1,0)
				
				ctx.Comment("near face")
				ctx.Polygon(4,RtToken("P"),RtFloatArray{.5,.5,.5, -.5,.5,.5, -.5,-.5,.5, .5,-.5,.5})
				ctx.Rotate(90,0,1,0)	

				ctx.Comment("left face")
				ctx.Polygon(4,RtToken("P"),RtFloatArray{.5,.5,.5, -.5,.5,.5, -.5,-.5,.5, .5,-.5,.5})
			ctx.TransformEnd()
			ctx.TransformBegin()			
				ctx.Comment("bottom face")
				ctx.Rotate(90,1,0,0)
				ctx.Polygon(4,RtToken("P"),RtFloatArray{.5,.5,.5, -.5,.5,.5, -.5,-.5,.5, .5,-.5,.5})
			ctx.TransformEnd()
			ctx.TransformBegin()
				ctx.Comment("top face")
				ctx.Rotate(-90,1,0,0)
				ctx.Polygon(4,RtToken("P"),RtFloatArray{.5,.5,.5, -.5,.5,.5, -.5,-.5,.5, .5,-.5,.5})
			ctx.TransformEnd()
		ctx.AttributeEnd("end unit cube")
	
		So(ctx.End(),ShouldBeNil)
	})
}



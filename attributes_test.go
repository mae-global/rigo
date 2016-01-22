package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func Test_Attributes(t *testing.T) {

	Convey("All Attributes",t,func() {

		ctx := New(nil)
		So(ctx,ShouldNotBeNil)

		So(ctx.Begin("output/attributes.rib"),ShouldBeNil)
		So(ctx.Comment("output from rigo, attributes_test.go"),ShouldBeNil)
		So(ctx.AttributeBegin(),ShouldBeNil)
		So(ctx.Color(RtColor{0.3,0.4,0.1,0.2}),ShouldBeNil)
		So(ctx.Opacity(RtColor{1,1,1,1}),ShouldBeNil)
		So(ctx.TextureCoordinates(0,0,2,-.5,-.5,1.75,3,3),ShouldBeNil)
		spot,err := ctx.LightSource(RtToken("spotlight"),RtToken("coneangle"),RtInt(5))
		So(spot,ShouldEqual,0)
		So(err,ShouldBeNil)
		ambient,err := ctx.LightSource(RtToken("ambientlight"),RtToken("lightcolor"),RtColor{.5,0,0},RtToken("intensity"),RtFloat(.6))
		So(err,ShouldBeNil)
		So(ambient,ShouldEqual,1)
		So(spot,ShouldNotEqual,ambient)
		
		So(ctx.AttributeEnd(),ShouldBeNil)
		So(ctx.End(),ShouldBeNil)
	})
}

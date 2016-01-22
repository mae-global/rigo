package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func Test_Options(t *testing.T) {

	Convey("All Options",t,func() {

		ctx := New(nil)
		So(ctx,ShouldNotBeNil)

		So(ctx.Begin("output/options.rib"),ShouldBeNil)
		So(ctx.Comment("output from rigo, options_test.go"),ShouldBeNil)
		So(ctx.Option(RtToken("foo"),RtToken("fov"),RtFloat(3.4),RtToken("color"),RtFloatArray{.4,.4,.4},RtToken("normal"),RtIntArray{1,1,0}),ShouldBeNil)
		So(ctx.Attribute(RtToken("displacementbound"),RtToken("sphere"),RtFloat(2.0)),ShouldBeNil)

		So(ctx.Geometry("teapot"),ShouldBeNil)

		So(ctx.MotionBegin(3,0.1,0.2,0.3),ShouldBeNil)
		So(ctx.MotionEnd(),ShouldBeNil)

		So(ctx.End(),ShouldBeNil)
	})
}

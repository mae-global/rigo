package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Transformations(t *testing.T) {

	Convey("All Transformations", t, func() {

		ctx := New(nil)
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("output/transformations.rib"), ShouldBeNil)
		So(ctx.Comment("output from rigo, transformations_test.go"), ShouldBeNil)
		So(ctx.TransformBegin(), ShouldBeNil)

		So(ctx.Identity(), ShouldBeNil)
		So(ctx.Transform(RtMatrix{.5, .0, .0, .0, .0, 1., .0, .0, .0, .0, 1., .0, .0, .0, .0, 1.}), ShouldBeNil)
		So(ctx.ConcatTransform(RtMatrix{.5, .0, .0, .0, .0, 1., .0, .0, .0, .0, 1., .0, .0, .0, .0, 1.}), ShouldBeNil)
		So(ctx.Perspective(90.0), ShouldBeNil)
		So(ctx.Translate(0.0, 1.0, 0.0), ShouldBeNil)
		So(ctx.Rotate(90.0, 0.0, 1.0, 0.0), ShouldBeNil)
		So(ctx.Scale(.5, 1, 1), ShouldBeNil)
		So(ctx.Skew(45.0, 0, 1.0, 0, 1.0, 0, 0), ShouldBeNil)
		So(ctx.CoordinateSystem("lamptop"), ShouldBeNil)
		So(ctx.CoordSysTransform("lamptop"), ShouldBeNil)

		So(ctx.TransformEnd(), ShouldBeNil)
		So(ctx.End(), ShouldBeNil)
	})
}

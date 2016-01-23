package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Surfaces(t *testing.T) {

	Convey("All Surfaces", t, func() {

		ctx := New(nil)
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("output/surfaces.rib"), ShouldBeNil)
		So(ctx.Comment("output from rigo, surfaces_test.go"), ShouldBeNil)

		So(ctx.SubdivisionMesh("catmull-clark", 5, RtIntArray{4, 4, 4, 4, 4, 4}, RtIntArray{0, 4, 3, 2, 5, 6, 1, 2, 3}, 1, RtTokenArray{"interpolateboundary"}, RtIntArray{0, 0}, RtIntArray{0}, RtFloatArray{0}), ShouldBeNil)
		So(ctx.Points(5, RtToken("P"), RtFloatArray{.5, -.5, -.5, 0, .5, 0, .5, 0}, RtToken("width"), RtFloatArray{.1, .12, .05, .02}), ShouldBeNil)
		So(ctx.Curves("cubic", 4, RtIntArray{4}, "nonperiodic", RtToken("constantwidth"), RtFloat(0.1)), ShouldBeNil)
		So(ctx.Blobby(2, 6, RtIntArray{1001, 0, 1003, 0, 16, 0, 201}, 6, RtFloatArray{1.5, 0, 0, 0, 0, 1.5, 0, 0, 0, 0, 1.5, 0, 0, 0, 1}, 1, RtStringArray{"flat.zfile"}), ShouldBeNil)

		So(ctx.End(), ShouldBeNil)
	})
}

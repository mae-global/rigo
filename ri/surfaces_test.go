package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Surfaces(t *testing.T) {

	Convey("All Surfaces", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("surfaces.rib"), ErrorShouldEqual, `Begin "surfaces.rib"`)
		So(ctx.Comment("output from rigo, surfaces_test.go"), ErrorShouldEqual, `# output from rigo, surfaces_test.go`)

		So(ctx.SubdivisionMesh("catmull-clark", 5, RtIntArray{4, 4, 4, 4, 4, 4}, RtIntArray{0, 4, 3, 2, 5, 6, 1, 2, 3}, 1, RtTokenArray{"interpolateboundary"}, RtIntArray{0, 0}, RtIntArray{0}, RtFloatArray{0}),
			ErrorShouldEqual, `SubdivisionMesh "catmull-clark" [4 4 4 4 4 4] [0 4 3 2 5 6 1 2 3] ["interpolateboundary"] [0 0] [0] [0]`)

		So(ctx.Points(5, RtToken("float[8] P"), RtFloatArray{.5, -.5, -.5, 0, .5, 0, .5, 0}, RtToken("float[4] width"), RtFloatArray{.1, .12, .05, .02}), 
									ErrorShouldEqual, `Points "float[8] P" [.5 -.5 -.5 0 .5 0 .5 0] "float[4] width" [.1 .12 .05 .02]`)

		So(ctx.Curves("cubic", 4, RtIntArray{4}, "nonperiodic", RtToken("float constantwidth"), RtFloat(0.1)), ErrorShouldEqual, `Curves "cubic" [4] "nonperiodic" "float constantwidth" [.1]`)
		So(ctx.Blobby(2, 6, RtIntArray{1001, 0, 1003, 0, 16, 0, 201}, 6, RtFloatArray{1.5, 0, 0, 0, 0, 1.5, 0, 0, 0, 0, 1.5, 0, 0, 0, 1}, 1, RtStringArray{"flat.zfile"}),
			ErrorShouldEqual, `Blobby 2 [1.5 0 0 0 0 1.5 0 0 0 0 1.5 0 0 0 1] ["flat.zfile"]`)

		So(ctx.End(), ErrorShouldEqual, `End`)
	})
}

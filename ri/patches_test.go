package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Patches(t *testing.T) {

	Convey("All Patches", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("patches.rib"), ErrorShouldEqual, `Begin "patches.rib"`)
		So(ctx.Comment("output from rigo, patches_test.go"), ErrorShouldEqual, `# output from rigo, patches_test.go`)
		So(ctx.Basis(RtBasis{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 1, RtBasis{1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2}, 3), ErrorShouldEqual, `Basis [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16] 1 [1 2 1 2 1 2 1 2 1 2 1 2 1 2 1 2] 3`)

		var points = make([]RtPoint, 0)
		points = append(points, RtPoint{0, 1, 0})
		points = append(points, RtPoint{0, 1, 1})
		points = append(points, RtPoint{0, 0, 1})
		points = append(points, RtPoint{0, 0, 0})
		So(ctx.Patch(BILINEAR, RtToken("P"), RtPointArray(points)), ErrorShouldEqual, `Patch "bilinear" "P" [0 1 0 0 1 1 0 0 1 0 0 0]`)

		So(ctx.PatchMesh(BICUBIC, 7, "nonperiodic", 4, "nonperiodic", RtToken("P"), RtPointArray(points)), ErrorShouldEqual, `PatchMesh "bicubic" 7 "nonperiodic" 4 "nonperiodic" "P" [0 1 0 0 1 1 0 0 1 0 0 0]`)
		So(ctx.NuPatch(9, 3, RtFloatArray{0, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 4}, 0, 4, 2, 2, RtFloatArray{0, 0, 1, 1}, 0, 1), ErrorShouldEqual, `NuPatch 9 3 [0 0 0 1 1 2 2 3 3 4 4 4] 0 4 2 2 [0 0 1 1] 0 1`)

		So(ctx.End(), ErrorShouldEqual, `End`)
	})
}

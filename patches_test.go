package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func Test_Patches(t *testing.T) {

	Convey("All Patches",t,func() {

		ctx := New(nil)
		So(ctx,ShouldNotBeNil)

		So(ctx.Begin("output/patches.rib"),ShouldBeNil)
		So(ctx.Comment("output from rigo, patches_test.go"),ShouldBeNil)
		So(ctx.Basis(RtBasis{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16},1,RtBasis{1,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2},3),ShouldBeNil)

		var points = make([]RtPoint,0)
		points = append(points,RtPoint{0,1,0})
		points = append(points,RtPoint{0,1,1})
		points = append(points,RtPoint{0,0,1})
		points = append(points,RtPoint{0,0,0})
		So(ctx.Patch(Bilinear,RtToken("P"),RtPointArray(points)),ShouldBeNil)

		So(ctx.PatchMesh(Bicubic,7,"nonperiodic",4,"nonperiodic",RtToken("P"),RtPointArray(points)),ShouldBeNil)
		So(ctx.NuPatch(9,3,RtFloatArray{0,0,0,1,1,2,2,3,3,4,4,4},0,4,2,2,RtFloatArray{0,0,1,1},0,1),ShouldBeNil)

		So(ctx.End(),ShouldBeNil)
	})
}

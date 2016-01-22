package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func Test_Surfaces(t *testing.T) {

	Convey("All Surfaces",t,func() {

		ctx := New(nil)
		So(ctx,ShouldNotBeNil)

		So(ctx.Begin("output/surfaces.rib"),ShouldBeNil)
		So(ctx.Comment("output from rigo, surfaces_test.go"),ShouldBeNil)

		So(ctx.SubdivisionMesh("catmull-clark",5,RtIntArray{4,4,4,4,4,4},RtIntArray{0,4,3,2,5,6,1,2,3},1,RtTokenArray{"interpolateboundary"},RtIntArray{0,0},RtIntArray{0},RtFloatArray{0}),ShouldBeNil)


		So(ctx.End(),ShouldBeNil)
	})
}

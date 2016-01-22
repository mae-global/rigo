package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func Test_Geometric(t *testing.T) {

	Convey("All Geometric",t,func() {

		ctx := New(nil)
		So(ctx,ShouldNotBeNil)

		So(ctx.Begin("output/geometric.rib"),ShouldBeNil)
		So(ctx.Comment("output from rigo, geometric_test.go"),ShouldBeNil)
		
		var points = make([]RtPoint,0)
		points = append(points,RtPoint{0,1,0})
		points = append(points,RtPoint{0,1,1})
		points = append(points,RtPoint{0,0,1})
		points = append(points,RtPoint{0,0,0})
		So(ctx.Polygon(4,RtToken("P"),RtPointArray(points)),ShouldBeNil)
		
		So(ctx.GeneralPolygon(2,RtIntArray{4,3},RtToken("P"),RtPointArray(points)),ShouldBeNil)
		So(ctx.PointsPolygons(2,RtIntArray{3,3,3},RtIntArray{0,3,2,0,1,3,1,4,3},RtToken("P"),RtPointArray(points)),ShouldBeNil)
		So(ctx.PointsGeneralPolygons(2,RtIntArray{2,2},RtIntArray{4,3,4,3},RtIntArray{0,1,4,3,6,7,8,1,2,5,4,9,10,11},RtToken("P"),RtPointArray(points)),ShouldBeNil)

		So(ctx.End(),ShouldBeNil)
	})
}

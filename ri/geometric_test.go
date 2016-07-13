package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Geometric(t *testing.T) {

	Convey("All Geometric", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("geometric.rib"), ErrorShouldEqual, `Begin "geometric.rib"`)
		So(ctx.Comment("output from rigo, geometric_test.go"), ErrorShouldEqual, `# output from rigo, geometric_test.go`)

		var points = make([]RtPoint, 0)
		points = append(points, RtPoint{0, 1, 0})
		points = append(points, RtPoint{0, 1, 1})
		points = append(points, RtPoint{0, 0, 1})
		points = append(points, RtPoint{0, 0, 0})
		So(ctx.Polygon(4, RtToken("point[4] P"), RtPointArray(points)), ErrorShouldEqual, `Polygon "point[4] P" [0 1 0 0 1 1 0 0 1 0 0 0]`)
		//So(ctx.PolygonV([]Rter{RtFloat(4)}, []Rter{RtToken("P")}, []Rter{RtPointArray(points)}), ErrorShouldEqual, `Polygon 4 "P" [0 1 0 0 1 1 0 0 1 0 0 0]`)

		So(ctx.GeneralPolygon(2, RtIntArray{4, 3}, RtToken("point[4] P"), RtPointArray(points)), ErrorShouldEqual, `GeneralPolygon [4 3] "point[4] P" [0 1 0 0 1 1 0 0 1 0 0 0]`)
		//So(ctx.GeneralPolygonV([]Rter{RtFloat(2), RtIntArray{4, 3}}, []Rter{RtToken("P")}, []Rter{RtPointArray(points)}),
		//	ErrorShouldEqual, `GeneralPolygon 2 [4 3] "P" [0 1 0 0 1 1 0 0 1 0 0 0]`)

		So(ctx.PointsPolygons(2, RtIntArray{3, 3, 3}, RtIntArray{0, 3, 2, 0, 1, 3, 1, 4, 3}, RtToken("point[4] P"), RtPointArray(points)), 
															ErrorShouldEqual, `PointsPolygons [3 3 3] [0 3 2 0 1 3 1 4 3] "point[4] P" [0 1 0 0 1 1 0 0 1 0 0 0]`)
		//So(ctx.PointsPolygonsV([]Rter{RtFloat(2), RtIntArray{3, 3, 3}, RtIntArray{0, 3, 2, 0, 1, 3, 1, 4, 3}},
			//[]Rter{RtToken("P")}, []Rter{RtPointArray(points)}), ErrorShouldEqual, `PointsPolygons 2 [3 3 3] [0 3 2 0 1 3 1 4 3] "P" [0 1 0 0 1 1 0 0 1 0 0 0]`)

		So(ctx.PointsGeneralPolygons(2, RtIntArray{2, 2}, RtIntArray{4, 3, 4, 3}, RtIntArray{0, 1, 4, 3, 6, 7, 8, 1, 2, 5, 4, 9, 10, 11},
			RtToken("point[4] P"), RtPointArray(points)), ErrorShouldEqual, `PointsGeneralPolygons [2 2] [4 3 4 3] [0 1 4 3 6 7 8 1 2 5 4 9 10 11] "point[4] P" [0 1 0 0 1 1 0 0 1 0 0 0]`)

	//	So(ctx.PointsGeneralPolygonsV([]Rter{RtFloat(2), RtIntArray{2, 2}, RtIntArray{4, 3, 4, 3}, RtIntArray{0, 1, 4, 3, 6, 7, 8, 1, 2, 5, 4, 9, 10, 11}},
	//		[]Rter{RtToken("P")}, []Rter{RtPointArray(points)}), ErrorShouldEqual, `PointsGeneralPolygons 2 [2 2] [4 3 4 3] [0 1 4 3 6 7 8 1 2 5 4 9 10 11] "P" [0 1 0 0 1 1 0 0 1 0 0 0]`)

		So(ctx.End(), ErrorShouldEqual, `End`)
	})
}

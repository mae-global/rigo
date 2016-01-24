package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Quadrics(t *testing.T) {

	Convey("All Quadrics", t, func() {

		ctx := New(nil,nil)
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("output/quadrics.rib"), ShouldBeNil)
		So(ctx.Comment("output from rigo, quadrics_test.go"), ShouldBeNil)

		So(ctx.Sphere(0.5, 0.0, 0.5, 360.0), ShouldBeNil)
		So(ctx.Cone(0.5, 0.5, 270.0), ShouldBeNil)
		So(ctx.Cylinder(0.5, 0.2, 1, 360), ShouldBeNil)
		So(ctx.Hyperboloid(RtPoint{0, 0, 1}, RtPoint{1, 0, 1}, 0.4), ShouldBeNil)
		So(ctx.Disk(0.3, 0.4, 0.1), ShouldBeNil)
		So(ctx.Torus(1, .3, 60, 90, 360), ShouldBeNil)

		So(ctx.End(), ShouldBeNil)
	})
}

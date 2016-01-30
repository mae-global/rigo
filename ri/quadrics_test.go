package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Quadrics(t *testing.T) {

	Convey("All Quadrics", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("quadrics.rib"), ErrorShouldEqual, `Begin "quadrics.rib"`)
		So(ctx.Comment("output from rigo, quadrics_test.go"), ErrorShouldEqual, `# output from rigo, quadrics_test.go`)

		So(ctx.Sphere(0.5, 0.0, 0.5, 360.0), ErrorShouldEqual, `Sphere .5 0 .5 360`)
		So(ctx.Cone(0.5, 0.5, 270.0), ErrorShouldEqual, `Cone .5 .5 270`)
		So(ctx.Cylinder(0.5, 0.2, 1, 360), ErrorShouldEqual, `Cylinder .5 .2 1 360`)
		So(ctx.Hyperboloid(RtPoint{0, 0, 1}, RtPoint{1, 0, 1}, 0.4), ErrorShouldEqual, `Hyperboloid 0 0 1 1 0 1 .4`)
		So(ctx.Disk(0.3, 0.4, 0.1), ErrorShouldEqual, `Disk .3 .4 .1`)
		So(ctx.Torus(1, .3, 60, 90, 360), ErrorShouldEqual, `Torus 1 .3 60 90 360`)

		So(ctx.End(), ErrorShouldEqual, `End`)
	})
}

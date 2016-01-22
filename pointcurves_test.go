package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func Test_PointCurves(t *testing.T) {

	Convey("All Point and Curves",t,func() {

		ctx := New(nil)
		So(ctx,ShouldNotBeNil)

		So(ctx.Begin("output/pointcurves.rib"),ShouldBeNil)
		So(ctx.Comment("output from rigo, pointcurves_test.go"),ShouldBeNil)

		So(ctx.Points(5,RtToken("P"),RtFloatArray{.5,-.5,-.5,0,.5,0,.5,0},RtToken("width"),RtFloatArray{.1,.12,.05,.02}),ShouldBeNil)
		So(ctx.Curves("cubic",4,RtIntArray{4},"nonperiodic",RtToken("constantwidth"),RtFloat(0.1)),ShouldBeNil)

		So(ctx.End(),ShouldBeNil)
	})
}

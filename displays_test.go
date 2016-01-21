package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func Test_Display(t *testing.T) {

	Convey("All Display",t,func() {

		ctx := New(nil)
		So(ctx,ShouldNotBeNil)

		So(ctx.Begin("output/displays.rib"),ShouldBeNil)
		So(ctx.Comment("output from rigo, displays_test.go"),ShouldBeNil)
		So(ctx.PixelVariance(.01),ShouldBeNil)	
		So(ctx.PixelSamples(2,2),ShouldBeNil)
		So(ctx.PixelFilter(GaussianFilter,2,1),ShouldBeNil)
		So(ctx.Exposure(1.5,2.3),ShouldBeNil)


		So(ctx.End(),ShouldBeNil)
	})	



}		


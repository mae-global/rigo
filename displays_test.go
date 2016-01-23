package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Display(t *testing.T) {

	Convey("All Display", t, func() {

		ctx := New(nil)
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("output/displays.rib"), ShouldBeNil)
		So(ctx.Comment("output from rigo, displays_test.go"), ShouldBeNil)
		So(ctx.PixelVariance(.01), ShouldBeNil)
		So(ctx.PixelSamples(2, 2), ShouldBeNil)
		So(ctx.PixelFilter(GaussianFilter, 2, 1), ShouldBeNil)
		So(ctx.Exposure(1.5, 2.3), ShouldBeNil)
		So(ctx.Imager("cmyk", RtString("foo"), RtInt(45)), ShouldBeNil)
		So(ctx.Quantize(RGBA, 2048, -1024, 3071, 1.0), ShouldBeNil)
		So(ctx.Display("pixar0", "framebuffer", "rgba", RtString("origin"), RtIntArray{10, 10}), ShouldBeNil)
		So(ctx.Hider("hidden", RtString("samples"), RtInt(3), RtString("detail"), RtFloat(1.0)), ShouldBeNil)
		So(ctx.ColorSamples(RtInt(3), RtFloatArray{.3, .3, .4}, RtFloatArray{1, 1, 1}), ShouldBeNil)
		So(ctx.RelativeDetail(RtFloat(0.6)), ShouldBeNil)

		So(ctx.End(), ShouldBeNil)
	})

}

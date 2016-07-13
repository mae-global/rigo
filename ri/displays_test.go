package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Display(t *testing.T) {

	Convey("All Display", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("displays.rib"), ErrorShouldEqual, `Begin "displays.rib"`)
		So(ctx.Comment("output from rigo, displays_test.go"), ErrorShouldEqual, `# output from rigo, displays_test.go`)

		So(ctx.PixelVariance(.01), ErrorShouldEqual, `PixelVariance .01`)
		//So(ctx.PixelVarianceV([]Rter{RtFloat(.01)}, nil, nil), ErrorShouldEqual, `PixelVariance .01`)

		So(ctx.PixelSamples(2, 2), ErrorShouldEqual, `PixelSamples 2 2`)
		//So(ctx.PixelSamplesV([]Rter{RtFloat(2), RtFloat(2)}, nil, nil), ErrorShouldEqual, `PixelSamples 2 2`)

		So(ctx.PixelFilter(GaussianFilter, 2, 1), ErrorShouldEqual, `PixelFilter "gaussian" 2 1`)
		//So(ctx.PixelFilterV([]Rter{GaussianFilter, RtFloat(2), RtFloat(1)}, nil, nil), ErrorShouldEqual, `PixelFilter "gaussian" 2 1`)

		So(ctx.Exposure(1.5, 2.3), ErrorShouldEqual, `Exposure 1.5 2.3`)
		//So(ctx.ExposureV([]Rter{RtFloat(1.5), RtFloat(2.3)}, nil, nil), ErrorShouldEqual, `Exposure 1.5 2.3`)

		So(ctx.Imager("cmyk", RtToken("int foo"), RtInt(45)), ErrorShouldEqual, `Imager "cmyk" "int foo" [45]`)
		//So(ctx.ImagerV([]Rter{RtString("cmyk")}, []Rter{RtToken("foo")}, []Rter{RtInt(45)}), ErrorShouldEqual, `Imager "cmyk" "foo" [45]`)

		So(ctx.Quantize(RGBA, 2048, -1024, 3071, 1.0), ErrorShouldEqual, `Quantize "rgba" 2048 -1024 3071 1`)
		//So(ctx.QuantizeV([]Rter{RGBA, RtFloat(2048), RtFloat(-1024), RtFloat(3071), RtFloat(1.0)}, nil, nil), ErrorShouldEqual, `Quantize "rgba" 2048 -1024 3071 1`)

		So(ctx.Display("pixar0", "framebuffer", "rgba", RtToken("int[2] origin"), RtIntArray{10, 10}), ErrorShouldEqual, `Display "pixar0" "framebuffer" "rgba" "int[2] origin" [10 10]`)
		//So(ctx.DisplayV([]Rter{RtToken("pixar0"), RtToken("framebuffer"), RtToken("rgba")},
			//[]Rter{RtToken("origin")}, []Rter{RtIntArray{10, 10}}), ErrorShouldEqual, `Display "pixar0" "framebuffer" "rgba" "origin" [10 10]`)

		So(ctx.Hider("hidden", RtToken("int samples"), RtInt(3), RtToken("float detail"), RtFloat(1.0)), ErrorShouldEqual, `Hider "hidden" "int samples" [3] "float detail" [1]`)
		//So(ctx.HiderV([]Rter{RtToken("hidden")}, []Rter{RtToken("samples"), RtToken("detail")}, []Rter{RtInt(3), RtFloat(1.0)}), ErrorShouldEqual, `Hider "hidden" "samples" [3] "detail" [1]`)

		So(ctx.ColorSamples(RtInt(3), RtFloatArray{.3, .3, .4}, RtFloatArray{1, 1, 1}), ErrorShouldEqual, `ColorSamples 3 [.3 .3 .4] [1 1 1]`)
		//So(ctx.ColorSamplesV([]Rter{RtInt(3), RtFloatArray{.3, .3, .4}, RtFloatArray{1, 1, 1}}, nil, nil), ErrorShouldEqual, `ColorSamples 3 [.3 .3 .4] [1 1 1]`)

		So(ctx.RelativeDetail(RtFloat(0.6)), ErrorShouldEqual, `RelativeDetail .6`)
		//So(ctx.RelativeDetailV([]Rter{RtFloat(0.6)}, nil, nil), ErrorShouldEqual, `RelativeDetail .6`)

		So(ctx.End(), ErrorShouldEqual, `End`)
	})

}

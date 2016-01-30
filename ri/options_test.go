package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Options(t *testing.T) {

	Convey("All Options", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("options.rib"), ErrorShouldEqual,`Begin "options.rib"`)
		So(ctx.Comment("output from rigo, options_test.go"), ErrorShouldEqual,`# output from rigo, options_test.go`)
		So(ctx.Option(RtToken("foo"), RtToken("fov"), RtFloat(3.4), RtToken("color"), RtFloatArray{.4, .4, .4}, RtToken("normal"), RtIntArray{1, 1, 0}), ErrorShouldEqual,`Option "foo" "fov" 3.4 "color" [.4 .4 .4] "normal" [1 1 0]`)
		So(ctx.Attribute(RtToken("displacementbound"), RtToken("sphere"), RtFloat(2.0)), ErrorShouldEqual,`Attribute "displacementbound" "sphere" 2`)

		So(ctx.Geometry("teapot"), ErrorShouldEqual,`Geometry "teapot"`)

		So(ctx.MotionBegin(3, 0.1, 0.2, 0.3), ErrorShouldEqual,`MotionBegin [.1 .2 .3]`)
		So(ctx.MotionEnd(), ErrorShouldEqual,`MotionEnd`)

		So(ctx.End(), ErrorShouldEqual,`End`)
	})
}

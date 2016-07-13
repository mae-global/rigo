package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Options(t *testing.T) {

	Convey("All Options", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("options.rib"), ErrorShouldEqual, `Begin "options.rib"`)
		So(ctx.Comment("output from rigo, options_test.go"), ErrorShouldEqual, `# output from rigo, options_test.go`)
		So(ctx.Option("foo", RtToken("float fov"), RtFloat(3.4), RtToken("float[3] color"), RtFloatArray{.4, .4, .4}, RtToken("int[3] normal"), RtIntArray{1, 1, 0}), 
										ErrorShouldEqual, `Option "foo" "float fov" [3.4] "float[3] color" [.4 .4 .4] "int[3] normal" [1 1 0]`)

		So(ctx.Attribute("displacementbound", RtToken("float sphere"), RtFloat(2.0)), ErrorShouldEqual, `Attribute "displacementbound" "float sphere" [2]`)

		So(ctx.Geometry("teapot"), ErrorShouldEqual, `Geometry "teapot"`)

		So(ctx.MotionBegin(3, 0.1, 0.2, 0.3), ErrorShouldEqual, `MotionBegin [.1 .2 .3]`)
		So(ctx.MotionEnd(), ErrorShouldEqual, `MotionEnd`)

		So(ctx.End(), ErrorShouldEqual, `End`)
	})
}

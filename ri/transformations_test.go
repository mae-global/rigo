package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Transformations(t *testing.T) {

	Convey("All Transformations", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("transformations.rib"), ErrorShouldEqual,`Begin "transformations.rib"`)
		So(ctx.Comment("output from rigo, transformations_test.go"), ErrorShouldEqual,`# output from rigo, transformations_test.go`)
		So(ctx.TransformBegin(), ErrorShouldEqual,`TransformBegin`)

		So(ctx.Identity(), ErrorShouldEqual,`Identity`)
		So(ctx.Transform(RtMatrix{.5, .0, .0, .0, .0, 1., .0, .0, .0, .0, 1., .0, .0, .0, .0, 1.}), ErrorShouldEqual,`Transform [.5 0 0 0 0 1 0 0 0 0 1 0 0 0 0 1]`)
		So(ctx.ConcatTransform(RtMatrix{.5, .0, .0, .0, .0, 1., .0, .0, .0, .0, 1., .0, .0, .0, .0, 1.}), ErrorShouldEqual,`ConcatTransform [.5 0 0 0 0 1 0 0 0 0 1 0 0 0 0 1]`)
		So(ctx.Perspective(90.0), ErrorShouldEqual,`Perspective 90`)
		So(ctx.Translate(0.0, 1.0, 0.0), ErrorShouldEqual,`Translate 0 1 0`)
		So(ctx.Rotate(90.0, 0.0, 1.0, 0.0), ErrorShouldEqual,`Rotate 90 0 1 0`)
		So(ctx.Scale(.5, 1, 1), ErrorShouldEqual,`Scale .5 1 1`)
		So(ctx.Skew(45.0, 0, 1.0, 0, 1.0, 0, 0), ErrorShouldEqual,`Skew [45 0 1 0 1 0 0]`)
		So(ctx.CoordinateSystem("lamptop"), ErrorShouldEqual,`CoordinateSystem "lamptop"`)
		So(ctx.CoordSysTransform("lamptop"), ErrorShouldEqual,`CoordSysTransform "lamptop"`)

		So(ctx.TransformEnd(), ErrorShouldEqual,`TransformEnd`)
		So(ctx.End(), ErrorShouldEqual,`End`)
	})
}

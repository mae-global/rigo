package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func ErrorShouldEqual(actual interface{}, expected ...interface{}) string {
	var aerr error
	var berr string

	if a, ok := actual.(error); ok {
		aerr = a
	} else {
		return "Inputs should be errors"
	}

	if len(expected) == 0 {
		return "Expected a string to test against"
	}

	if b, ok := expected[0].(string); ok {
		berr = b
	} else {
		return "Inputs should be errors"
	}

	if aerr.Error() != berr {
		return "Actual: " + aerr.Error() + "\nExpecting: " + berr
	}

	return ""
}

func Test_Attributes(t *testing.T) {

	Convey("All Attributes", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("attributes.rib"), ErrorShouldEqual, `Begin "attributes.rib"`)
		So(ctx.Comment("output from rigo, attributes_test.go"), ErrorShouldEqual, `# output from rigo, attributes_test.go`)
		So(ctx.AttributeBegin(), ErrorShouldEqual, `AttributeBegin`)
		So(ctx.Color(RtColor{0.3, 0.4, 0.1, 0.2}), ErrorShouldEqual, `Color [.3 .4 .1 .2]`)
		So(ctx.Opacity(RtColor{1, 1, 1, 1}), ErrorShouldEqual, `Opacity [1 1 1 1]`)
		So(ctx.TextureCoordinates(0, 0, 2, -.5, -.5, 1.75, 3, 3), ErrorShouldEqual, `TextureCoordinates [0 0 2 -.5 -.5 1.75 3 3]`)

		spot, err := ctx.LightSource(RtToken("spotlight"), RtToken("coneangle"), RtInt(5))
		So(spot, ShouldEqual, "0")
		So(err, ErrorShouldEqual, `LightSource "spotlight" "0" "coneangle" [5]`)
		ambient, err := ctx.LightSource(RtToken("ambientlight"), RtToken("lightcolor"), RtColor{.5, 0, 0}, RtToken("intensity"), RtFloat(.6))
		So(err, ErrorShouldEqual, `LightSource "ambientlight" "1" "lightcolor" [.5 0 0] "intensity" [.6]`)
		So(ambient, ShouldEqual, "1")
		So(spot, ShouldNotEqual, ambient)

		So(ctx.Illuminate(spot, RtBoolean(true)), ErrorShouldEqual, `Illuminate "0" 1`)
		So(ctx.Illuminate(ambient, RtBoolean(false)), ErrorShouldEqual, `Illuminate "1" 0`)
		So(ctx.Surface("wood", RtToken("roughness"), RtFloat(0.3), RtToken("Kd"), RtFloat(1.0), RtToken("float ringwidth"), RtFloat(0.25)), ErrorShouldEqual, `Surface "wood" "roughness" [.3] "Kd" [1] "float ringwidth" [.25]`)
		So(ctx.Displacement("displaceit"), ErrorShouldEqual, `Displacement "displaceit"`)
		So(ctx.Atmosphere("fog"), ErrorShouldEqual, `Atmosphere "fog"`)
		So(ctx.Interior("water"), ErrorShouldEqual, `Interior "water"`)
		So(ctx.Exterior("fog"), ErrorShouldEqual, `Exterior "fog"`)

		So(ctx.ShadingRate(1.0), ErrorShouldEqual, `ShadingRate 1`)
		So(ctx.ShadingInterpolation("smooth"), ErrorShouldEqual, `ShadingInterpolation "smooth"`)
		So(ctx.Matte(true), ErrorShouldEqual, `Matte 1`)
		So(ctx.Bound(RtBound{0, .5, 0, .5, .9, 1}), ErrorShouldEqual, `Bound [0 .5 0 .5 .9 1]`)
		So(ctx.Detail(RtBound{10, 20, 42, 69, 0, 1}), ErrorShouldEqual, `Detail [10 20 42 69 0 1]`)
		So(ctx.DetailRange(0, 0, 10, 20), ErrorShouldEqual, `DetailRange [0 0 10 20]`)
		So(ctx.GeometricApproximation("flatness", 2.5), ErrorShouldEqual, `GeometricApproximation "flatness" 2.5`)
		So(ctx.Orientation("lh"), ErrorShouldEqual, `Orientation "lh"`)
		So(ctx.ReverseOrientation(), ErrorShouldEqual, `ReverseOrientation`)
		So(ctx.Sides(2), ErrorShouldEqual, `Sides 2`)

		So(ctx.AttributeEnd(), ErrorShouldEqual, `AttributeEnd`)
		So(ctx.End(), ErrorShouldEqual, `End`)
	})
}

func Test_AttributesExample(t *testing.T) {

	Convey("Attributes Example from pg. 74 of 'Advanced RenderMan'", t, func() {
		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		ctx.Begin("attributes.rib")
		ctx.AttributeBegin()
		ctx.Comment("Hemispherical wooden dome in a dense fog")
		ctx.Declare("density", "uniform float")
		ctx.Color(RtColor{.2, .45, .8})
		ctx.Opacity(RtColor{1, 1, 1})
		ctx.Sides(2)
		ctx.ShadingRate(1.0)
		ctx.Surface("paintedplastic", RtToken("texturename"), RtString("wood.tx"))
		ctx.Atmosphere("myfog", RtToken("density"), RtFloat(1.48))
		ctx.Sphere(1, 0, 1, 360)
		ctx.Disk(0, 1, 360)
		ctx.AttributeEnd()
		ctx.End()
	})
}

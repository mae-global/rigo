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
	//	So(ctx.AttributeBeginV(nil, nil, nil), ErrorShouldEqual, `AttributeBegin`)

		colour := RtColor{0.3, 0.4, 0.1, 0.2}

		So(ctx.Color(colour), ErrorShouldEqual, `Color [.3 .4 .1 .2]`)
		//So(ctx.ColorV([]Rter{colour}, nil, nil), ErrorShouldEqual, `Color [.3 .4 .1 .2]`)

		colour = RtColor{1, 1, 1, 1}

		So(ctx.Opacity(colour), ErrorShouldEqual, `Opacity [1 1 1 1]`)
		//So(ctx.OpacityV([]Rter{colour}, nil, nil), ErrorShouldEqual, `Opacity [1 1 1 1]`)

		So(ctx.TextureCoordinates(0, 0, 2, -.5, -.5, 1.75, 3, 3), ErrorShouldEqual, `TextureCoordinates [0 0 2 -.5 -.5 1.75 3 3]`)
	//	So(ctx.TextureCoordinatesV([]Rter{RtFloatArray{0, 0, 2, -.5, -.5, 1.73, 3, 3}}, nil, nil), ErrorShouldEqual, `TextureCoordinates [0 0 2 -.5 -.5 1.73 3 3]`)

		spot, err := ctx.LightSource("spotlight", RtToken("int coneangle"), RtInt(5))
		So(spot, ShouldEqual, "0")
		So(err, ErrorShouldEqual, `LightSource "spotlight" "0" "int coneangle" [5]`)
		ambient, err := ctx.LightSource("ambientlight", RtToken("color lightcolor"), RtColor{.5, 0, 0}, RtToken("float intensity"), RtFloat(.6))
		So(err, ErrorShouldEqual, `LightSource "ambientlight" "1" "color lightcolor" [.5 0 0] "float intensity" [.6]`)
		So(ambient, ShouldEqual, "1")
		So(spot, ShouldNotEqual, ambient)

		So(ctx.Illuminate(spot, RtBoolean(true)), ErrorShouldEqual, `Illuminate "0" 1`)
	//	So(ctx.IlluminateV([]Rter{spot, RtBoolean(true)}, nil, nil), ErrorShouldEqual, `Illuminate "0" 1`)

		So(ctx.Illuminate(ambient, RtBoolean(false)), ErrorShouldEqual, `Illuminate "1" 0`)

		So(ctx.Surface("wood", RtToken("float roughness"), RtFloat(0.3), RtToken("float Kd"), RtFloat(1.0), RtToken("float ringwidth"), RtFloat(0.25)), 
									ErrorShouldEqual, `Surface "wood" "float roughness" [.3] "float Kd" [1] "float ringwidth" [.25]`)
	//	So(ctx.SurfaceV([]Rter{RtString("wood")},
		//	[]Rter{RtToken("roughness"), RtToken("Kd"), RtToken("float ringwidth")},
			//[]Rter{RtFloat(0.3), RtFloat(1.0), RtFloat(0.25)}), ErrorShouldEqual, `Surface "wood" "roughness" [.3] "Kd" [1] "float ringwidth" [.25]`)

		So(ctx.Displacement("displaceit"), ErrorShouldEqual, `Displacement "displaceit"`)
//		So(ctx.DisplacementV([]Rter{RtString("displaceit")}, nil, nil), ErrorShouldEqual, `Displacement "displaceit"`)

		So(ctx.Atmosphere("fog"), ErrorShouldEqual, `Atmosphere "fog"`)
	//	So(ctx.AtmosphereV([]Rter{RtString("fog")}, nil, nil), ErrorShouldEqual, `Atmosphere "fog"`)

		So(ctx.Interior("water"), ErrorShouldEqual, `Interior "water"`)
	//	So(ctx.InteriorV([]Rter{RtString("water")}, nil, nil), ErrorShouldEqual, `Interior "water"`)

		So(ctx.Exterior("fog"), ErrorShouldEqual, `Exterior "fog"`)
		//So(ctx.ExteriorV([]Rter{RtString("fog")}, nil, nil), ErrorShouldEqual, `Exterior "fog"`)

		So(ctx.ShadingRate(1.0), ErrorShouldEqual, `ShadingRate 1`)
		//So(ctx.ShadingRateV([]Rter{RtFloat(1.0)}, nil, nil), ErrorShouldEqual, `ShadingRate 1`)

		So(ctx.ShadingInterpolation("smooth"), ErrorShouldEqual, `ShadingInterpolation "smooth"`)
	//	So(ctx.ShadingInterpolationV([]Rter{RtString("smooth")}, nil, nil), ErrorShouldEqual, `ShadingInterpolation "smooth"`)

		So(ctx.Matte(true), ErrorShouldEqual, `Matte 1`)
	//	So(ctx.MatteV([]Rter{RtBoolean(true)}, nil, nil), ErrorShouldEqual, `Matte 1`)

		So(ctx.Bound(RtBound{0, .5, 0, .5, .9, 1}), ErrorShouldEqual, `Bound [0 .5 0 .5 .9 1]`)
	//	So(ctx.BoundV([]Rter{RtBound{0, .5, 0, .5, .9, 1}}, nil, nil), ErrorShouldEqual, `Bound [0 .5 0 .5 .9 1]`)

		So(ctx.Detail(RtBound{10, 20, 42, 69, 0, 1}), ErrorShouldEqual, `Detail [10 20 42 69 0 1]`)
	//	So(ctx.DetailV([]Rter{RtBound{10, 20, 42, 69, 0, 1}}, nil, nil), ErrorShouldEqual, `Detail [10 20 42 69 0 1]`)

		So(ctx.DetailRange(0, 0, 10, 20), ErrorShouldEqual, `DetailRange [0 0 10 20]`)
	//	So(ctx.DetailRangeV([]Rter{RtFloatArray{0, 0, 10, 20}}, nil, nil), ErrorShouldEqual, `DetailRange [0 0 10 20]`)

		So(ctx.GeometricApproximation("flatness", 2.5), ErrorShouldEqual, `GeometricApproximation "flatness" 2.5`)
		//So(ctx.GeometricApproximationV([]Rter{RtString("flatness"), RtFloat(2.5)}, nil, nil), ErrorShouldEqual, `GeometricApproximation "flatness" 2.5`)

		So(ctx.Orientation("lh"), ErrorShouldEqual, `Orientation "lh"`)
		//So(ctx.OrientationV([]Rter{RtString("lh")}, nil, nil), ErrorShouldEqual, `Orientation "lh"`)

		So(ctx.ReverseOrientation(), ErrorShouldEqual, `ReverseOrientation`)
		//So(ctx.ReverseOrientationV(nil, nil, nil), ErrorShouldEqual, `ReverseOrientation`)

		So(ctx.Sides(2), ErrorShouldEqual, `Sides 2`)
		//So(ctx.SidesV([]Rter{RtFloat(2)}, nil, nil), ErrorShouldEqual, `Sides 2`)

		So(ctx.AttributeEnd(), ErrorShouldEqual, `AttributeEnd`)
		//So(ctx.AttributeEndV(nil, nil, nil), ErrorShouldEqual, `AttributeEnd`)

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

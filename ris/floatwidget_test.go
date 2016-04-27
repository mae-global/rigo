package ris

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	. "github.com/mae-global/rigo/ri"
)

func ShouldEqualFloat(actual interface{}, expected ...interface{}) string {

	a := actual.(Rter)
	b := expected[0].(Rter)

	ac, ok := a.(RtFloat)
	if !ok {
		return "Expected: RtFloat\nActual: "
	}
	bc, ok := b.(RtFloat)
	if !ok {
		return "Expected: RtFloat for comparison"
	}

	if ac == bc {
		return ""
	}

	return fmt.Sprintf("Expected: %v\nActual: %v", ac, bc)
}

func Test_FloatWidget(t *testing.T) {

	Convey("FloatWidget", t, func() {

		shader, err := Parse("TestFloatWidget", "-", []byte(floatwidget))
		So(err, ShouldBeNil)
		So(shader, ShouldNotBeNil)

		fw, ok := shader.Widget("intensity").(*RtFloatWidget)
		So(ok, ShouldBeTrue)
		So(fw, ShouldNotBeNil)

		So(fw.Name(), ShouldEqual, "intensity")
		So(fw.NameSpec(), ShouldEqual, "float intensity")
		So(fw.Label(), ShouldEqual, "Intensity")
		So(fw.Help(), ShouldBeEmpty)
		So(fw.Value(), ShouldEqualFloat, RtFloat(0.9))

		min, max := fw.Bounds()
		So(min, ShouldEqualFloat, RtFloat(0.0))
		So(max, ShouldEqualFloat, RtFloat(1.0))

		So(fw.Set(RtFloat(0.5)), ShouldBeNil)
	})
}

const floatwidget = `
<args format="1.0">
	<shaderType>
		<tag value="bxdf" />
	</shaderType>
	<param name="intensity" type="float" label="Intensity" default="0.9" widget="float" min="0.0" max="1.0">
		<tags>
			<tag value="float" />
		</tags>
	</param>
	<rfmdata node="101" classification="shader/surface:rendernode/RenderMan/bxdf:swatch/rmanSwatch" />
</args>
`

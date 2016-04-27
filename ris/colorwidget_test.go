package ris

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	. "github.com/mae-global/rigo/ri"
)

func ShouldEqualColor(actual interface{}, expected ...interface{}) string {

	a := actual.(Rter)
	b := expected[0].(Rter)

	ac, ok := a.(RtColor)
	if !ok {
		return "Expected: RtColor\nActual: "
	}
	bc, ok := b.(RtColor)
	if !ok {
		return "Expected RtColor for comparison"
	}

	if ac.Equal(bc) {
		return ""
	}

	return fmt.Sprintf("Expected: %v\nActual: %v", ac, bc)
}

func Test_ColorWidget(t *testing.T) {

	Convey("ColorWidget", t, func() {

		shader, err := Parse("TestColorWidget", "-", []byte(colorwidget))
		So(err, ShouldBeNil)
		So(shader, ShouldNotBeNil)

		cw, ok := shader.Widget("color").(*RtColorWidget)
		So(ok, ShouldBeTrue)
		So(cw, ShouldNotBeNil)

		So(cw.Name(), ShouldEqual, "color")
		So(cw.NameSpec(), ShouldEqual, "color color")
		So(cw.Label(), ShouldEqual, "Color")
		So(cw.Help(), ShouldBeEmpty)
		So(cw.Value(), ShouldEqualColor, RtColor{.1, .2, .3})

		min, max := cw.Bounds()
		So(min, ShouldEqualColor, RtColor{0, 0, 0})
		So(max, ShouldEqualColor, RtColor{1, 1, 1})

		So(cw.Set(RtColor{.45, .45, .45}), ShouldBeNil)
	})
}

const colorwidget = `
<args format="1.0">
	<shaderType>
		<tag value="bxdf" />
	</shaderType>
	<param name="color" type="color" label="Color" default=".1 .2 .3" widget="color" min="0 0 0" max="1 1 1">
		<tags>
			<tag value="color" />
		</tags>
	</param>
	<rfmdata node="101" classification="shader/surface:rendernode/RenderMan/bxdf:swatch/rmanSwatch" />
</args>
`

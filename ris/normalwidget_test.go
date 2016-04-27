package ris

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	. "github.com/mae-global/rigo/ri"
)

func ShouldEqualNormal(actual interface{}, expected ...interface{}) string {

	a := actual.(Rter)
	b := expected[0].(Rter)

	ac, ok := a.(RtNormal)
	if !ok {
		return "Expected: RtNormal\nActual: "
	}
	bc, ok := b.(RtNormal)
	if !ok {
		return "Expected RtNormal for comparison"
	}

	if ac.Equal(bc) {
		return ""
	}

	return fmt.Sprintf("Expected: %v\nActual: %v", ac, bc)
}

func Test_NormalWidget(t *testing.T) {

	Convey("NormalWidget", t, func() {

		shader, err := Parse("TestNormalWidget", "-", []byte(normalwidget))
		So(err, ShouldBeNil)
		So(shader, ShouldNotBeNil)

		cw, ok := shader.Widget("surface").(*RtNormalWidget)
		So(ok, ShouldBeTrue)
		So(cw, ShouldNotBeNil)

		So(cw.Name(), ShouldEqual, "surface")
		So(cw.NameSpec(), ShouldEqual, "normal surface")
		So(cw.Label(), ShouldEqual, "Surface")
		So(cw.Help(), ShouldBeEmpty)
		So(cw.Value(), ShouldEqualNormal, RtNormal{.1, .2, .3})

		min, max := cw.Bounds()
		So(min, ShouldEqualNormal, RtNormal{0, 0, 0})
		So(max, ShouldEqualNormal, RtNormal{1, 1, 1})

		So(cw.Set(RtNormal{.45, .45, .45}), ShouldBeNil)
	})
}

const normalwidget = `
<args format="1.0">
	<shaderType>
		<tag value="bxdf" />
	</shaderType>
	<param name="surface" type="normal" label="Surface" default=".1 .2 .3" widget="normal" min="0 0 0" max="1 1 1">
		<tags>
			<tag value="normal" />
		</tags>
	</param>
	<rfmdata node="101" classification="shader/surface:rendernode/RenderMan/bxdf:swatch/rmanSwatch" />
</args>
`

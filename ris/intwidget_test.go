package ris

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"

	. "github.com/mae-global/rigo/ri"
)

func ShouldEqualInt(actual interface{},expected ...interface{}) string {

	a := actual.(Rter)
	b := expected[0].(Rter)

	ac,ok := a.(RtInt)
	if !ok {
		return "Expected: RtInt\nActual: "
	}
	bc,ok := b.(RtInt)
	if !ok {
		return "Expected: RtInt for comparison"
	}

	if ac == bc {
		return ""
	}

	return fmt.Sprintf("Expected: %v\nActual: %v",ac,bc)
}

func Test_IntWidget(t *testing.T) {

	Convey("IntWidget",t,func() {
		
		shader,err := Parse("TestIntWidget","-",[]byte(intwidget))
		So(err,ShouldBeNil)
		So(shader,ShouldNotBeNil)

		iw,ok := shader.Widget("fov").(*RtIntWidget)
		So(ok,ShouldBeTrue)
		So(iw,ShouldNotBeNil)

		So(iw.Name(),ShouldEqual,"fov")
		So(iw.NameSpec(),ShouldEqual,"int fov")
		So(iw.Label(),ShouldEqual,"Fov")
		So(iw.Help(),ShouldBeEmpty)
		So(iw.Value(),ShouldEqualInt,RtInt(30))

		min,max := iw.Bounds()
		So(min,ShouldEqualInt,RtInt(0))
		So(max,ShouldEqualInt,RtInt(100))
		
		So(iw.Set(RtInt(45)),ShouldBeNil)
	})
}

const intwidget = `
<args format="1.0">
	<shaderType>
		<tag value="bxdf" />
	</shaderType>
	<param name="fov" type="int" label="Fov" default="30" widget="int" min="0" max="100">
		<tags>
			<tag value="int" />
		</tags>
	</param>
	<rfmdata node="101" classification="shader/surface:rendernode/RenderMan/bxdf:swatch/rmanSwatch" />
</args>
`


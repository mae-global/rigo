package ri

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"

)

func Test_ParseString(t *testing.T) {

	Convey("Parse RIB string",t,func() {

		err := ParseString(RIBExample0,nil)
		So(err,ShouldBeNil)

	})
}

const RIBExample0 = `##RenderMan RIB-Structure 1.1
version 3.04
Projection "perspective" "fov" 30.0
Color [1 0 0]
Sphere 1 -1 1 360
`

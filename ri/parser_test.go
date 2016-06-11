package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"strings"
)

func Test_ParseString(t *testing.T) {

	Convey("Parse RIB string", t, func() {

		err := parse(strings.NewReader(RIBExample0), nil)
		So(err, ShouldBeNil)

	})
}

const RIBExample0 = `##RenderMan RIB-Structure 1.1
version 3.04
Declare "d" "uniform point"
Option "limits" "bucketsize" [6 6]
Projection "perspective" "float fov" [30.0]
Color [1 0 0]
AttributeBegin
	Attribute "identifier" "shadinggroup" ["tophalf"]
	Surface "plastic" 
	Sphere 1 -1 1 360
AttributeEnd
`

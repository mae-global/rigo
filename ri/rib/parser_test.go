/* rigo/ri/parser/parser_test.go */
package rib

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_SimpleParse(t *testing.T) {

	Convey("Simple Parser",t,func() {
		
		So(len(RIBExample1),ShouldEqual,169)
		So(Parse([]byte(RIBExample1)),ShouldBeNil)
		
	})
}	

const RIBExample1 = `##RenderMan RIB-Structure 1.1
version 3.04 
Display "sphere.tif" "file" "rgb" 
Format 320 240 1 
Translate 0 0 6 
WorldBegin  
Color [1 0 0] 
Sphere 1 -1 1 360 
WorldEnd`
	

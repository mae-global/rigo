/* rigo/handles/handles_test.go */
package handles

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"	
	"fmt"
)

func Test_NumberGenerators(t *testing.T) {

	Convey("Light Number Generator",t,func() {
	
		gen := new(LightNumberGenerator)
		So(gen,ShouldNotBeNil)

		h,err := gen.Generate()
		So(err,ShouldBeNil)
		So(h.String(),ShouldEqual,"\"0\"")

		for i := 1; i < 10; i++ {
			h,err := gen.Generate()
			So(err,ShouldBeNil)
			So(h.String(),ShouldEqual,fmt.Sprintf("\"%d\"",i))
		}
	})

	Convey("Object Number Generator",t,func() {

		gen := new(ObjectNumberGenerator)
		So(gen,ShouldNotBeNil)
		
		h,err := gen.Generate()
		So(err,ShouldBeNil)
		So(h.String(),ShouldEqual,"\"0\"")

		for i := 1; i < 10; i++ {
			h,err := gen.Generate()
			So(err,ShouldBeNil)
			So(h.String(),ShouldEqual,fmt.Sprintf("\"%d\"",i))
		}
	})
}

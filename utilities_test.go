package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
)


func Test_Utilities(t *testing.T) {

	Convey("Utilities",t,func() {

		Convey("check for mismatch error",func() {
			out,err := serialise(RtInt(1),RtInt(2),RtInt(3))
			So(err,ShouldEqual,ErrBadParamlist)
			So(out,ShouldBeEmpty)
		})

		Convey("check correct serialisation",func() {
			out,err := serialise(RtInt(1),RtInt(2),RtInt(3),RtInt(4))
			So(err,ShouldBeNil)
			So(out,ShouldEqual,"1 2 3 4")
		})

		Convey("check annotation parsing",func() {
			out := parseAnnotations(RtAnnotation("hello"),RtAnnotation("there"),RtAnnotation("Alice"))
			So(len(out),ShouldEqual,3)
			So(serialiseToString(out...),ShouldEqual,"#hello there Alice")
		})

	})

	list := []Rter{RtToken("test"),RtFloat(1.2),RtFloatArray{.1,.2,.3,.4}}
	if t,ok := list[0].(RtToken); ok {
		fmt.Printf("token %s\n",t)
	}
	if t,ok := list[1].(RtFloat); ok {
		fmt.Printf("float %s\n",t)
	}
	if t,ok := list[2].(RtFloatArray); ok {
		fmt.Printf("float array len=%d %s\n",len(t),t)
	}


}

package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func Test_Utilities(t *testing.T) {

	Convey("Utilities",t,func() {

		Convey("check for mismatch error",func() {
			out,err := serialise(RtInt(1),RtInt(2),RtInt(3))
			So(err,ShouldEqual,ErrParameterlistMismatch)
			So(out,ShouldBeEmpty)
		})

		Convey("check corect serialisation",func() {
			out,err := serialise(RtInt(1),RtInt(2),RtInt(3),RtInt(4))
			So(err,ShouldBeNil)
			So(out,ShouldEqual,"1 2 3 4")
		})
	})
}

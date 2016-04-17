package ri

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Utilities(t *testing.T) {

	Convey("Utilities", t, func() {

		Convey("check for mismatch error", func() {
			out, err := serialise(RtInt(1), RtInt(2), RtInt(3))
			So(err, ShouldEqual, ErrBadParamlist)
			So(out, ShouldBeEmpty)
		})

		Convey("check correct serialisation", func() {
			out, err := serialise(RtInt(1), RtInt(2), RtInt(3), RtInt(4))
			So(err, ShouldBeNil)
			So(out, ShouldEqual, "1 2 3 4")
		})

		Convey("check annotation parsing", func() {
			out := parseAnnotations(RtAnnotation("hello"), RtAnnotation("there"), RtAnnotation("Alice"))
			So(len(out), ShouldEqual, 3)
			So(serialiseToString(out...), ShouldEqual, "#hello there Alice")
		})

		Convey("reduce", func() {
			So(reduce(0.5), ShouldEqual, ".5")
			So(reduce(-0.5), ShouldEqual, "-.5")
			So(reduce(5.00010), ShouldEqual, "5.0001")
			So(reduce(0.0), ShouldEqual, "0")
		})

		Convey("reducev", func() {
			So(reducev([]RtFloat{0.05, 1.0500, 0.0, -.1}), ShouldEqual, ".05 1.05 0 -.1")
		})

		Convey("ClassTypeNameCount",func() {
			c,t,n,count := ClassTypeNameCount("varying float[2] st")
			So(c,ShouldEqual,RtToken("varying"))
			So(t,ShouldEqual,RtToken("float"))
			So(n,ShouldEqual,ST)
			So(count,ShouldEqual,2)
		})

	})

	list := []Rter{RtToken("test"), RtFloat(1.2), RtFloatArray{.1, .2, .3, .4}}
	if t, ok := list[0].(RtToken); ok {
		fmt.Printf("token %s\n", t)
	}
	if t, ok := list[1].(RtFloat); ok {
		fmt.Printf("float %s\n", t)
	}
	if t, ok := list[2].(RtFloatArray); ok {
		fmt.Printf("float array len=%d %s\n", len(t), t)
	}

}

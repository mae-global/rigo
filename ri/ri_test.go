package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Ri(t *testing.T) {

	Convey("Ri", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("ri.rib"), ErrorShouldEqual, `Begin "ri.rib"`)
		So(ctx.Surface("wood", RtToken("roughness"), RtToken("in error"), RtInt(5)), ShouldEqual, ErrBadArgument)
		So(ctx.Surface("wood", RtToken("roughness"), RtToken("in error"), RtToken("of")), ShouldEqual, ErrBadArgument)
	})
}

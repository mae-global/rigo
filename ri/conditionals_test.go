package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Conditionals(t *testing.T) {

	Convey("All Conditionals", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		So(ctx.IfBegin("$user:renderpass == 'shadow'"), ErrorShouldEqual, `IfBegin "$user:renderpass == 'shadow'"`)
		So(ctx.IfBeginV([]Rter{RtString("$user:renderpass == 'shadow'")},nil,nil), ErrorShouldEqual, `IfBegin "$user:renderpass == 'shadow'"`)

		So(ctx.ElseIf("$user:renderpass == 'beauty'"), ErrorShouldEqual, `ElseIf "$user:renderpass == 'beauty'"`)
		So(ctx.ElseIfV([]Rter{RtString("$user:renderpass == 'beauty'")},nil,nil), ErrorShouldEqual, `ElseIf "$user:renderpass == 'beauty'"`)

		So(ctx.Else(), ErrorShouldEqual, `Else`)
		So(ctx.ElseV(nil,nil,nil), ErrorShouldEqual, `Else`)

		So(ctx.IfEnd(), ErrorShouldEqual, `IfEnd`)
		So(ctx.IfEndV(nil,nil,nil), ErrorShouldEqual, `IfEnd`)
	})
}

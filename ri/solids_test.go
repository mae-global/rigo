package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Solids(t *testing.T) {

	Convey("All Solids", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("solids.rib"), ErrorShouldEqual,`Begin "solids.rib"`)
		So(ctx.Comment("output from rigo, solids_test.go"), ErrorShouldEqual,`# output from rigo, solids_test.go`)
		So(ctx.SolidBegin("union"), ErrorShouldEqual,`SolidBegin "union"`)

		oh, err := ctx.ObjectBegin()
		So(err, ErrorShouldEqual,`ObjectBegin 0`)
		So(ctx.ObjectEnd(), ErrorShouldEqual,`ObjectEnd`)
		So(ctx.ObjectInstance(oh), ErrorShouldEqual,`ObjectInstance 0`)

		So(ctx.SolidEnd(), ErrorShouldEqual,`SolidEnd`)
		So(ctx.End(), ErrorShouldEqual,`End`)
	})
}

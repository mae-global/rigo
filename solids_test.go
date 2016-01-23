package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Solids(t *testing.T) {

	Convey("All Solids", t, func() {

		ctx := New(nil)
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("output/solids.rib"), ShouldBeNil)
		So(ctx.Comment("output from rigo, solids_test.go"), ShouldBeNil)
		So(ctx.SolidBegin("union"), ShouldBeNil)

		oh, err := ctx.ObjectBegin()
		So(err, ShouldBeNil)
		So(ctx.ObjectEnd(), ShouldBeNil)
		So(ctx.ObjectInstance(oh), ShouldBeNil)

		So(ctx.SolidEnd(), ShouldBeNil)
		So(ctx.End(), ShouldBeNil)
	})
}

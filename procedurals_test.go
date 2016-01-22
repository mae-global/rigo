package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func Test_Procedurals(t *testing.T) {

	Convey("All Procedurals",t,func() {

		ctx := New(nil)
		So(ctx,ShouldNotBeNil)

		So(ctx.Begin("output/procedurals.rib"),ShouldBeNil)
		So(ctx.Comment("output from rigo, procedurals_test.go"),ShouldBeNil)

		So(ctx.Procedural(RtStringArray{"sodacan.rib"},RtBound{-1,1,-1,1,0,6},ProcDelayedReadArchive,ProcFree),ShouldBeNil)

		So(ctx.End(),ShouldBeNil)
	})
}

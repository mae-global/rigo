package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Procedurals(t *testing.T) {

	Convey("All Procedurals", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("procedurals.rib"), ErrorShouldEqual,`Begin "procedurals.rib"`)
		So(ctx.Comment("output from rigo, procedurals_test.go"), ErrorShouldEqual,`# output from rigo, procedurals_test.go`)

		So(ctx.Procedural(RtStringArray{"sodacan.rib"}, RtBound{-1, 1, -1, 1, 0, 6}, ProcDelayedReadArchive, ProcFree), ErrorShouldEqual,`Procedural "DelayedReadArchive" ["sodacan.rib"] [-1 1 -1 1 0 6]`)

		So(ctx.End(), ErrorShouldEqual,`End`)	
	})
}

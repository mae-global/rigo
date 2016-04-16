package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"fmt"
)

func Test_Archive(t *testing.T) {

	Convey("All Archive", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		aw,err := ctx.ArchiveBegin("test")
		fmt.Printf("err = [%s]\n",err)
		So(err,ErrorShouldEqual, `ArchiveBegin "test"`)
		So(aw,ShouldNotBeNil)

		n,err := aw.Write([]byte("hello there"))
		So(err,ShouldBeNil)
		So(n,ShouldEqual,11)
			
		So(ctx.ArchiveEnd("test"),ErrorShouldEqual,`ArchiveEnd "test"`)
		So(ctx.ArchiveInstance("test"),ErrorShouldEqual,`ArchiveInstance "test"`)
	})
}

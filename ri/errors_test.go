package ri

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrorHandlers(t *testing.T) {

	Convey("Error Handling",t,func() {

		ctx := NewTest()
		So(ctx,ShouldNotBeNil)

		So(ctx.ErrorHandler(ErrorPrint),ErrorShouldEqual,`ErrorHandler "print"`)
		So(ctx.ErrorHandler(ErrorIgnore),ErrorShouldEqual,`ErrorHandler "ignore"`)
		So(ctx.ErrorHandler(ErrorAbort),ErrorShouldEqual,`ErrorHandler "abort"`)

	})
}

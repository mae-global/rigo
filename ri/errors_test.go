package ri

import (
	"fmt"	
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	__hiddenPanic int
)

func testpanic(str string) {
	__hiddenPanic ++
}

func resetpanic() {
	__hiddenpanic --
}


type exampleErrorHandler int

func (e exampleErrorHandler) Name() RtErrorHandler { return "abort" }
func (e exampleErrorHandler) Handle(err RtError) error {
	fmt.Printf("example %s\n",err)
	testpanic(err.Msg)
	return nil	
}


func Test_ErrorHandlers(t *testing.T) {

	Convey("Error Handling",t,func() {

		ctx := NewTest()
		So(ctx,ShouldNotBeNil)

		So(ctx.ErrorHandler(ErrorPrint(0)),ErrorShouldEqual,`ErrorHandler "print"`)
		So(ctx.ErrorHandler(ErrorIgnore(0)),ErrorShouldEqual,`ErrorHandler "ignore"`)
		So(ctx.ErrorHandler(ErrorAbort(0)),ErrorShouldEqual,`ErrorHandler "abort"`)

	})

	Convey("Error Test",t,func() {

		ctx := NewTest()
		So(ctx,ShouldNotBeNil)

		ctx.ErrorHandler(exampleErrorHandler(0))
		
		ctx.Attribute("test",RtToken("constant foo bar"),RtInt(1))

		So(__hiddenPanic,ShouldEqual,1); resetpanic()

	})			

}

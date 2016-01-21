/* rigo/state_test.go */
package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"bytes"
)

func Test_State(t *testing.T) {

	Convey("Begin",t,func() {
		Convey("twice with same name",func() {

			So(Begin("test.rib"),ShouldBeNil)
			So(Begin("test.rib"),ShouldEqual,ErrContextAlreadyExists)			
		})
	})

	Convey("End",t,func() {
		Convey("end twice",func() {
			
			So(End(),ShouldBeNil)
			So(End(),ShouldEqual,ErrNoActiveContext)
		})
	})

	Convey("BeginWriter/end",t,func() {
		buf := bytes.NewBuffer(nil)
		So(BeginWriter("test.rib",buf),ShouldBeNil)
		So(End(),ShouldBeNil)
		So(len(buf.Bytes()),ShouldNotEqual,0)	
	})	


	Convey("GetContext",t,func() {
		Convey("get the active context",func() {

			So(GetContext(),ShouldBeNil)
		})

		Convey("get valid context",func() {
			
			So(Begin("test.rib"),ShouldBeNil)
			ctx := GetContext()
			So(ctx,ShouldNotBeNil)
			
			So(Context(ctx),ShouldBeNil)
		})
	})
}		


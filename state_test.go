/* rigo/state_test.go */
package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"bytes"
	"os"
)

type buffered struct {
	buf *bytes.Buffer
}

func (buf *buffered) Bytes() []byte {
	return buf.buf.Bytes()
}

func (buf *buffered) Close() error {
	return nil
}

func (buf *buffered) Write(content []byte) (int,error) {
	return buf.buf.Write(content)
}

func newbuffered() *buffered {
	return &buffered{bytes.NewBuffer(nil)}
}


func Test_State(t *testing.T) {

	Convey("Context",t,func() {

		ctx := New(nil)
		So(ctx,ShouldNotBeNil)

		/* try a test.rib file */
		So(ctx.Begin("test.rib"),ShouldBeNil)
		So(ctx.End(),ShouldBeNil)

		So(ctx.End(),ShouldEqual,ErrNoActiveContext)

		info,err := os.Stat("test.rib")
		So(err,ShouldBeNil)
		So(info.IsDir(),ShouldBeFalse)

		os.Remove("test.rib")
	})

	Convey("Context - buffered file",t,func() {

		buf := newbuffered()
		ctx := New(buf)
		So(ctx,ShouldNotBeNil)

		So(ctx.Begin("test.rib"),ShouldBeNil)
		So(ctx.FrameBegin(1),ShouldBeNil)
		So(ctx.FrameEnd(),ShouldBeNil)
		So(ctx.End(),ShouldBeNil)

		So(len(buf.Bytes()),ShouldNotEqual,0)

	})

}		


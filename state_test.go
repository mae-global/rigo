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

func (buf *buffered) Write(depth int,content string) error {
	_,err := buf.buf.Write([]byte(content))
	return err
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

	Convey("All",t,func() {

		ctx := New(nil)
		So(ctx,ShouldNotBeNil)

		So(ctx.Begin("output/states.rib"),ShouldBeNil)
		So(ctx.Comment("output from rigo, state_test.go"),ShouldBeNil)
		So(ctx.FrameBegin(1),ShouldBeNil)
		So(ctx.Comment("random comment"),ShouldBeNil)
		So(ctx.Projection(Perspective,RtString("fov"),RtFloat(45.3)),ShouldBeNil)
		So(ctx.Clipping(0.1,10000),ShouldBeNil)
		So(ctx.ClippingPlane(3,0,0,0,0,-1),ShouldBeNil)
		So(ctx.DepthOfField(22,45,1200),ShouldBeNil)
		So(ctx.Shutter(0.1,0.9),ShouldBeNil)
		So(ctx.FrameEnd(),ShouldBeNil)
		So(ctx.End(),ShouldBeNil)
	})	



}		


/* rigo/state_test.go */
package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_State(t *testing.T) {

	Convey("Context", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		/* try a test.rib file */
		So(ctx.Begin("test.rib"), ErrorShouldEqual, `Begin "test.rib"`)
		So(ctx.End(), ErrorShouldEqual, `End`)
	})

	Convey("All", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("states.rib"), ErrorShouldEqual, `Begin "states.rib"`)
		So(ctx.Comment("output from rigo, state_test.go"), ErrorShouldEqual, `# output from rigo, state_test.go`)
		So(ctx.FrameBegin(1), ErrorShouldEqual, `FrameBegin 1`)
		So(ctx.Comment("random comment"), ErrorShouldEqual, `# random comment`)
		So(ctx.Projection(PERSPECTIVE, RtToken("float fov"), RtFloat(45.3)), ErrorShouldEqual, `Projection "perspective" "float fov" [45.3]`)
		So(ctx.Clipping(0.1, 10000), ErrorShouldEqual, `Clipping .1 10000`)
		So(ctx.ClippingPlane(3, 0, 0, 0, 0, -1), ErrorShouldEqual, `ClippingPlane 3 0 0 0 0 -1`)
		So(ctx.DepthOfField(22, 45, 1200), ErrorShouldEqual, `DepthOfField 22 45 1200`)
		So(ctx.Shutter(0.1, 0.9), ErrorShouldEqual, `Shutter .1 .9`)
		So(ctx.FrameEnd(), ErrorShouldEqual, `FrameEnd`)
		So(ctx.End(), ErrorShouldEqual, `End`)
	})

}

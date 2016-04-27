/* rigo/presets_test.go */
package rigo

import (
	"fmt"
	"testing"

	. "github.com/mae-global/rigo/ri"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Presets(t *testing.T) {

	Convey("Presets", t, func() {

		Convey("Strict Pipeline -- no extentions", func() {

			ctx, pipe := StrictPipeline()
			So(ctx, ShouldNotBeNil)

			ctx.Begin("output/strict/simple.rib")
			ctx.Display("sphere.tif", "file", "rgb")
			ctx.Format(320, 240, 1)
			ctx.Projection(PERSPECTIVE, RtString("fov"), RtFloat(30))
			ctx.Translate(0, 0, 6)
			ctx.WorldBegin()
			ctx.LightSource("ambientlight", RtString("intensity"), RtFloat(0.5))
			ctx.LightSource("distantlight", RtString("intensity"), RtFloat(1.2), RtString("form"), RtIntArray{0, 0, -6}, RtString("to"), RtIntArray{0, 0, 0})
			ctx.Color(RtColor{1, 0, 0})
			ctx.Sphere(1, -1, 1, 360)
			ctx.WorldEnd()

			So(ctx.End(), ShouldBeNil)

			/* output gathered stats */
			p := pipe.GetByName(PipeToStats{}.Name())
			So(p, ShouldNotBeNil)
			s, ok := p.(*PipeToStats)
			So(s, ShouldNotBeNil)
			So(ok, ShouldBeTrue)

			p = pipe.GetByName(PipeTimer{}.Name())
			So(p, ShouldNotBeNil)
			t, ok := p.(*PipeTimer)
			So(t, ShouldNotBeNil)
			So(ok, ShouldBeTrue)

			fmt.Printf("%s%s", s, t)
		})
	})
}

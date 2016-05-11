package rigo

import (
	"fmt"
	"os/user"
	"testing"
	"time"

	. "github.com/mae-global/rigo/ri"
	. "github.com/mae-global/rigo/ris"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_RISExamples(t *testing.T) {

	Convey("RIS Examples", t, func() {
		Convey("PxrConstant -- bxdf shader", func() {

			cuser, err := user.Current()
			So(err, ShouldBeNil)

			pipe := DefaultFilePipe()
			So(pipe, ShouldNotBeNil)

			ctx := NewContext(pipe, nil, &Configuration{PrettyPrint: true})
			ri := RI(ctx)

			ri.Begin("output/risBxdfPxrConstant.rib")
			ri.ArchiveRecord("structure", "Scene Bxdf Constant")
			ri.ArchiveRecord("structure", "Creator %s", Author)
			ri.ArchiveRecord("structure", "CreationDate %s", time.Now())
			ri.ArchiveRecord("structure", "For %s", cuser.Username)
			ri.ArchiveRecord("structure", "Frames 1")

			ri.Display("risbxdf_sphere.tif", "file", "rgb")
			ri.Format(320, 240, 1)
			ri.PixelFilter(GaussianFilter, 4, 4)
			ri.Imager("background", RtToken("color color"), RtColor{.6, .6, .6}, RtToken("float alpha"), RtFloat(1))

			ri.Projection(PERSPECTIVE, RtToken("fov"), RtFloat(30))
			ri.Translate(0, 0, 6)
			ri.WorldBegin()
			ri.Color(RtColor{1, 0, 0})

			/* wrap the context with the RIS interface */
			ris := RIS(ctx)
			/* load the PxrConstant bxdf shader, $RMANTREE/lib/RIS/bxdf/Args/PxrConstant.args is parsed
			 * for the constant shader.
			 */
			constant, err := ris.Bxdf("PxrConstant", "-")
			So(err, ShouldBeNil)
			So(constant, ShouldNotBeNil)

			/* we want to manipulate the emit color of the shader so we
			 * ask the shader for the widget interface
			 */
			w := constant.Widget("emitColor")
			So(w, ShouldNotBeNil)

			/* set the color via the widget; SetValue(Rter) */
			So(w.SetValue(RtColor{0.3, 0.2, 0.1}), ShouldBeNil)

			/* we know it's a color so we can convert directly to the
			 * RtColorWidget concrete type
			 */
			cw, ok := w.(*RtColorWidget)
			So(ok, ShouldBeTrue)
			So(cw, ShouldNotBeNil)

			/* we can directly set the color via our RtColorWidget :
			 * overriding our previous SetValue
			 */
			So(cw.Set(RtColor{0.1, 0.2, 0.3}), ShouldBeNil)

			/* As the widget (w & cw) is linked to the shader (constant), we can override
			 * what we set with the widget directly at the shader */
			So(constant.SetValue(RtToken("emitColor"), RtColor{0.45, 0.45, 0.45}), ShouldBeNil)
			/* The widget should be in sync with what we set at the shader */
			So(cw.Value().Equal(RtColor{0.45, 0.45, 0.45}), ShouldBeTrue)

			/* Here we over the shader again by including the emitColor inline; the constant.Handle()
			 * includes the handle created for the shader */
			ri.Bxdf("PxrConstant", constant.Handle(), RtToken("color emitColor"), RtColor{1, 0.25, 0.25})

			/* this is another way of writing our shader : using the emitColor that we actually set above */
			ri.Bxdf(constant.Name(), constant.Handle())

			ri.Sphere(1, -1, 1, 360)
			ri.WorldEnd()

			So(ri.End(), ShouldBeNil)

			/* gather and print the statistics and time the
			 * pipe took */
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

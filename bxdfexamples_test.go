package rigo

import (
	"testing"
	"os/user"
	"time"
	"fmt"
	
	
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/mae-global/rigo/ri"

	"github.com/mae-global/rigo/bxdf"
	"github.com/mae-global/rigo/bxdf/args"
)

func Test_BxdfExamples(t *testing.T) {

	Convey("Bxdf Constant Example",t,func() {

		cuser,err := user.Current()
		So(err,ShouldBeNil)

		pipe := DefaultFilePipe()
		So(pipe,ShouldNotBeNil)

		ctx := New(pipe,&Configuration{PrettyPrint:true})
		ctx.Begin("output/exampleBxdfConstant.rib")
		ctx.ArchiveRecord("structure","Scene Bxdf Constant")
		ctx.ArchiveRecord("structure","Creator %s",Author)
		ctx.ArchiveRecord("structure","CreationDate %s",time.Now())
		ctx.ArchiveRecord("structure","For %s",cuser.Username)
		ctx.ArchiveRecord("structure","Frame 1")

		ctx.Display("bxdf_sphere.tif","file","rgb")
		ctx.Format(320,240,1)
		ctx.Projection(PERSPECTIVE,RtString("fov"),RtFloat(30))
		ctx.Translate(0,0,6)
		ctx.WorldBegin()
		ctx.Color(RtColor{1,0,0})
		
		/* Create a bxdf interface from the PxrConstant.args file */
		constant,err := args.ParseFile("","PxrConstant")
		So(err,ShouldBeNil)
		So(constant,ShouldNotBeNil)

		/* grab a widget of the emit color attribute */
		w := constant.Widget("emitColor")
		So(w,ShouldNotBeNil)

		/* convert the interface to color widget */
		cw,ok := w.(*bxdf.RtColorWidget)
		So(ok,ShouldBeTrue)
		So(cw,ShouldNotBeNil)

		/* adjust the value via the Widget interface directly */
		So(w.SetValue(RtColor{0.3,0.2,0.1}),ShouldBeNil)
		
		/* adjust the emit color */
		So(cw.Set(RtColor{0.1,0.2,0.3}),ShouldBeNil)

		/* change the emit color directly */
		So(constant.SetValue(RtToken("emitColor"),RtColor{0.45,0.45,0.45}),ShouldBeNil)
		So(cw.Value().Equal(RtColor{0.45,0.45,0.45}),ShouldBeTrue)
		
		/* write the bxdf using the user function call */
		ctx.User(constant) /* constant will be tokenised into the RIB output */
		ctx.Sphere(1,-1,1,360)
		ctx.WorldEnd()

		So(ctx.End(),ShouldBeNil)

		p := pipe.GetByName(PipeToStats{}.Name())
		So(p,ShouldNotBeNil)
		s,ok := p.(*PipeToStats)
		So(s,ShouldNotBeNil)
		So(ok,ShouldBeTrue)	
	

		p = pipe.GetByName(PipeTimer{}.Name())
		So(p,ShouldNotBeNil)
		t,ok := p.(*PipeTimer)
		So(t,ShouldNotBeNil)
		So(ok,ShouldBeTrue)

		fmt.Printf("%s%s",s,t)
	})
}

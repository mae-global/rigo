/* rigo/fragment_test.go */
package rigo

import (
	"fmt"
	"os/user"
	"testing"
	"time"

	. "github.com/mae-global/rigo/ri"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ExampleOrangeBall(t *testing.T) {

	Convey("Example Orange Ball",t,func() {

		cuser,err := user.Current()
		So(err,ShouldBeNil)

		pipe := DefaultFilePipe()
		So(pipe,ShouldNotBeNil)

		ctx := New(pipe,&Configuration{PrettyPrint:true})
		ctx.Begin("output/exampleOrangeBall.rib")
		ctx.ArchiveRecord("structure","Scene Orange Ball")
		ctx.ArchiveRecord("structure","Creator %s-%s",Author,Version)
		ctx.ArchiveRecord("structure","CreationDate %s",time.Now())
		ctx.ArchiveRecord("structure","For %s",cuser.Username)
		ctx.ArchiveRecord("structure","Frames 5")
		
		frag := NewFragment("orangeball_fragment")
		So(frag,ShouldNotBeNil)
		/* grab the Renderman Interface from the fragment */
		fri := frag.Ri()
		So(fri,ShouldNotBeNil)

		fri.Format(640,480,-1)
		fri.ShadingRate(1)
		fri.Projection(Perspective,RtString("fov"),RtInt(30))
		fri.FrameAspectRatio(1.33)
		fri.Identity()
		fri.LightSource("distantlight",RtInt(1))
		fri.Translate(0,0,5)
		fri.WorldBegin()
		fri.Identity()
		fri.AttributeBegin()
		fri.Color(RtColor{1.0,0.6,0.0})
		fri.Surface("plastic",RtString("Ka"),RtFloat(1),RtString("Kd"),RtFloat(0.5),
												  RtString("Ks"),RtFloat(1),RtString("roughness"),RtFloat(0.1))
		fri.TransformBegin()
		fri.Rotate(90,1,0,0)
		fri.Sphere(1,-1,1,360)
		fri.TransformEnd()
		fri.AttributeEnd()
		fri.WorldEnd()

		So(frag.Statements(),ShouldEqual,17)

		for frame := 1; frame <= 5; frame++ {
			ctx.FrameBegin(RtInt(frame))
			ctx.Display(RtToken(fmt.Sprintf("orange_%03d.tif",frame)),"file","rgba")		

			frag.Replay(ctx)

			ctx.FrameEnd()
		}

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




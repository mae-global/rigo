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

		ri,pipe := DefaultPipeline(&Configuration{PrettyPrint:true})

		ri.Begin("output/exampleOrangeBall.rib")
		ri.ArchiveRecord("structure","Scene Orange Ball")
		ri.ArchiveRecord("structure","Creator %s",Author)
		ri.ArchiveRecord("structure","CreationDate %s",time.Now())
		ri.ArchiveRecord("structure","For %s",cuser.Username)
		ri.ArchiveRecord("structure","Frames 5")

		light,err := ri.LightHandle()
		So(err,ShouldBeNil)
		So(light.String(),ShouldEqual,"\"0\"")
		
		frag := DefaultFragment("orangeball_fragment")
		So(frag,ShouldNotBeNil)

		/* grab the Renderman Interface from the fragment */
		fri := RI(frag)
		So(fri,ShouldNotBeNil)

		fri.Format(640,480,-1)
		fri.ShadingRate(1)
		fri.Projection(PERSPECTIVE,RtString("fov"),RtInt(30))
		fri.FrameAspectRatio(1.33)
		fri.Identity()
		fri.LightSource("distantlight")
		fri.Illuminate(light,true)
		fri.Translate(0,0,5)
		fri.WorldBegin()
		fri.Identity()
		fri.AttributeBegin()
		fri.Color(RtColor{1.0,0.6,0.0})
		fri.Surface("plastic",RtToken("Ka"),RtFloat(1),RtToken("Kd"),RtFloat(0.5),
												  RtToken("Ks"),RtFloat(1),RtToken("roughness"),RtFloat(0.1))
		fri.TransformBegin()
		fri.Rotate(90,1,0,0)
		fri.Sphere(1,-1,1,360)
		fri.TransformEnd()
		fri.AttributeEnd()
		fri.WorldEnd()

		So(frag.Statements(),ShouldEqual,18)

		for frame := 1; frame <= 5; frame++ {
			ri.FrameBegin(RtInt(frame))
			ri.Display(RtToken(fmt.Sprintf("orange_%03d.tif",frame)),"file","rgba")		

			frag.Replay(ri)

			ri.FrameEnd()
		}

		So(ri.End(),ShouldBeNil)

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




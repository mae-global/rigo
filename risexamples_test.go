package rigo

import (
	"testing"
	"os/user"
	"time"
	"fmt"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/mae-global/rigo/ri"
  . "github.com/mae-global/rigo/ris"
)

func Test_RISExamples(t *testing.T) {

	Convey("RIS Examples",t,func() {
		Convey("PxrConstant -- bxdf shader",func() {

			cuser,err := user.Current()
			So(err,ShouldBeNil)

			pipe := DefaultFilePipe()
			So(pipe,ShouldNotBeNil)
			
			ctx := NewContext(pipe,nil,nil,nil,&Configuration{PrettyPrint:true})
			ri := RI(ctx)
			ri.Begin("output/risBxdfPxrConstant.rib")
			ri.ArchiveRecord("structure","Scene Bxdf Constant")
			ri.ArchiveRecord("structure","Creator %s",Author)
			ri.ArchiveRecord("structure","CreationDate %s",time.Now())
			ri.ArchiveRecord("structure","For %s",cuser.Username)
			ri.ArchiveRecord("structure","Frames 1")
			
			ri.Display("risbxdf_sphere.tif","file","rgb")
			ri.Format(320,240,1)	
			ri.PixelFilter(GaussianFilter,4,4)
			ri.Imager("background",RtToken("color color"),RtColor{.6,.6,.6},RtToken("float alpha"),RtFloat(1))

			ri.Projection(PERSPECTIVE,RtToken("fov"),RtFloat(30))
			ri.Translate(0,0,6)
			ri.WorldBegin()
			ri.Color(RtColor{1,0,0})

			ris := RIS(ctx)
			constant,err := ris.Bxdf("PxrConstant","-")
			So(err,ShouldBeNil)
			So(constant,ShouldNotBeNil)

			w := constant.Widget("emitColor")
			So(w,ShouldNotBeNil)

			cw,ok := w.(*RtColorWidget)
			So(ok,ShouldBeTrue)
			So(cw,ShouldNotBeNil)

			So(w.SetValue(RtColor{0.3,0.2,0.1}),ShouldBeNil)
			So(cw.Set(RtColor{0.1,0.2,0.3}),ShouldBeNil)
			So(constant.SetValue(RtToken("emitColor"),RtColor{0.45,0.45,0.45}),ShouldBeNil)
			So(cw.Value().Equal(RtColor{0.45,0.45,0.45}),ShouldBeTrue)

			ri.Bxdf("PxrConstant",constant.Handle(),RtToken("color emitColor"),RtColor{1,0.25,0.25})
			ri.Sphere(1,-1,1,360)
			ri.WorldEnd()

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
	})
}
			

package rigo

import (
	"testing"
	"fmt"

	. "github.com/mae-global/rigo/ri"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_HandleCallbacks(t *testing.T) {

	lighthandle := func(name RtName,light RtLightHandle) {

		fmt.Printf("LightHandle -- %s %s\n",name,light)
	}

	Convey("Handle callbacks pipelining",t,func() {

		pipe := NewPipe()
		callback := new(PipeHookHandleCallback)
		callback.LightHandler = lighthandle
	
		pipe.Append(callback)	

		mgr := NewHandleManager(nil,NewPrefixLightUniqueGenerator("light_"),nil)

		ri := RI(NewContext(pipe,mgr,nil))
		
		So(ri.Begin("output/filters-simple.rib"),ShouldBeNil)
		ri.Display("sphere.tif", "file", "rgb")
		ri.Format(320, 240, 1)
		ri.Projection(PERSPECTIVE, RtToken("fov"), RtFloat(30))
		ri.Translate(0, 0, 6)
		ri.WorldBegin()

			ambient,_ := ri.LightSource("ambientlight", RtToken("intensity"), RtFloat(0.5))
			So(ambient,ShouldNotBeNil)

			ri.LightSource("distantlight", RtToken("intensity"), RtFloat(1.2), RtToken("from"), RtIntArray{0, 0, -6}, RtToken("to"), RtIntArray{0, 0, 0})
			ri.Illuminate(ambient,false)	
			ri.Color(RtColor{1, 0, 0})
			ri.Sphere(1, -1, 1, 360)

		ri.WorldEnd()
		So(ri.End(),ShouldBeNil)
	
	})
}

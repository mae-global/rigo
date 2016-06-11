package issues

import (
	"testing"
	
	"github.com/mae-global/rigo"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Issue0001(t *testing.T) {

	Convey("Issue 0001 -- _version_ injection",t,func() {

		Convey("Problem",func() {

			buf := new(rigo.PipeToDebugBuffer)

			pipe := rigo.NewEmptyPipe()
			pipe.Append(buf)
	
			ctx := rigo.NewContext(pipe, nil, &rigo.Configuration{PrettyPrint: false})
			ri := rigo.RI(ctx)
			ri.Begin("issue0001.rib") 

			So(ri.ParseRIBString(RIBExample), ShouldBeNil)
			So(ri.End(), ShouldBeNil) 
			So(buf.String(),ShouldEqual,`Begin "issue0001.rib"  version 3.04  version 3.04  End  `)		
		})
	
		Convey("Fix",func() {

			buf := new(rigo.PipeToDebugBuffer)

			pipe := rigo.NewEmptyPipe()
			pipe.Append(new(rigo.PipeIssue0001Fix)).Append(buf)

			ctx := rigo.NewContext(pipe,nil,&rigo.Configuration{PrettyPrint: false})
			ri := rigo.RI(ctx)
			ri.Begin("issue0001.rib")

			So(ri.ParseRIBString(RIBExample), ShouldBeNil)
			So(ri.End(),ShouldBeNil)
			So(buf.String(),ShouldEqual,`Begin "issue0001.rib"  version 3.04  End  `)
		})
	})			
}

const RIBExample = `##RenderMan RIB-Structure 1.1
version 3.04
`


package args 

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"

	"fmt"
)

func Test_Parser(t *testing.T) {

	Convey("Parser",t,func() {

		args,err := Parse([]byte(constant_args))
		So(err,ShouldBeNil)
		So(args,ShouldNotBeNil)

		fmt.Printf("args is %v\n",args)

	})
}


const constant_args = `
<args format="1.0">
    <shaderType>
        <tag value="bxdf"/>
    </shaderType>
    <param label="Emit Color" name="emitColor" type="color" 
           default="1. 1. 1." widget="color">
        <tags>
            <tag value="color"/>
        </tags>
    </param>
    <param label="Presence" 
           name="presence" 
           type="float" 
           default="1" min="0" max="1"
           widget="default">
        <tags>
            <tag value="float"/>
        </tags>
        <help>
           help text was here
        </help>
    </param>
    <rfmdata nodeid="1053405" 
             classification="shader/surface:rendernode/RenderMan/bxdf:swatch/rmanSwatch"/>
</args>
`


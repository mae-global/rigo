package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_RIS(t *testing.T) {

	Convey("All RIS",t,func() {
		ctx := NewTest()
		So(ctx,ShouldNotBeNil)

		So(ctx.Integrator("PxrVCM","handle",RtToken("int maxPathLength"),RtInt(18),RtToken("int mergePaths"),RtInt(1)),ErrorShouldEqual,`Integrator "PxrVCM" "handle" "int maxPathLength" [18] "int mergePaths" [1]`)
		
		So(ctx.Bxdf("PxrDisney","dis1",RtToken("reference color baseColor"),RtString("mixer:result"),RtToken("reference float roughness"),RtString("tex1:result"),RtToken("float metallic"),RtFloat(1)),ErrorShouldEqual,`Bxdf "PxrDisney" "dis1" "reference color baseColor" ["mixer:result"] "reference float roughness" ["tex1:result"] "float metallic" [1]`)

		So(ctx.Pattern("PxrTexture","tex1",RtToken("string filename"),RtString("checker.tx")),ErrorShouldEqual,`Pattern "PxrTexture" "tex1" "string filename" ["checker.tx"]`)

	})
}

package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_External(t *testing.T) {

	Convey("All External", t, func() {

		ctx := NewTest()
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("external.rib"), ErrorShouldEqual, `Begin "external.rib"`)
		So(ctx.Comment("output from rigo, external_test.go"), ErrorShouldEqual, `# output from rigo, external_test.go`)

		So(ctx.MakeTexture("globe.pic", "globe.tx", "periodic", "clamp", GaussianFilter, 2.0, 2.0), ErrorShouldEqual, `MakeTexture "globe.pic" "globe.tx" "periodic" "clamp" "gaussian" 2 2`)
		So(ctx.MakeLatLongEnvironment("long.pic", "long.tx", CatmullRomFilter, 3, 3), ErrorShouldEqual, `MakeLatLongEnvironment "long.pic" "long.tx" "catmull-rom" 3 3`)
		So(ctx.MakeCubeFaceEnvironment("foo.x", "foo.nx", "foo.y", "foo.ny", "foo.z", "foo.nz", "foo.env", 95.0, TriangleFilter, 2.0, 2.0), ErrorShouldEqual, `MakeCubeFaceEnvironment "foo.x" "foo.nx" "foo.y" "foo.ny" "foo.z" "foo.nz" "foo.env" 95 "triangle" 2 2`)
		So(ctx.MakeShadow("shadow.pic", "shadow.tex"), ErrorShouldEqual, `MakeShadow "shadow.pic" "shadow.tex"`)
		So(ctx.ReadArchive("sodacan.rib", ReadArchiveCallback), ErrorShouldEqual, `ReadArchive "sodacan.rib"`)
		So(ctx.ArchiveRecord("comment", "this is just a test"), ErrorShouldEqual, `# this is just a test`)
		So(ctx.ArchiveRecord("structure", "hello there, %s!", "Alice"), ErrorShouldEqual, `## hello there, Alice!`)
		So(ctx.ArchiveRecord("verbatim", "This should not be currently implemented"), ErrorShouldEqual, `Verbatim This should not be currently implemented`)

		So(ctx.End(), ErrorShouldEqual, `End`)
	})
}

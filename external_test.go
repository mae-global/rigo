package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_External(t *testing.T) {

	Convey("All External", t, func() {

		ctx := New(nil)
		So(ctx, ShouldNotBeNil)

		So(ctx.Begin("output/external.rib"), ShouldBeNil)
		So(ctx.Comment("output from rigo, external_test.go"), ShouldBeNil)

		So(ctx.MakeTexture("globe.pic", "globe.tx", "periodic", "clamp", GaussianFilter, 2.0, 2.0), ShouldBeNil)
		So(ctx.MakeLatLongEnvironment("long.pic", "long.tx", CatmullRomFilter, 3, 3), ShouldBeNil)
		So(ctx.MakeCubeFaceEnvironment("foo.x", "foo.nx", "foo.y", "foo.ny", "foo.z", "foo.nz", "foo.env", 95.0, TriangleFilter, 2.0, 2.0), ShouldBeNil)
		So(ctx.MakeShadow("shadow.pic", "shadow.tex"), ShouldBeNil)
		So(ctx.ReadArchive("sodacan.rib", ReadArchiveCallback), ShouldBeNil)
		So(ctx.ArchiveRecord("comment", "this is just a test"), ShouldBeNil)
		So(ctx.ArchiveRecord("structure", "hello there, %s!", "Alice"), ShouldBeNil)
		So(ctx.ArchiveRecord("verbatim", "This should not be currently implemented"), ShouldBeNil)

		So(ctx.End(), ShouldBeNil)
	})
}

package rigo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Pipes(t *testing.T) {

	Convey("All Pipes", t, func() {

		pipe := DefaultFilePipe()
		So(pipe, ShouldNotBeNil)
		So(pipe.Run("foo", nil, nil, Info{}), ShouldEqual, ErrNoActiveContext)
	})
}

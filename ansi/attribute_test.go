package ansi_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/makyo/gotui/ansi"
)

func TestAttribute(t *testing.T) {
	Convey("When using an attribute", t, func() {

		Convey("It should have a start", func() {
			So(ansi.Bold.Start(), ShouldEqual, "\x1b[1m")
		})

		Convey("It should have an end", func() {
			So(ansi.Bold.End(), ShouldEqual, "\x1b[22m")
		})

		Convey("It should be able to apply to a string", func() {
			So(ansi.Bold.Apply("rose"), ShouldEqual, "\x1b[1mrose\x1b[22m")
		})
	})
}

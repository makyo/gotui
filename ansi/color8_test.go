package ansi_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/makyo/gotui/ansi"
)

func TestColor8(t *testing.T) {
	color, err := ansi.Colors8.Find("black")
	Convey("One should be able to find a color", t, func() {
		So(err, ShouldBeNil)
	})

	Convey("When using 3/4-bit color", t, func() {

		Convey("When using it for the foreground", func() {

			Convey("It should have a start", func() {
				So(color.FGStart(), ShouldEqual, "\x1b[30m")
			})

			Convey("It should be able to apply to a string", func() {
				So(color.FG("rose"), ShouldEqual, "\x1b[30mrose\x1b[39m")
			})
		})

		Convey("When using it for the background", func() {

			Convey("It should have a start", func() {
				So(color.BGStart(), ShouldEqual, "\x1b[40m")
			})

			Convey("It should be able to apply to a string", func() {
				So(color.BG("rose"), ShouldEqual, "\x1b[40mrose\x1b[49m")
			})
		})
	})
}

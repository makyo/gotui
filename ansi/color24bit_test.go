package ansi_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/makyo/gotui/ansi"
)

func TestColor24bit(t *testing.T) {
	Convey("One should be able to find a color", t, func() {

		Convey("By hex code", func() {
			_, err := ansi.Colors24bit.Find("#000000")
			So(err, ShouldBeNil)
		})

		Convey("By RGB code", func() {
			_, err := ansi.Colors24bit.Find("rgb(0, 0, 0)")
			So(err, ShouldBeNil)
		})

		Convey("By HSL code", func() {
			_, err := ansi.Colors24bit.Find("hsl(0, 0%, 0%)")
			So(err, ShouldBeNil)
		})
	})

	Convey("When using 3/4-bit color", t, func() {
		color, _ := ansi.Colors24bit.Find("#000000")

		Convey("When using it for the foreground", func() {

			Convey("It should have a start", func() {
				So(color.FGStart(), ShouldEqual, "\x1b[38;2;0;0;0m")
			})

			Convey("It should be able to apply to a string", func() {
				So(color.FG("rose"), ShouldEqual, "\x1b[38;2;0;0;0mrose\x1b[39m")
			})
		})

		Convey("When using it for the background", func() {

			Convey("It should have a start", func() {
				So(color.BGStart(), ShouldEqual, "\x1b[48;2;0;0;0m")
			})

			Convey("It should be able to apply to a string", func() {
				So(color.BG("rose"), ShouldEqual, "\x1b[48;2;0;0;0mrose\x1b[49m")
			})
		})
	})
}

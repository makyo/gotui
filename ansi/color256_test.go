package ansi_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/makyo/gotui/ansi"
)

func TestColor256(t *testing.T) {
	Convey("One should be able to find a color", t, func() {

		Convey("By ID", func() {
			_, err := ansi.Colors256.Find("0")
			So(err, ShouldBeNil)
		})

		Convey("By hex code", func() {
			_, err := ansi.Colors256.Find("#000000")
			So(err, ShouldBeNil)
		})

		Convey("By short RGB code", func() {
			_, err := ansi.Colors256.Find("rgb000")
			So(err, ShouldBeNil)
		})

		Convey("By RGB code", func() {
			_, err := ansi.Colors256.Find("rgb(0, 0, 0)")
			So(err, ShouldBeNil)
		})

		Convey("By HSL code", func() {
			_, err := ansi.Colors256.Find("hsl(0, 0%, 0%)")
			So(err, ShouldBeNil)
		})
	})

	Convey("When using 3/4-bit color", t, func() {
		color, _ := ansi.Colors256.Find("black")

		Convey("When using it for the foreground", func() {

			Convey("It should have a start", func() {
				So(color.FGStart(), ShouldEqual, "\x1b[38;5;0m")
			})

			Convey("It should be able to apply to a string", func() {
				So(color.FG("rose"), ShouldEqual, "\x1b[38;5;0mrose\x1b[39m")
			})
		})

		Convey("When using it for the background", func() {

			Convey("It should have a start", func() {
				So(color.BGStart(), ShouldEqual, "\x1b[48;5;0m")
			})

			Convey("It should be able to apply to a string", func() {
				So(color.BG("rose"), ShouldEqual, "\x1b[48;5;0mrose\x1b[49m")
			})
		})
	})
}

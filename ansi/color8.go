package ansi

import (
	"fmt"
)

type color8 uint8

func (c color8) FGStart() string {
	return fmt.Sprintf("\x1b[%dm", 30+c)
}

func (c color8) BGStart() string {
	return fmt.Sprintf("\x1b[%dm", 40+c)
}

func (c color8) FG(s string) string {
	return fmt.Sprintf("%s%s%s", c.FGStart(), s, FGEnd)
}

func (c color8) BG(s string) string {
	return fmt.Sprintf("%s%s%s", c.BGStart(), s, BGEnd)
}

type colors8 map[string]uint8

func (c colors8) Find(what string) (color8, error) {
	col, ok := c[what]
	if ok {
		return color8(col), nil
	}
	return 0, ColorNotFound
}

var Colors8 colors8 = map[string]uint8{
	"black":   0,
	"red":     1,
	"green":   2,
	"yellow":  3,
	"blue":    4,
	"magenta": 5,
	"cyan":    6,
	"white":   7,

	"brightblack":   60,
	"brightred":     61,
	"brightgreen":   62,
	"brightyellow":  63,
	"brightblue":    64,
	"brightmagenta": 65,
	"brightcyan":    66,
	"brightwhite":   67,
}

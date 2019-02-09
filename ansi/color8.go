package ansi

import (
	"fmt"
	"strings"
)

// color8 represents a 3/4-bit color
type color8 uint8

// FGStart returns the ANSI code to start writing runes in this color
// as the foreground.
func (c color8) FGStart() string {
	return fmt.Sprintf("\x1b[%dm", 30+c)
}

// BGStart returns the ANSI code to start writing runes in this color
// as the background.
func (c color8) BGStart() string {
	return fmt.Sprintf("\x1b[%dm", 40+c)
}

// FG turns the text for the provided string the specified color by surrounding
// it with the start and end codes.
func (c color8) FG(s string) string {
	return fmt.Sprintf("%s%s%s", c.FGStart(), s, FGEnd)
}

// BG turns the background for the provided string the specified color by
// surrounding it with the start and end codes.
func (c color8) BG(s string) string {
	return fmt.Sprintf("%s%s%s", c.BGStart(), s, BGEnd)
}

type colors8 map[string]uint8

// Find returns the color with the given name, or an error if one could not be
// found. It accepts the common name of the 3/4-bit color (case-insensitive).
// It will also find "bright" colors, which usually also work in most
// terminals, even if that just means defaulting to the non-bright verison in
// practice.
func (c colors8) Find(what string) (Color, error) {
	col, ok := c[strings.ToLower(what)]
	if ok {
		return color8(col), nil
	}
	return nil, ColorNotFound
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

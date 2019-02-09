package ansi

import (
	"fmt"
)

// Attribute describes an attribute that a string of runes might have.
// (e.g: bold, undelrine, etc)
type Attribute struct {
	// ANSI codes usually have one code to turn a functionality on, and
	// another to turn it off. These are labeled start and end here as they
	// often show up in pairs like this rather than being left on. The
	// universal end, of course, is 0.
	start, end uint8
}

// Reset is the ANSI code to reset all attributes and colors of a string.
const Reset string = "\x1b[0m"

// Start prints the start ANSI escape code for the attribute.
func (a Attribute) Start() string {
	return fmt.Sprintf("\x1b[%dm", a.start)
}

// End prints the ANSI escape code to turn off the attribute's functionality.
func (a Attribute) End() string {
	return fmt.Sprintf("\x1b[%dm", a.end)
}

// Apply turns on the attribute for the given string by surrounding it with
// the start and end codes.
func (a Attribute) Apply(s string) string {
	return fmt.Sprintf("%s%s%s", a.Start(), s, a.End())
}

var (
	Bold                    Attribute = Attribute{start: 1, end: 22}
	Faint                             = Attribute{start: 2, end: 22}
	Italic                            = Attribute{start: 3, end: 23}
	Underline                         = Attribute{start: 4, end: 24}
	Blink                             = Attribute{start: 5, end: 25}
	Flash                             = Attribute{start: 6, end: 25}
	Reverse                           = Attribute{start: 7, end: 27}
	Conceal                           = Attribute{start: 8, end: 28}
	CrossedOut                        = Attribute{start: 9, end: 29}
	AltFont1                          = Attribute{start: 11, end: 10}
	AltFont2                          = Attribute{start: 12, end: 10}
	AltFont3                          = Attribute{start: 13, end: 10}
	AltFont4                          = Attribute{start: 14, end: 10}
	AltFont5                          = Attribute{start: 15, end: 10}
	AltFont6                          = Attribute{start: 16, end: 10}
	AltFont7                          = Attribute{start: 17, end: 10}
	AltFont8                          = Attribute{start: 18, end: 10}
	AltFont9                          = Attribute{start: 19, end: 10}
	Fraktur                           = Attribute{start: 20, end: 23}
	DoubleUnderline                   = Attribute{start: 21, end: 24}
	Framed                            = Attribute{start: 51, end: 54}
	Encircled                         = Attribute{start: 52, end: 54}
	Overlined                         = Attribute{start: 53, end: 55}
	IdeogramUnderline                 = Attribute{start: 60, end: 65}
	IdeogramDoubleUnderline           = Attribute{start: 61, end: 65}
	IdeogramOverline                  = Attribute{start: 62, end: 65}
	IdeogramDoubleOverline            = Attribute{start: 63, end: 65}
	IdeogramStressMarking             = Attribute{start: 64, end: 65}
)

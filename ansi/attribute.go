package ansi

import (
	"fmt"
)

type Attribute struct {
	start, end uint8
}

func (a Attribute) Start() string {
	return fmt.Sprintf("\x1b[%dm", a.start)
}

func (a Attribute) End() string {
	return fmt.Sprintf("\x1b[%dm", a.end)
}

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

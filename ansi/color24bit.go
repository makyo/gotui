package ansi

import (
	"fmt"
	"strconv"
)

type color24bit rgb

// FGStart returns the ANSI code to start writing runes in this color
// as the foreground.
func (c color24bit) FGStart() string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", c.R, c.G, c.B)
}

// BGStart returns the ANSI code to start writing runes in this color
// as the background.
func (c color24bit) BGStart() string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", c.R, c.G, c.B)
}

// FG turns the text for the provided string the specified color by surrounding
// it with the start and end codes.
func (c color24bit) FG(s string) string {
	return fmt.Sprintf("%s%s%s", c.FGStart(), s, FGEnd)
}

// BG turns the background for the provided string the specified color by
// surrounding it with the start and end codes.
func (c color24bit) BG(s string) string {
	return fmt.Sprintf("%s%s%s", c.BGStart(), s, BGEnd)
}

type colors24bit struct{}

func (c colors24bit) Find(what string) (Color, error) {
	if matches := hexRegexp.FindAllStringSubmatch(what, -1); len(matches) == 1 && len(matches[0]) == 4 {
		r, _ := strconv.ParseUint(matches[0][1], 16, 8)
		g, _ := strconv.ParseUint(matches[0][2], 16, 8)
		b, _ := strconv.ParseUint(matches[0][3], 16, 8)
		return color24bit{R: uint8(r), G: uint8(g), B: uint8(b)}, nil
	} else if matches := rgbRegexp.FindAllStringSubmatch(what, -1); len(matches) == 1 && len(matches[0]) == 4 {
		var r, g, b int
		r, _ = strconv.Atoi(matches[0][1])
		g, _ = strconv.Atoi(matches[0][2])
		b, _ = strconv.Atoi(matches[0][3])
		if r > 255 || g > 255 || b > 255 {
			return nil, InvalidColorSpec
		}
		return color24bit{R: uint8(r), G: uint8(g), B: uint8(b)}, nil
	} else if matches := hslRegexp.FindAllStringSubmatch(what, -1); len(matches) == 1 && len(matches[0]) == 4 {
		var h, s, l int
		h, _ = strconv.Atoi(matches[0][1])
		s, _ = strconv.Atoi(matches[0][2])
		l, _ = strconv.Atoi(matches[0][3])
		if h > 360 || s > 100 || l > 100 {
			return nil, InvalidColorSpec
		}
		r, g, b := decodeHSL(h, s, l)
		return color24bit{R: r, G: g, B: b}, nil
	} else {
		return nil, ColorNotFound
	}
}

var Colors24bit colors24bit

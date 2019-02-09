package ansi

import (
	"fmt"
	"strconv"
)

type color24bit struct {
	R, G, B uint8
}

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

// decodeHSL returns RGB values on a scale from 0-255 given the hue, saturation,
// and lightness values. This conversion is not straight-forward, and the author
// doesn't totally understand it. Unashamed StackExchange-ing resulted in this.
func decodeHSL(_h, _s, _l int) (uint8, uint8, uint8) {
	h := float32(_h) / 360.0
	s := float32(_s) / 100.0
	l := float32(_l) / 100.0
	var r, g, b float32
	if s == 0 {
		r = l
		g = l
		b = l
	} else {
		var q float32
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q
		r = hue2rgb(p, q, h+1.0/3.0)
		g = hue2rgb(p, q, h)
		b = hue2rgb(p, q, h-1.0/3.0)
	}
	return uint8(r * 255), uint8(g * 255), uint8(b * 255)
}

// hue2rgb converts a hue value to an RGB value.
func hue2rgb(p, q, t float32) float32 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 0.5 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}

var Colors24bit colors24bit

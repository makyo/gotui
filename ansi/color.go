package ansi

import (
	"errors"
	"regexp"
)

var (
	// ColorNotFound is returned when calling Find does not come up with
	// anything.
	ColorNotFound error = errors.New("color not found")

	// InvalidColorSpec is returned when calling Find with an invalid string.
	InvalidColorSpec = errors.New("invalid color spec")

	hslRegexp *regexp.Regexp = regexp.MustCompile("^(?i)hsl\\((\\d+),\\s*(\\d+)%,\\s*(\\d+)%\\)$")
	rgbRegexp                = regexp.MustCompile("^(?i)rgb\\((\\d+),\\s*(\\d+),\\s*(\\d+)\\)$")
	hexRegexp                = regexp.MustCompile("^#([[:xdigit:]]{2})([[:xdigit:]]{2})([[:xdigit:]]{2})$")
)

// Color describes a color that a string of runes might have. This can be
// applied to both foreground and background.
type Color interface {
	FGStart() string
	BGStart() string
	FG(string) string
	BG(string) string
}

// Colors describes a colorspace that can be searched by a Find method.
type Colors interface {
	Find(string) (Color, error)
}

const (
	FGEnd string = "\x1b[39m"
	BGEnd        = "\x1b[49m"
)

// rgb represents a Red/Green/Blue specification for a color.
type rgb struct {
	R, G, B uint8
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

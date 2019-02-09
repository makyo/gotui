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

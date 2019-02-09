package ansi

import (
	"errors"
)

var ColorNotFound error = errors.New("color not found")

type Color interface {
	FGStart() string
	BGStart() string
	FG(string) string
	BG(string) string
	Reset() string
}

const (
	FGEnd string = "\x1b[39m"
	BGEnd        = "\x1b[49m"
)

package ansi

import (
	"fmt"
	"regexp"
	"strconv"
)

// color256 represents an 8-bit color and the ways in which it might be found.
type color256 struct {
	Id   uint8
	Name string
	RGB  rgb
}

// FGStart returns the ANSI code to start writing runes in this color
// as the foreground.
func (c color256) FGStart() string {
	return fmt.Sprintf("\x1b[38;5;%dm", c.Id)
}

// BGStart returns the ANSI code to start writing runes in this color
// as the background.
func (c color256) BGStart() string {
	return fmt.Sprintf("\x1b[48;5;%dm", c.Id)
}

// FG turns the text for the provided string the specified color by surrounding
// it with the start and end codes.
func (c color256) FG(s string) string {
	return fmt.Sprintf("%s%s%s", c.FGStart(), s, FGEnd)
}

// BG turns the background for the provided string the specified color by
// surrounding it with the start and end codes.
func (c color256) BG(s string) string {
	return fmt.Sprintf("%s%s%s", c.BGStart(), s, BGEnd)
}

type colors256 []color256

// Find returns the color with the given name, or an error if one could not be
// found. It accepts:
// * The color name
// * The color number as specified in the escape code
// * A CSS-style hex string (#rrggbb)
// * A CSS-style HSL specification (hsl(h, s%, l%))
// * A CSS style RGB specification (rgb(r, g, b))
// * A short-code rgb specification (rgb###)
// Note that, despite the wide array of possibilities provided by the CSS-style
// arguments, not every one exists in a 256-color set, so providing just any
// argument will not get you a color.
func (c colors256) Find(what string) (Color, error) {
	i, err := strconv.Atoi(what)
	if err == nil {
		return c.FindById(i)
	} else if len(what) == 6 && what[:3] == "rgb" {
		return c.FindByShortRGB(what)
	} else if matches := hexRegexp.FindAllStringSubmatch(what, -1); len(matches) == 1 && len(matches[0]) == 4 {
		r, _ := strconv.ParseUint(matches[0][1], 16, 8)
		g, _ := strconv.ParseUint(matches[0][2], 16, 8)
		b, _ := strconv.ParseUint(matches[0][3], 16, 8)
		return c.FindByRGB(int(r), int(g), int(b))
	} else if matches := rgbRegexp.FindAllStringSubmatch(what, -1); len(matches) == 1 && len(matches[0]) == 4 {
		var r, g, b int
		r, _ = strconv.Atoi(matches[0][1])
		g, _ = strconv.Atoi(matches[0][2])
		b, _ = strconv.Atoi(matches[0][3])
		return c.FindByRGB(r, g, b)
	} else if matches := hslRegexp.FindAllStringSubmatch(what, -1); len(matches) == 1 && len(matches[0]) == 4 {
		var h, s, l int
		h, _ = strconv.Atoi(matches[0][1])
		s, _ = strconv.Atoi(matches[0][2])
		l, _ = strconv.Atoi(matches[0][3])
		if h > 360 || s > 100 || l > 100 {
			return nil, InvalidColorSpec
		}
		r, g, b := decodeHSL(h, s, l)
		return c.FindByRGB(int(r), int(g), int(b))
	}
	return c.FindByName(what)
}

// findByShortRGB finds a color by RGB short code, which take the form of rgb###
// where each # is a digit from 0-5 inclusive. This works like CSS color codes,
// in that rgb000 is black, and rgb555 is white. This covers 216 of the 256
// colors in this space.
func (c colors256) FindByShortRGB(what string) (Color, error) {
	r, err := strconv.Atoi(string(what[3]))
	if err != nil {
		return nil, err
	}
	g, err := strconv.Atoi(string(what[4]))
	if err != nil {
		return nil, err
	}
	b, err := strconv.Atoi(string(what[5]))
	if err != nil {
		return nil, err
	}
	return c.FindById(16 + r*36 + 6*g + b)
}

// FindByRGB finds a color based on its CSS red/blue/green value.
func (c colors256) FindByRGB(r, g, b int) (Color, error) {
	for _, col := range c {
		if int(col.RGB.R) == r && int(col.RGB.G) == g && int(col.RGB.B) == b {
			return col, nil
		}
	}
	return nil, ColorNotFound
}

// FindById finds a color based on its ID, where ID is:
//     0-  7:  standard colors (as in ESC [ 30–37 m)
//     8- 15:  high intensity colors (as in ESC [ 90–97 m)
//    16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//   232-255:  grayscale from black to white in 24 steps
func (c colors256) FindById(what int) (Color, error) {
	if what >= 0 && what < len(c) {
		return c[what], nil
	}
	return nil, ColorNotFound
}

// FindByName finds a color based on its name.
func (c colors256) FindByName(what string) (Color, error) {
	for _, col := range c {
		if b, _ := regexp.MatchString(fmt.Sprintf("(?i)%s", col.Name), what); b {
			return col, nil
		}
	}
	return nil, ColorNotFound
}

var Colors256 colors256 = []color256{
	color256{
		Id:   0,
		Name: "Black",
		RGB: rgb{
			R: 0,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   1,
		Name: "Maroon",
		RGB: rgb{
			R: 128,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   2,
		Name: "Green",
		RGB: rgb{
			R: 0,
			G: 128,
			B: 0,
		},
	},
	color256{
		Id:   3,
		Name: "Olive",
		RGB: rgb{
			R: 128,
			G: 128,
			B: 0,
		},
	},
	color256{
		Id:   4,
		Name: "Navy",
		RGB: rgb{
			R: 0,
			G: 0,
			B: 128,
		},
	},
	color256{
		Id:   5,
		Name: "Purple",
		RGB: rgb{
			R: 128,
			G: 0,
			B: 128,
		},
	},
	color256{
		Id:   6,
		Name: "Teal",
		RGB: rgb{
			R: 0,
			G: 128,
			B: 128,
		},
	},
	color256{
		Id:   7,
		Name: "Silver",
		RGB: rgb{
			R: 192,
			G: 192,
			B: 192,
		},
	},
	color256{
		Id:   8,
		Name: "Grey",
		RGB: rgb{
			R: 128,
			G: 128,
			B: 128,
		},
	},
	color256{
		Id:   9,
		Name: "Red",
		RGB: rgb{
			R: 255,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   10,
		Name: "Lime",
		RGB: rgb{
			R: 0,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   11,
		Name: "Yellow",
		RGB: rgb{
			R: 255,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   12,
		Name: "Blue",
		RGB: rgb{
			R: 0,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   13,
		Name: "Fuchsia",
		RGB: rgb{
			R: 255,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   14,
		Name: "Aqua",
		RGB: rgb{
			R: 0,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   15,
		Name: "White",
		RGB: rgb{
			R: 255,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   16,
		Name: "Grey0",
		RGB: rgb{
			R: 0,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   17,
		Name: "NavyBlue",
		RGB: rgb{
			R: 0,
			G: 0,
			B: 95,
		},
	},
	color256{
		Id:   18,
		Name: "DarkBlue",
		RGB: rgb{
			R: 0,
			G: 0,
			B: 135,
		},
	},
	color256{
		Id:   19,
		Name: "Blue3",
		RGB: rgb{
			R: 0,
			G: 0,
			B: 175,
		},
	},
	color256{
		Id:   20,
		Name: "Blue3",
		RGB: rgb{
			R: 0,
			G: 0,
			B: 215,
		},
	},
	color256{
		Id:   21,
		Name: "Blue1",
		RGB: rgb{
			R: 0,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   22,
		Name: "DarkGreen",
		RGB: rgb{
			R: 0,
			G: 95,
			B: 0,
		},
	},
	color256{
		Id:   23,
		Name: "DeepSkyBlue4",
		RGB: rgb{
			R: 0,
			G: 95,
			B: 95,
		},
	},
	color256{
		Id:   24,
		Name: "DeepSkyBlue4",
		RGB: rgb{
			R: 0,
			G: 95,
			B: 135,
		},
	},
	color256{
		Id:   25,
		Name: "DeepSkyBlue4",
		RGB: rgb{
			R: 0,
			G: 95,
			B: 175,
		},
	},
	color256{
		Id:   26,
		Name: "DodgerBlue3",
		RGB: rgb{
			R: 0,
			G: 95,
			B: 215,
		},
	},
	color256{
		Id:   27,
		Name: "DodgerBlue2",
		RGB: rgb{
			R: 0,
			G: 95,
			B: 255,
		},
	},
	color256{
		Id:   28,
		Name: "Green4",
		RGB: rgb{
			R: 0,
			G: 135,
			B: 0,
		},
	},
	color256{
		Id:   29,
		Name: "SpringGreen4",
		RGB: rgb{
			R: 0,
			G: 135,
			B: 95,
		},
	},
	color256{
		Id:   30,
		Name: "Turquoise4",
		RGB: rgb{
			R: 0,
			G: 135,
			B: 135,
		},
	},
	color256{
		Id:   31,
		Name: "DeepSkyBlue3",
		RGB: rgb{
			R: 0,
			G: 135,
			B: 175,
		},
	},
	color256{
		Id:   32,
		Name: "DeepSkyBlue3",
		RGB: rgb{
			R: 0,
			G: 135,
			B: 215,
		},
	},
	color256{
		Id:   33,
		Name: "DodgerBlue1",
		RGB: rgb{
			R: 0,
			G: 135,
			B: 255,
		},
	},
	color256{
		Id:   34,
		Name: "Green3",
		RGB: rgb{
			R: 0,
			G: 175,
			B: 0,
		},
	},
	color256{
		Id:   35,
		Name: "SpringGreen3",
		RGB: rgb{
			R: 0,
			G: 175,
			B: 95,
		},
	},
	color256{
		Id:   36,
		Name: "DarkCyan",
		RGB: rgb{
			R: 0,
			G: 175,
			B: 135,
		},
	},
	color256{
		Id:   37,
		Name: "LightSeaGreen",
		RGB: rgb{
			R: 0,
			G: 175,
			B: 175,
		},
	},
	color256{
		Id:   38,
		Name: "DeepSkyBlue2",
		RGB: rgb{
			R: 0,
			G: 175,
			B: 215,
		},
	},
	color256{
		Id:   39,
		Name: "DeepSkyBlue1",
		RGB: rgb{
			R: 0,
			G: 175,
			B: 255,
		},
	},
	color256{
		Id:   40,
		Name: "Green3",
		RGB: rgb{
			R: 0,
			G: 215,
			B: 0,
		},
	},
	color256{
		Id:   41,
		Name: "SpringGreen3",
		RGB: rgb{
			R: 0,
			G: 215,
			B: 95,
		},
	},
	color256{
		Id:   42,
		Name: "SpringGreen2",
		RGB: rgb{
			R: 0,
			G: 215,
			B: 135,
		},
	},
	color256{
		Id:   43,
		Name: "Cyan3",
		RGB: rgb{
			R: 0,
			G: 215,
			B: 175,
		},
	},
	color256{
		Id:   44,
		Name: "DarkTurquoise",
		RGB: rgb{
			R: 0,
			G: 215,
			B: 215,
		},
	},
	color256{
		Id:   45,
		Name: "Turquoise2",
		RGB: rgb{
			R: 0,
			G: 215,
			B: 255,
		},
	},
	color256{
		Id:   46,
		Name: "Green1",
		RGB: rgb{
			R: 0,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   47,
		Name: "SpringGreen2",
		RGB: rgb{
			R: 0,
			G: 255,
			B: 95,
		},
	},
	color256{
		Id:   48,
		Name: "SpringGreen1",
		RGB: rgb{
			R: 0,
			G: 255,
			B: 135,
		},
	},
	color256{
		Id:   49,
		Name: "MediumSpringGreen",
		RGB: rgb{
			R: 0,
			G: 255,
			B: 175,
		},
	},
	color256{
		Id:   50,
		Name: "Cyan2",
		RGB: rgb{
			R: 0,
			G: 255,
			B: 215,
		},
	},
	color256{
		Id:   51,
		Name: "Cyan1",
		RGB: rgb{
			R: 0,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   52,
		Name: "DarkRed",
		RGB: rgb{
			R: 95,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   53,
		Name: "DeepPink4",
		RGB: rgb{
			R: 95,
			G: 0,
			B: 95,
		},
	},
	color256{
		Id:   54,
		Name: "Purple4",
		RGB: rgb{
			R: 95,
			G: 0,
			B: 135,
		},
	},
	color256{
		Id:   55,
		Name: "Purple4",
		RGB: rgb{
			R: 95,
			G: 0,
			B: 175,
		},
	},
	color256{
		Id:   56,
		Name: "Purple3",
		RGB: rgb{
			R: 95,
			G: 0,
			B: 215,
		},
	},
	color256{
		Id:   57,
		Name: "BlueViolet",
		RGB: rgb{
			R: 95,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   58,
		Name: "Orange4",
		RGB: rgb{
			R: 95,
			G: 95,
			B: 0,
		},
	},
	color256{
		Id:   59,
		Name: "Grey37",
		RGB: rgb{
			R: 95,
			G: 95,
			B: 95,
		},
	},
	color256{
		Id:   60,
		Name: "MediumPurple4",
		RGB: rgb{
			R: 95,
			G: 95,
			B: 135,
		},
	},
	color256{
		Id:   61,
		Name: "SlateBlue3",
		RGB: rgb{
			R: 95,
			G: 95,
			B: 175,
		},
	},
	color256{
		Id:   62,
		Name: "SlateBlue3",
		RGB: rgb{
			R: 95,
			G: 95,
			B: 215,
		},
	},
	color256{
		Id:   63,
		Name: "RoyalBlue1",
		RGB: rgb{
			R: 95,
			G: 95,
			B: 255,
		},
	},
	color256{
		Id:   64,
		Name: "Chartreuse4",
		RGB: rgb{
			R: 95,
			G: 135,
			B: 0,
		},
	},
	color256{
		Id:   65,
		Name: "DarkSeaGreen4",
		RGB: rgb{
			R: 95,
			G: 135,
			B: 95,
		},
	},
	color256{
		Id:   66,
		Name: "PaleTurquoise4",
		RGB: rgb{
			R: 95,
			G: 135,
			B: 135,
		},
	},
	color256{
		Id:   67,
		Name: "SteelBlue",
		RGB: rgb{
			R: 95,
			G: 135,
			B: 175,
		},
	},
	color256{
		Id:   68,
		Name: "SteelBlue3",
		RGB: rgb{
			R: 95,
			G: 135,
			B: 215,
		},
	},
	color256{
		Id:   69,
		Name: "CornflowerBlue",
		RGB: rgb{
			R: 95,
			G: 135,
			B: 255,
		},
	},
	color256{
		Id:   70,
		Name: "Chartreuse3",
		RGB: rgb{
			R: 95,
			G: 175,
			B: 0,
		},
	},
	color256{
		Id:   71,
		Name: "DarkSeaGreen4",
		RGB: rgb{
			R: 95,
			G: 175,
			B: 95,
		},
	},
	color256{
		Id:   72,
		Name: "CadetBlue",
		RGB: rgb{
			R: 95,
			G: 175,
			B: 135,
		},
	},
	color256{
		Id:   73,
		Name: "CadetBlue",
		RGB: rgb{
			R: 95,
			G: 175,
			B: 175,
		},
	},
	color256{
		Id:   74,
		Name: "SkyBlue3",
		RGB: rgb{
			R: 95,
			G: 175,
			B: 215,
		},
	},
	color256{
		Id:   75,
		Name: "SteelBlue1",
		RGB: rgb{
			R: 95,
			G: 175,
			B: 255,
		},
	},
	color256{
		Id:   76,
		Name: "Chartreuse3",
		RGB: rgb{
			R: 95,
			G: 215,
			B: 0,
		},
	},
	color256{
		Id:   77,
		Name: "PaleGreen3",
		RGB: rgb{
			R: 95,
			G: 215,
			B: 95,
		},
	},
	color256{
		Id:   78,
		Name: "SeaGreen3",
		RGB: rgb{
			R: 95,
			G: 215,
			B: 135,
		},
	},
	color256{
		Id:   79,
		Name: "Aquamarine3",
		RGB: rgb{
			R: 95,
			G: 215,
			B: 175,
		},
	},
	color256{
		Id:   80,
		Name: "MediumTurquoise",
		RGB: rgb{
			R: 95,
			G: 215,
			B: 215,
		},
	},
	color256{
		Id:   81,
		Name: "SteelBlue1",
		RGB: rgb{
			R: 95,
			G: 215,
			B: 255,
		},
	},
	color256{
		Id:   82,
		Name: "Chartreuse2",
		RGB: rgb{
			R: 95,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   83,
		Name: "SeaGreen2",
		RGB: rgb{
			R: 95,
			G: 255,
			B: 95,
		},
	},
	color256{
		Id:   84,
		Name: "SeaGreen1",
		RGB: rgb{
			R: 95,
			G: 255,
			B: 135,
		},
	},
	color256{
		Id:   85,
		Name: "SeaGreen1",
		RGB: rgb{
			R: 95,
			G: 255,
			B: 175,
		},
	},
	color256{
		Id:   86,
		Name: "Aquamarine1",
		RGB: rgb{
			R: 95,
			G: 255,
			B: 215,
		},
	},
	color256{
		Id:   87,
		Name: "DarkSlateGray2",
		RGB: rgb{
			R: 95,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   88,
		Name: "DarkRed",
		RGB: rgb{
			R: 135,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   89,
		Name: "DeepPink4",
		RGB: rgb{
			R: 135,
			G: 0,
			B: 95,
		},
	},
	color256{
		Id:   90,
		Name: "DarkMagenta",
		RGB: rgb{
			R: 135,
			G: 0,
			B: 135,
		},
	},
	color256{
		Id:   91,
		Name: "DarkMagenta",
		RGB: rgb{
			R: 135,
			G: 0,
			B: 175,
		},
	},
	color256{
		Id:   92,
		Name: "DarkViolet",
		RGB: rgb{
			R: 135,
			G: 0,
			B: 215,
		},
	},
	color256{
		Id:   93,
		Name: "Purple",
		RGB: rgb{
			R: 135,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   94,
		Name: "Orange4",
		RGB: rgb{
			R: 135,
			G: 95,
			B: 0,
		},
	},
	color256{
		Id:   95,
		Name: "LightPink4",
		RGB: rgb{
			R: 135,
			G: 95,
			B: 95,
		},
	},
	color256{
		Id:   96,
		Name: "Plum4",
		RGB: rgb{
			R: 135,
			G: 95,
			B: 135,
		},
	},
	color256{
		Id:   97,
		Name: "MediumPurple3",
		RGB: rgb{
			R: 135,
			G: 95,
			B: 175,
		},
	},
	color256{
		Id:   98,
		Name: "MediumPurple3",
		RGB: rgb{
			R: 135,
			G: 95,
			B: 215,
		},
	},
	color256{
		Id:   99,
		Name: "SlateBlue1",
		RGB: rgb{
			R: 135,
			G: 95,
			B: 255,
		},
	},
	color256{
		Id:   100,
		Name: "Yellow4",
		RGB: rgb{
			R: 135,
			G: 135,
			B: 0,
		},
	},
	color256{
		Id:   101,
		Name: "Wheat4",
		RGB: rgb{
			R: 135,
			G: 135,
			B: 95,
		},
	},
	color256{
		Id:   102,
		Name: "Grey53",
		RGB: rgb{
			R: 135,
			G: 135,
			B: 135,
		},
	},
	color256{
		Id:   103,
		Name: "LightSlateGrey",
		RGB: rgb{
			R: 135,
			G: 135,
			B: 175,
		},
	},
	color256{
		Id:   104,
		Name: "MediumPurple",
		RGB: rgb{
			R: 135,
			G: 135,
			B: 215,
		},
	},
	color256{
		Id:   105,
		Name: "LightSlateBlue",
		RGB: rgb{
			R: 135,
			G: 135,
			B: 255,
		},
	},
	color256{
		Id:   106,
		Name: "Yellow4",
		RGB: rgb{
			R: 135,
			G: 175,
			B: 0,
		},
	},
	color256{
		Id:   107,
		Name: "DarkOliveGreen3",
		RGB: rgb{
			R: 135,
			G: 175,
			B: 95,
		},
	},
	color256{
		Id:   108,
		Name: "DarkSeaGreen",
		RGB: rgb{
			R: 135,
			G: 175,
			B: 135,
		},
	},
	color256{
		Id:   109,
		Name: "LightSkyBlue3",
		RGB: rgb{
			R: 135,
			G: 175,
			B: 175,
		},
	},
	color256{
		Id:   110,
		Name: "LightSkyBlue3",
		RGB: rgb{
			R: 135,
			G: 175,
			B: 215,
		},
	},
	color256{
		Id:   111,
		Name: "SkyBlue2",
		RGB: rgb{
			R: 135,
			G: 175,
			B: 255,
		},
	},
	color256{
		Id:   112,
		Name: "Chartreuse2",
		RGB: rgb{
			R: 135,
			G: 215,
			B: 0,
		},
	},
	color256{
		Id:   113,
		Name: "DarkOliveGreen3",
		RGB: rgb{
			R: 135,
			G: 215,
			B: 95,
		},
	},
	color256{
		Id:   114,
		Name: "PaleGreen3",
		RGB: rgb{
			R: 135,
			G: 215,
			B: 135,
		},
	},
	color256{
		Id:   115,
		Name: "DarkSeaGreen3",
		RGB: rgb{
			R: 135,
			G: 215,
			B: 175,
		},
	},
	color256{
		Id:   116,
		Name: "DarkSlateGray3",
		RGB: rgb{
			R: 135,
			G: 215,
			B: 215,
		},
	},
	color256{
		Id:   117,
		Name: "SkyBlue1",
		RGB: rgb{
			R: 135,
			G: 215,
			B: 255,
		},
	},
	color256{
		Id:   118,
		Name: "Chartreuse1",
		RGB: rgb{
			R: 135,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   119,
		Name: "LightGreen",
		RGB: rgb{
			R: 135,
			G: 255,
			B: 95,
		},
	},
	color256{
		Id:   120,
		Name: "LightGreen",
		RGB: rgb{
			R: 135,
			G: 255,
			B: 135,
		},
	},
	color256{
		Id:   121,
		Name: "PaleGreen1",
		RGB: rgb{
			R: 135,
			G: 255,
			B: 175,
		},
	},
	color256{
		Id:   122,
		Name: "Aquamarine1",
		RGB: rgb{
			R: 135,
			G: 255,
			B: 215,
		},
	},
	color256{
		Id:   123,
		Name: "DarkSlateGray1",
		RGB: rgb{
			R: 135,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   124,
		Name: "Red3",
		RGB: rgb{
			R: 175,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   125,
		Name: "DeepPink4",
		RGB: rgb{
			R: 175,
			G: 0,
			B: 95,
		},
	},
	color256{
		Id:   126,
		Name: "MediumVioletRed",
		RGB: rgb{
			R: 175,
			G: 0,
			B: 135,
		},
	},
	color256{
		Id:   127,
		Name: "Magenta3",
		RGB: rgb{
			R: 175,
			G: 0,
			B: 175,
		},
	},
	color256{
		Id:   128,
		Name: "DarkViolet",
		RGB: rgb{
			R: 175,
			G: 0,
			B: 215,
		},
	},
	color256{
		Id:   129,
		Name: "Purple",
		RGB: rgb{
			R: 175,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   130,
		Name: "DarkOrange3",
		RGB: rgb{
			R: 175,
			G: 95,
			B: 0,
		},
	},
	color256{
		Id:   131,
		Name: "IndianRed",
		RGB: rgb{
			R: 175,
			G: 95,
			B: 95,
		},
	},
	color256{
		Id:   132,
		Name: "HotPink3",
		RGB: rgb{
			R: 175,
			G: 95,
			B: 135,
		},
	},
	color256{
		Id:   133,
		Name: "MediumOrchId3",
		RGB: rgb{
			R: 175,
			G: 95,
			B: 175,
		},
	},
	color256{
		Id:   134,
		Name: "MediumOrchId",
		RGB: rgb{
			R: 175,
			G: 95,
			B: 215,
		},
	},
	color256{
		Id:   135,
		Name: "MediumPurple2",
		RGB: rgb{
			R: 175,
			G: 95,
			B: 255,
		},
	},
	color256{
		Id:   136,
		Name: "DarkGoldenrod",
		RGB: rgb{
			R: 175,
			G: 135,
			B: 0,
		},
	},
	color256{
		Id:   137,
		Name: "LightSalmon3",
		RGB: rgb{
			R: 175,
			G: 135,
			B: 95,
		},
	},
	color256{
		Id:   138,
		Name: "RosyBrown",
		RGB: rgb{
			R: 175,
			G: 135,
			B: 135,
		},
	},
	color256{
		Id:   139,
		Name: "Grey63",
		RGB: rgb{
			R: 175,
			G: 135,
			B: 175,
		},
	},
	color256{
		Id:   140,
		Name: "MediumPurple2",
		RGB: rgb{
			R: 175,
			G: 135,
			B: 215,
		},
	},
	color256{
		Id:   141,
		Name: "MediumPurple1",
		RGB: rgb{
			R: 175,
			G: 135,
			B: 255,
		},
	},
	color256{
		Id:   142,
		Name: "Gold3",
		RGB: rgb{
			R: 175,
			G: 175,
			B: 0,
		},
	},
	color256{
		Id:   143,
		Name: "DarkKhaki",
		RGB: rgb{
			R: 175,
			G: 175,
			B: 95,
		},
	},
	color256{
		Id:   144,
		Name: "NavajoWhite3",
		RGB: rgb{
			R: 175,
			G: 175,
			B: 135,
		},
	},
	color256{
		Id:   145,
		Name: "Grey69",
		RGB: rgb{
			R: 175,
			G: 175,
			B: 175,
		},
	},
	color256{
		Id:   146,
		Name: "LightSteelBlue3",
		RGB: rgb{
			R: 175,
			G: 175,
			B: 215,
		},
	},
	color256{
		Id:   147,
		Name: "LightSteelBlue",
		RGB: rgb{
			R: 175,
			G: 175,
			B: 255,
		},
	},
	color256{
		Id:   148,
		Name: "Yellow3",
		RGB: rgb{
			R: 175,
			G: 215,
			B: 0,
		},
	},
	color256{
		Id:   149,
		Name: "DarkOliveGreen3",
		RGB: rgb{
			R: 175,
			G: 215,
			B: 95,
		},
	},
	color256{
		Id:   150,
		Name: "DarkSeaGreen3",
		RGB: rgb{
			R: 175,
			G: 215,
			B: 135,
		},
	},
	color256{
		Id:   151,
		Name: "DarkSeaGreen2",
		RGB: rgb{
			R: 175,
			G: 215,
			B: 175,
		},
	},
	color256{
		Id:   152,
		Name: "LightCyan3",
		RGB: rgb{
			R: 175,
			G: 215,
			B: 215,
		},
	},
	color256{
		Id:   153,
		Name: "LightSkyBlue1",
		RGB: rgb{
			R: 175,
			G: 215,
			B: 255,
		},
	},
	color256{
		Id:   154,
		Name: "GreenYellow",
		RGB: rgb{
			R: 175,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   155,
		Name: "DarkOliveGreen2",
		RGB: rgb{
			R: 175,
			G: 255,
			B: 95,
		},
	},
	color256{
		Id:   156,
		Name: "PaleGreen1",
		RGB: rgb{
			R: 175,
			G: 255,
			B: 135,
		},
	},
	color256{
		Id:   157,
		Name: "DarkSeaGreen2",
		RGB: rgb{
			R: 175,
			G: 255,
			B: 175,
		},
	},
	color256{
		Id:   158,
		Name: "DarkSeaGreen1",
		RGB: rgb{
			R: 175,
			G: 255,
			B: 215,
		},
	},
	color256{
		Id:   159,
		Name: "PaleTurquoise1",
		RGB: rgb{
			R: 175,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   160,
		Name: "Red3",
		RGB: rgb{
			R: 215,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   161,
		Name: "DeepPink3",
		RGB: rgb{
			R: 215,
			G: 0,
			B: 95,
		},
	},
	color256{
		Id:   162,
		Name: "DeepPink3",
		RGB: rgb{
			R: 215,
			G: 0,
			B: 135,
		},
	},
	color256{
		Id:   163,
		Name: "Magenta3",
		RGB: rgb{
			R: 215,
			G: 0,
			B: 175,
		},
	},
	color256{
		Id:   164,
		Name: "Magenta3",
		RGB: rgb{
			R: 215,
			G: 0,
			B: 215,
		},
	},
	color256{
		Id:   165,
		Name: "Magenta2",
		RGB: rgb{
			R: 215,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   166,
		Name: "DarkOrange3",
		RGB: rgb{
			R: 215,
			G: 95,
			B: 0,
		},
	},
	color256{
		Id:   167,
		Name: "IndianRed",
		RGB: rgb{
			R: 215,
			G: 95,
			B: 95,
		},
	},
	color256{
		Id:   168,
		Name: "HotPink3",
		RGB: rgb{
			R: 215,
			G: 95,
			B: 135,
		},
	},
	color256{
		Id:   169,
		Name: "HotPink2",
		RGB: rgb{
			R: 215,
			G: 95,
			B: 175,
		},
	},
	color256{
		Id:   170,
		Name: "OrchId",
		RGB: rgb{
			R: 215,
			G: 95,
			B: 215,
		},
	},
	color256{
		Id:   171,
		Name: "MediumOrchId1",
		RGB: rgb{
			R: 215,
			G: 95,
			B: 255,
		},
	},
	color256{
		Id:   172,
		Name: "Orange3",
		RGB: rgb{
			R: 215,
			G: 135,
			B: 0,
		},
	},
	color256{
		Id:   173,
		Name: "LightSalmon3",
		RGB: rgb{
			R: 215,
			G: 135,
			B: 95,
		},
	},
	color256{
		Id:   174,
		Name: "LightPink3",
		RGB: rgb{
			R: 215,
			G: 135,
			B: 135,
		},
	},
	color256{
		Id:   175,
		Name: "Pink3",
		RGB: rgb{
			R: 215,
			G: 135,
			B: 175,
		},
	},
	color256{
		Id:   176,
		Name: "Plum3",
		RGB: rgb{
			R: 215,
			G: 135,
			B: 215,
		},
	},
	color256{
		Id:   177,
		Name: "Violet",
		RGB: rgb{
			R: 215,
			G: 135,
			B: 255,
		},
	},
	color256{
		Id:   178,
		Name: "Gold3",
		RGB: rgb{
			R: 215,
			G: 175,
			B: 0,
		},
	},
	color256{
		Id:   179,
		Name: "LightGoldenrod3",
		RGB: rgb{
			R: 215,
			G: 175,
			B: 95,
		},
	},
	color256{
		Id:   180,
		Name: "Tan",
		RGB: rgb{
			R: 215,
			G: 175,
			B: 135,
		},
	},
	color256{
		Id:   181,
		Name: "MistyRose3",
		RGB: rgb{
			R: 215,
			G: 175,
			B: 175,
		},
	},
	color256{
		Id:   182,
		Name: "Thistle3",
		RGB: rgb{
			R: 215,
			G: 175,
			B: 215,
		},
	},
	color256{
		Id:   183,
		Name: "Plum2",
		RGB: rgb{
			R: 215,
			G: 175,
			B: 255,
		},
	},
	color256{
		Id:   184,
		Name: "Yellow3",
		RGB: rgb{
			R: 215,
			G: 215,
			B: 0,
		},
	},
	color256{
		Id:   185,
		Name: "Khaki3",
		RGB: rgb{
			R: 215,
			G: 215,
			B: 95,
		},
	},
	color256{
		Id:   186,
		Name: "LightGoldenrod2",
		RGB: rgb{
			R: 215,
			G: 215,
			B: 135,
		},
	},
	color256{
		Id:   187,
		Name: "LightYellow3",
		RGB: rgb{
			R: 215,
			G: 215,
			B: 175,
		},
	},
	color256{
		Id:   188,
		Name: "Grey84",
		RGB: rgb{
			R: 215,
			G: 215,
			B: 215,
		},
	},
	color256{
		Id:   189,
		Name: "LightSteelBlue1",
		RGB: rgb{
			R: 215,
			G: 215,
			B: 255,
		},
	},
	color256{
		Id:   190,
		Name: "Yellow2",
		RGB: rgb{
			R: 215,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   191,
		Name: "DarkOliveGreen1",
		RGB: rgb{
			R: 215,
			G: 255,
			B: 95,
		},
	},
	color256{
		Id:   192,
		Name: "DarkOliveGreen1",
		RGB: rgb{
			R: 215,
			G: 255,
			B: 135,
		},
	},
	color256{
		Id:   193,
		Name: "DarkSeaGreen1",
		RGB: rgb{
			R: 215,
			G: 255,
			B: 175,
		},
	},
	color256{
		Id:   194,
		Name: "Honeydew2",
		RGB: rgb{
			R: 215,
			G: 255,
			B: 215,
		},
	},
	color256{
		Id:   195,
		Name: "LightCyan1",
		RGB: rgb{
			R: 215,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   196,
		Name: "Red1",
		RGB: rgb{
			R: 255,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   197,
		Name: "DeepPink2",
		RGB: rgb{
			R: 255,
			G: 0,
			B: 95,
		},
	},
	color256{
		Id:   198,
		Name: "DeepPink1",
		RGB: rgb{
			R: 255,
			G: 0,
			B: 135,
		},
	},
	color256{
		Id:   199,
		Name: "DeepPink1",
		RGB: rgb{
			R: 255,
			G: 0,
			B: 175,
		},
	},
	color256{
		Id:   200,
		Name: "Magenta2",
		RGB: rgb{
			R: 255,
			G: 0,
			B: 215,
		},
	},
	color256{
		Id:   201,
		Name: "Magenta1",
		RGB: rgb{
			R: 255,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   202,
		Name: "OrangeRed1",
		RGB: rgb{
			R: 255,
			G: 95,
			B: 0,
		},
	},
	color256{
		Id:   203,
		Name: "IndianRed1",
		RGB: rgb{
			R: 255,
			G: 95,
			B: 95,
		},
	},
	color256{
		Id:   204,
		Name: "IndianRed1",
		RGB: rgb{
			R: 255,
			G: 95,
			B: 135,
		},
	},
	color256{
		Id:   205,
		Name: "HotPink",
		RGB: rgb{
			R: 255,
			G: 95,
			B: 175,
		},
	},
	color256{
		Id:   206,
		Name: "HotPink",
		RGB: rgb{
			R: 255,
			G: 95,
			B: 215,
		},
	},
	color256{
		Id:   207,
		Name: "MediumOrchId1",
		RGB: rgb{
			R: 255,
			G: 95,
			B: 255,
		},
	},
	color256{
		Id:   208,
		Name: "DarkOrange",
		RGB: rgb{
			R: 255,
			G: 135,
			B: 0,
		},
	},
	color256{
		Id:   209,
		Name: "Salmon1",
		RGB: rgb{
			R: 255,
			G: 135,
			B: 95,
		},
	},
	color256{
		Id:   210,
		Name: "LightCoral",
		RGB: rgb{
			R: 255,
			G: 135,
			B: 135,
		},
	},
	color256{
		Id:   211,
		Name: "PaleVioletRed1",
		RGB: rgb{
			R: 255,
			G: 135,
			B: 175,
		},
	},
	color256{
		Id:   212,
		Name: "OrchId2",
		RGB: rgb{
			R: 255,
			G: 135,
			B: 215,
		},
	},
	color256{
		Id:   213,
		Name: "OrchId1",
		RGB: rgb{
			R: 255,
			G: 135,
			B: 255,
		},
	},
	color256{
		Id:   214,
		Name: "Orange1",
		RGB: rgb{
			R: 255,
			G: 175,
			B: 0,
		},
	},
	color256{
		Id:   215,
		Name: "SandyBrown",
		RGB: rgb{
			R: 255,
			G: 175,
			B: 95,
		},
	},
	color256{
		Id:   216,
		Name: "LightSalmon1",
		RGB: rgb{
			R: 255,
			G: 175,
			B: 135,
		},
	},
	color256{
		Id:   217,
		Name: "LightPink1",
		RGB: rgb{
			R: 255,
			G: 175,
			B: 175,
		},
	},
	color256{
		Id:   218,
		Name: "Pink1",
		RGB: rgb{
			R: 255,
			G: 175,
			B: 215,
		},
	},
	color256{
		Id:   219,
		Name: "Plum1",
		RGB: rgb{
			R: 255,
			G: 175,
			B: 255,
		},
	},
	color256{
		Id:   220,
		Name: "Gold1",
		RGB: rgb{
			R: 255,
			G: 215,
			B: 0,
		},
	},
	color256{
		Id:   221,
		Name: "LightGoldenrod2",
		RGB: rgb{
			R: 255,
			G: 215,
			B: 95,
		},
	},
	color256{
		Id:   222,
		Name: "LightGoldenrod2",
		RGB: rgb{
			R: 255,
			G: 215,
			B: 135,
		},
	},
	color256{
		Id:   223,
		Name: "NavajoWhite1",
		RGB: rgb{
			R: 255,
			G: 215,
			B: 175,
		},
	},
	color256{
		Id:   224,
		Name: "MistyRose1",
		RGB: rgb{
			R: 255,
			G: 215,
			B: 215,
		},
	},
	color256{
		Id:   225,
		Name: "Thistle1",
		RGB: rgb{
			R: 255,
			G: 215,
			B: 255,
		},
	},
	color256{
		Id:   226,
		Name: "Yellow1",
		RGB: rgb{
			R: 255,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   227,
		Name: "LightGoldenrod1",
		RGB: rgb{
			R: 255,
			G: 255,
			B: 95,
		},
	},
	color256{
		Id:   228,
		Name: "Khaki1",
		RGB: rgb{
			R: 255,
			G: 255,
			B: 135,
		},
	},
	color256{
		Id:   229,
		Name: "Wheat1",
		RGB: rgb{
			R: 255,
			G: 255,
			B: 175,
		},
	},
	color256{
		Id:   230,
		Name: "Cornsilk1",
		RGB: rgb{
			R: 255,
			G: 255,
			B: 215,
		},
	},
	color256{
		Id:   231,
		Name: "Grey100",
		RGB: rgb{
			R: 255,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   232,
		Name: "Grey3",
		RGB: rgb{
			R: 8,
			G: 8,
			B: 8,
		},
	},
	color256{
		Id:   233,
		Name: "Grey7",
		RGB: rgb{
			R: 18,
			G: 18,
			B: 18,
		},
	},
	color256{
		Id:   234,
		Name: "Grey11",
		RGB: rgb{
			R: 28,
			G: 28,
			B: 28,
		},
	},
	color256{
		Id:   235,
		Name: "Grey15",
		RGB: rgb{
			R: 38,
			G: 38,
			B: 38,
		},
	},
	color256{
		Id:   236,
		Name: "Grey19",
		RGB: rgb{
			R: 48,
			G: 48,
			B: 48,
		},
	},
	color256{
		Id:   237,
		Name: "Grey23",
		RGB: rgb{
			R: 58,
			G: 58,
			B: 58,
		},
	},
	color256{
		Id:   238,
		Name: "Grey27",
		RGB: rgb{
			R: 68,
			G: 68,
			B: 68,
		},
	},
	color256{
		Id:   239,
		Name: "Grey30",
		RGB: rgb{
			R: 78,
			G: 78,
			B: 78,
		},
	},
	color256{
		Id:   240,
		Name: "Grey35",
		RGB: rgb{
			R: 88,
			G: 88,
			B: 88,
		},
	},
	color256{
		Id:   241,
		Name: "Grey39",
		RGB: rgb{
			R: 98,
			G: 98,
			B: 98,
		},
	},
	color256{
		Id:   242,
		Name: "Grey42",
		RGB: rgb{
			R: 108,
			G: 108,
			B: 108,
		},
	},
	color256{
		Id:   243,
		Name: "Grey46",
		RGB: rgb{
			R: 118,
			G: 118,
			B: 118,
		},
	},
	color256{
		Id:   244,
		Name: "Grey50",
		RGB: rgb{
			R: 128,
			G: 128,
			B: 128,
		},
	},
	color256{
		Id:   245,
		Name: "Grey54",
		RGB: rgb{
			R: 138,
			G: 138,
			B: 138,
		},
	},
	color256{
		Id:   246,
		Name: "Grey58",
		RGB: rgb{
			R: 148,
			G: 148,
			B: 148,
		},
	},
	color256{
		Id:   247,
		Name: "Grey62",
		RGB: rgb{
			R: 158,
			G: 158,
			B: 158,
		},
	},
	color256{
		Id:   248,
		Name: "Grey66",
		RGB: rgb{
			R: 168,
			G: 168,
			B: 168,
		},
	},
	color256{
		Id:   249,
		Name: "Grey70",
		RGB: rgb{
			R: 178,
			G: 178,
			B: 178,
		},
	},
	color256{
		Id:   250,
		Name: "Grey74",
		RGB: rgb{
			R: 188,
			G: 188,
			B: 188,
		},
	},
	color256{
		Id:   251,
		Name: "Grey78",
		RGB: rgb{
			R: 198,
			G: 198,
			B: 198,
		},
	},
	color256{
		Id:   252,
		Name: "Grey82",
		RGB: rgb{
			R: 208,
			G: 208,
			B: 208,
		},
	},
	color256{
		Id:   253,
		Name: "Grey85",
		RGB: rgb{
			R: 218,
			G: 218,
			B: 218,
		},
	},
	color256{
		Id:   254,
		Name: "Grey89",
		RGB: rgb{
			R: 228,
			G: 228,
			B: 228,
		},
	},
	color256{
		Id:   255,
		Name: "Grey93",
		RGB: rgb{
			R: 238,
			G: 238,
			B: 238,
		},
	},
}

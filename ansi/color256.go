package ansi

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	hslRegexp      *regexp.Regexp = regexp.MustCompile("^(?i)hsl\\((\\d+),\\s*(\\d+)%,\\s*(\\d+)\\)$")
	rgbRegexp                     = regexp.MustCompile("^(?i)rgb\\((\\d+),\\s*(\\d+),\\s*(\\d+)\\)$")
	shortRGBRegexp                = regexp.MustCompile("rgb(\\d{3})")
)

type hsl struct {
	H    uint16
	S, L uint8
}

func (c *hsl) String() string {
	return fmt.Sprintf("hsl(%d, %d%%, %d%%)", c.H, c.S, c.L)
}

type rgb struct {
	R, G, B uint8
}

func (c *rgb) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", c.R, c.G, c.B)
}

type color256 struct {
	Id uint8
	Name,
	Hex string
	HSL hsl
	RGB rgb
}

func (c color256) FGStart() string {
	return fmt.Sprintf("\x1b[38;5;%dm", c.Id)
}

func (c color256) BGStart() string {
	return fmt.Sprintf("\x1b[48;5;%dm", c.Id)
}

func (c color256) FG(s string) string {
	return fmt.Sprintf("%s%s%s", c.FGStart(), s, FGEnd)
}

func (c color256) BG(s string) string {
	return fmt.Sprintf("%s%s%s", c.BGStart(), s, BGEnd)
}

type colors256 []color256

func (c colors256) Find(what string) (color256, error) {
	i, err := strconv.Atoi(what)
	if err == nil {
		return c.FindId(i)
	}
	if string(what[0]) == "#" {
		return c.FindHex(what)
	}
	if len(what) > 4 {
		switch what[:4] {
		case "rgb(":
			return c.FindRGB(what)
			break
		case "hsl(":
			return c.FindHSL(what)
			break
		}
	}
	return c.FindName(what)
}

func (c colors256) FindRGB(what string) (color256, error) {
	var r, g, b int
	if matches := rgbRegexp.FindAllString(what, -1); len(matches) == 3 {
		r, _ = strconv.Atoi(matches[0])
		g, _ = strconv.Atoi(matches[1])
		b, _ = strconv.Atoi(matches[2])
	}
	for _, col := range c {
		if int(col.RGB.R) == r && int(col.RGB.G) == g && int(col.RGB.B) == b {
			return col, nil
		}
	}
	return color256{}, ColorNotFound
}

func (c colors256) FindHSL(what string) (color256, error) {
	var h, s, l int
	if matches := hslRegexp.FindAllString(what, -1); len(matches) == 3 {
		h, _ = strconv.Atoi(matches[0])
		s, _ = strconv.Atoi(matches[1])
		l, _ = strconv.Atoi(matches[2])
	}
	for _, col := range c {
		if int(col.HSL.H) == h && int(col.HSL.S) == s && int(col.HSL.L) == l {
			return col, nil
		}
	}
	return color256{}, ColorNotFound
}

func (c colors256) FindHex(what string) (color256, error) {
	for _, col := range c {
		if b, _ := regexp.MatchString(fmt.Sprintf("(?i)%s", col.Hex), what); b {
			return col, nil
		}
	}
	return color256{}, ColorNotFound
}

func (c colors256) FindId(what int) (color256, error) {
	if what >= 0 && what < len(c) {
		return c[what], nil
	}
	return color256{}, ColorNotFound
}

func (c colors256) FindName(what string) (color256, error) {
	for _, col := range c {
		if b, _ := regexp.MatchString(fmt.Sprintf("(?i)%s", col.Name), what); b {
			return col, nil
		}
	}
	return color256{}, ColorNotFound
}

var Colors256 colors256 = []color256{
	color256{
		Id:   0,
		Name: "Black",
		Hex:  "#000000",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 0,
		},
		RGB: rgb{
			R: 0,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   1,
		Name: "Maroon",
		Hex:  "#800000",
		HSL: hsl{
			H: 0,
			S: 100,
			L: 25,
		},
		RGB: rgb{
			R: 128,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   2,
		Name: "Green",
		Hex:  "#008000",
		HSL: hsl{
			H: 120,
			S: 100,
			L: 25,
		},
		RGB: rgb{
			R: 0,
			G: 128,
			B: 0,
		},
	},
	color256{
		Id:   3,
		Name: "Olive",
		Hex:  "#808000",
		HSL: hsl{
			H: 60,
			S: 100,
			L: 25,
		},
		RGB: rgb{
			R: 128,
			G: 128,
			B: 0,
		},
	},
	color256{
		Id:   4,
		Name: "Navy",
		Hex:  "#000080",
		HSL: hsl{
			H: 240,
			S: 100,
			L: 25,
		},
		RGB: rgb{
			R: 0,
			G: 0,
			B: 128,
		},
	},
	color256{
		Id:   5,
		Name: "Purple",
		Hex:  "#800080",
		HSL: hsl{
			H: 300,
			S: 100,
			L: 25,
		},
		RGB: rgb{
			R: 128,
			G: 0,
			B: 128,
		},
	},
	color256{
		Id:   6,
		Name: "Teal",
		Hex:  "#008080",
		HSL: hsl{
			H: 180,
			S: 100,
			L: 25,
		},
		RGB: rgb{
			R: 0,
			G: 128,
			B: 128,
		},
	},
	color256{
		Id:   7,
		Name: "Silver",
		Hex:  "#c0c0c0",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 75,
		},
		RGB: rgb{
			R: 192,
			G: 192,
			B: 192,
		},
	},
	color256{
		Id:   8,
		Name: "Grey",
		Hex:  "#808080",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 50,
		},
		RGB: rgb{
			R: 128,
			G: 128,
			B: 128,
		},
	},
	color256{
		Id:   9,
		Name: "Red",
		Hex:  "#ff0000",
		HSL: hsl{
			H: 0,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 255,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   10,
		Name: "Lime",
		Hex:  "#00ff00",
		HSL: hsl{
			H: 120,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 0,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   11,
		Name: "Yellow",
		Hex:  "#ffff00",
		HSL: hsl{
			H: 60,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 255,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   12,
		Name: "Blue",
		Hex:  "#0000ff",
		HSL: hsl{
			H: 240,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 0,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   13,
		Name: "Fuchsia",
		Hex:  "#ff00ff",
		HSL: hsl{
			H: 300,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 255,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   14,
		Name: "Aqua",
		Hex:  "#00ffff",
		HSL: hsl{
			H: 180,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 0,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   15,
		Name: "White",
		Hex:  "#ffffff",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 100,
		},
		RGB: rgb{
			R: 255,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   16,
		Name: "Grey0",
		Hex:  "#000000",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 0,
		},
		RGB: rgb{
			R: 0,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   17,
		Name: "NavyBlue",
		Hex:  "#00005f",
		HSL: hsl{
			H: 240,
			S: 100,
			L: 18,
		},
		RGB: rgb{
			R: 0,
			G: 0,
			B: 95,
		},
	},
	color256{
		Id:   18,
		Name: "DarkBlue",
		Hex:  "#000087",
		HSL: hsl{
			H: 240,
			S: 100,
			L: 26,
		},
		RGB: rgb{
			R: 0,
			G: 0,
			B: 135,
		},
	},
	color256{
		Id:   19,
		Name: "Blue3",
		Hex:  "#0000af",
		HSL: hsl{
			H: 240,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 0,
			G: 0,
			B: 175,
		},
	},
	color256{
		Id:   20,
		Name: "Blue3",
		Hex:  "#0000d7",
		HSL: hsl{
			H: 240,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 0,
			G: 0,
			B: 215,
		},
	},
	color256{
		Id:   21,
		Name: "Blue1",
		Hex:  "#0000ff",
		HSL: hsl{
			H: 240,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 0,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   22,
		Name: "DarkGreen",
		Hex:  "#005f00",
		HSL: hsl{
			H: 120,
			S: 100,
			L: 18,
		},
		RGB: rgb{
			R: 0,
			G: 95,
			B: 0,
		},
	},
	color256{
		Id:   23,
		Name: "DeepSkyBlue4",
		Hex:  "#005f5f",
		HSL: hsl{
			H: 180,
			S: 100,
			L: 18,
		},
		RGB: rgb{
			R: 0,
			G: 95,
			B: 95,
		},
	},
	color256{
		Id:   24,
		Name: "DeepSkyBlue4",
		Hex:  "#005f87",
		HSL: hsl{
			H: 197,
			S: 100,
			L: 26,
		},
		RGB: rgb{
			R: 0,
			G: 95,
			B: 135,
		},
	},
	color256{
		Id:   25,
		Name: "DeepSkyBlue4",
		Hex:  "#005faf",
		HSL: hsl{
			H: 207,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 0,
			G: 95,
			B: 175,
		},
	},
	color256{
		Id:   26,
		Name: "DodgerBlue3",
		Hex:  "#005fd7",
		HSL: hsl{
			H: 213,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 0,
			G: 95,
			B: 215,
		},
	},
	color256{
		Id:   27,
		Name: "DodgerBlue2",
		Hex:  "#005fff",
		HSL: hsl{
			H: 217,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 0,
			G: 95,
			B: 255,
		},
	},
	color256{
		Id:   28,
		Name: "Green4",
		Hex:  "#008700",
		HSL: hsl{
			H: 120,
			S: 100,
			L: 26,
		},
		RGB: rgb{
			R: 0,
			G: 135,
			B: 0,
		},
	},
	color256{
		Id:   29,
		Name: "SpringGreen4",
		Hex:  "#00875f",
		HSL: hsl{
			H: 162,
			S: 100,
			L: 26,
		},
		RGB: rgb{
			R: 0,
			G: 135,
			B: 95,
		},
	},
	color256{
		Id:   30,
		Name: "Turquoise4",
		Hex:  "#008787",
		HSL: hsl{
			H: 180,
			S: 100,
			L: 26,
		},
		RGB: rgb{
			R: 0,
			G: 135,
			B: 135,
		},
	},
	color256{
		Id:   31,
		Name: "DeepSkyBlue3",
		Hex:  "#0087af",
		HSL: hsl{
			H: 193,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 0,
			G: 135,
			B: 175,
		},
	},
	color256{
		Id:   32,
		Name: "DeepSkyBlue3",
		Hex:  "#0087d7",
		HSL: hsl{
			H: 202,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 0,
			G: 135,
			B: 215,
		},
	},
	color256{
		Id:   33,
		Name: "DodgerBlue1",
		Hex:  "#0087ff",
		HSL: hsl{
			H: 208,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 0,
			G: 135,
			B: 255,
		},
	},
	color256{
		Id:   34,
		Name: "Green3",
		Hex:  "#00af00",
		HSL: hsl{
			H: 120,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 0,
			G: 175,
			B: 0,
		},
	},
	color256{
		Id:   35,
		Name: "SpringGreen3",
		Hex:  "#00af5f",
		HSL: hsl{
			H: 152,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 0,
			G: 175,
			B: 95,
		},
	},
	color256{
		Id:   36,
		Name: "DarkCyan",
		Hex:  "#00af87",
		HSL: hsl{
			H: 166,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 0,
			G: 175,
			B: 135,
		},
	},
	color256{
		Id:   37,
		Name: "LightSeaGreen",
		Hex:  "#00afaf",
		HSL: hsl{
			H: 180,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 0,
			G: 175,
			B: 175,
		},
	},
	color256{
		Id:   38,
		Name: "DeepSkyBlue2",
		Hex:  "#00afd7",
		HSL: hsl{
			H: 191,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 0,
			G: 175,
			B: 215,
		},
	},
	color256{
		Id:   39,
		Name: "DeepSkyBlue1",
		Hex:  "#00afff",
		HSL: hsl{
			H: 198,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 0,
			G: 175,
			B: 255,
		},
	},
	color256{
		Id:   40,
		Name: "Green3",
		Hex:  "#00d700",
		HSL: hsl{
			H: 120,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 0,
			G: 215,
			B: 0,
		},
	},
	color256{
		Id:   41,
		Name: "SpringGreen3",
		Hex:  "#00d75f",
		HSL: hsl{
			H: 146,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 0,
			G: 215,
			B: 95,
		},
	},
	color256{
		Id:   42,
		Name: "SpringGreen2",
		Hex:  "#00d787",
		HSL: hsl{
			H: 157,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 0,
			G: 215,
			B: 135,
		},
	},
	color256{
		Id:   43,
		Name: "Cyan3",
		Hex:  "#00d7af",
		HSL: hsl{
			H: 168,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 0,
			G: 215,
			B: 175,
		},
	},
	color256{
		Id:   44,
		Name: "DarkTurquoise",
		Hex:  "#00d7d7",
		HSL: hsl{
			H: 180,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 0,
			G: 215,
			B: 215,
		},
	},
	color256{
		Id:   45,
		Name: "Turquoise2",
		Hex:  "#00d7ff",
		HSL: hsl{
			H: 189,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 0,
			G: 215,
			B: 255,
		},
	},
	color256{
		Id:   46,
		Name: "Green1",
		Hex:  "#00ff00",
		HSL: hsl{
			H: 120,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 0,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   47,
		Name: "SpringGreen2",
		Hex:  "#00ff5f",
		HSL: hsl{
			H: 142,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 0,
			G: 255,
			B: 95,
		},
	},
	color256{
		Id:   48,
		Name: "SpringGreen1",
		Hex:  "#00ff87",
		HSL: hsl{
			H: 151,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 0,
			G: 255,
			B: 135,
		},
	},
	color256{
		Id:   49,
		Name: "MediumSpringGreen",
		Hex:  "#00ffaf",
		HSL: hsl{
			H: 161,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 0,
			G: 255,
			B: 175,
		},
	},
	color256{
		Id:   50,
		Name: "Cyan2",
		Hex:  "#00ffd7",
		HSL: hsl{
			H: 170,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 0,
			G: 255,
			B: 215,
		},
	},
	color256{
		Id:   51,
		Name: "Cyan1",
		Hex:  "#00ffff",
		HSL: hsl{
			H: 180,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 0,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   52,
		Name: "DarkRed",
		Hex:  "#5f0000",
		HSL: hsl{
			H: 0,
			S: 100,
			L: 18,
		},
		RGB: rgb{
			R: 95,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   53,
		Name: "DeepPink4",
		Hex:  "#5f005f",
		HSL: hsl{
			H: 300,
			S: 100,
			L: 18,
		},
		RGB: rgb{
			R: 95,
			G: 0,
			B: 95,
		},
	},
	color256{
		Id:   54,
		Name: "Purple4",
		Hex:  "#5f0087",
		HSL: hsl{
			H: 282,
			S: 100,
			L: 26,
		},
		RGB: rgb{
			R: 95,
			G: 0,
			B: 135,
		},
	},
	color256{
		Id:   55,
		Name: "Purple4",
		Hex:  "#5f00af",
		HSL: hsl{
			H: 272,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 95,
			G: 0,
			B: 175,
		},
	},
	color256{
		Id:   56,
		Name: "Purple3",
		Hex:  "#5f00d7",
		HSL: hsl{
			H: 266,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 95,
			G: 0,
			B: 215,
		},
	},
	color256{
		Id:   57,
		Name: "BlueViolet",
		Hex:  "#5f00ff",
		HSL: hsl{
			H: 262,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 95,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   58,
		Name: "Orange4",
		Hex:  "#5f5f00",
		HSL: hsl{
			H: 60,
			S: 100,
			L: 18,
		},
		RGB: rgb{
			R: 95,
			G: 95,
			B: 0,
		},
	},
	color256{
		Id:   59,
		Name: "Grey37",
		Hex:  "#5f5f5f",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 37,
		},
		RGB: rgb{
			R: 95,
			G: 95,
			B: 95,
		},
	},
	color256{
		Id:   60,
		Name: "MediumPurple4",
		Hex:  "#5f5f87",
		HSL: hsl{
			H: 240,
			S: 17,
			L: 45,
		},
		RGB: rgb{
			R: 95,
			G: 95,
			B: 135,
		},
	},
	color256{
		Id:   61,
		Name: "SlateBlue3",
		Hex:  "#5f5faf",
		HSL: hsl{
			H: 240,
			S: 33,
			L: 52,
		},
		RGB: rgb{
			R: 95,
			G: 95,
			B: 175,
		},
	},
	color256{
		Id:   62,
		Name: "SlateBlue3",
		Hex:  "#5f5fd7",
		HSL: hsl{
			H: 240,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 95,
			G: 95,
			B: 215,
		},
	},
	color256{
		Id:   63,
		Name: "RoyalBlue1",
		Hex:  "#5f5fff",
		HSL: hsl{
			H: 240,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 95,
			G: 95,
			B: 255,
		},
	},
	color256{
		Id:   64,
		Name: "Chartreuse4",
		Hex:  "#5f8700",
		HSL: hsl{
			H: 77,
			S: 100,
			L: 26,
		},
		RGB: rgb{
			R: 95,
			G: 135,
			B: 0,
		},
	},
	color256{
		Id:   65,
		Name: "DarkSeaGreen4",
		Hex:  "#5f875f",
		HSL: hsl{
			H: 120,
			S: 17,
			L: 45,
		},
		RGB: rgb{
			R: 95,
			G: 135,
			B: 95,
		},
	},
	color256{
		Id:   66,
		Name: "PaleTurquoise4",
		Hex:  "#5f8787",
		HSL: hsl{
			H: 180,
			S: 17,
			L: 45,
		},
		RGB: rgb{
			R: 95,
			G: 135,
			B: 135,
		},
	},
	color256{
		Id:   67,
		Name: "SteelBlue",
		Hex:  "#5f87af",
		HSL: hsl{
			H: 210,
			S: 33,
			L: 52,
		},
		RGB: rgb{
			R: 95,
			G: 135,
			B: 175,
		},
	},
	color256{
		Id:   68,
		Name: "SteelBlue3",
		Hex:  "#5f87d7",
		HSL: hsl{
			H: 220,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 95,
			G: 135,
			B: 215,
		},
	},
	color256{
		Id:   69,
		Name: "CornflowerBlue",
		Hex:  "#5f87ff",
		HSL: hsl{
			H: 225,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 95,
			G: 135,
			B: 255,
		},
	},
	color256{
		Id:   70,
		Name: "Chartreuse3",
		Hex:  "#5faf00",
		HSL: hsl{
			H: 87,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 95,
			G: 175,
			B: 0,
		},
	},
	color256{
		Id:   71,
		Name: "DarkSeaGreen4",
		Hex:  "#5faf5f",
		HSL: hsl{
			H: 120,
			S: 33,
			L: 52,
		},
		RGB: rgb{
			R: 95,
			G: 175,
			B: 95,
		},
	},
	color256{
		Id:   72,
		Name: "CadetBlue",
		Hex:  "#5faf87",
		HSL: hsl{
			H: 150,
			S: 33,
			L: 52,
		},
		RGB: rgb{
			R: 95,
			G: 175,
			B: 135,
		},
	},
	color256{
		Id:   73,
		Name: "CadetBlue",
		Hex:  "#5fafaf",
		HSL: hsl{
			H: 180,
			S: 33,
			L: 52,
		},
		RGB: rgb{
			R: 95,
			G: 175,
			B: 175,
		},
	},
	color256{
		Id:   74,
		Name: "SkyBlue3",
		Hex:  "#5fafd7",
		HSL: hsl{
			H: 200,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 95,
			G: 175,
			B: 215,
		},
	},
	color256{
		Id:   75,
		Name: "SteelBlue1",
		Hex:  "#5fafff",
		HSL: hsl{
			H: 210,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 95,
			G: 175,
			B: 255,
		},
	},
	color256{
		Id:   76,
		Name: "Chartreuse3",
		Hex:  "#5fd700",
		HSL: hsl{
			H: 93,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 95,
			G: 215,
			B: 0,
		},
	},
	color256{
		Id:   77,
		Name: "PaleGreen3",
		Hex:  "#5fd75f",
		HSL: hsl{
			H: 120,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 95,
			G: 215,
			B: 95,
		},
	},
	color256{
		Id:   78,
		Name: "SeaGreen3",
		Hex:  "#5fd787",
		HSL: hsl{
			H: 140,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 95,
			G: 215,
			B: 135,
		},
	},
	color256{
		Id:   79,
		Name: "Aquamarine3",
		Hex:  "#5fd7af",
		HSL: hsl{
			H: 160,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 95,
			G: 215,
			B: 175,
		},
	},
	color256{
		Id:   80,
		Name: "MediumTurquoise",
		Hex:  "#5fd7d7",
		HSL: hsl{
			H: 180,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 95,
			G: 215,
			B: 215,
		},
	},
	color256{
		Id:   81,
		Name: "SteelBlue1",
		Hex:  "#5fd7ff",
		HSL: hsl{
			H: 195,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 95,
			G: 215,
			B: 255,
		},
	},
	color256{
		Id:   82,
		Name: "Chartreuse2",
		Hex:  "#5fff00",
		HSL: hsl{
			H: 97,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 95,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   83,
		Name: "SeaGreen2",
		Hex:  "#5fff5f",
		HSL: hsl{
			H: 120,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 95,
			G: 255,
			B: 95,
		},
	},
	color256{
		Id:   84,
		Name: "SeaGreen1",
		Hex:  "#5fff87",
		HSL: hsl{
			H: 135,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 95,
			G: 255,
			B: 135,
		},
	},
	color256{
		Id:   85,
		Name: "SeaGreen1",
		Hex:  "#5fffaf",
		HSL: hsl{
			H: 150,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 95,
			G: 255,
			B: 175,
		},
	},
	color256{
		Id:   86,
		Name: "Aquamarine1",
		Hex:  "#5fffd7",
		HSL: hsl{
			H: 165,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 95,
			G: 255,
			B: 215,
		},
	},
	color256{
		Id:   87,
		Name: "DarkSlateGray2",
		Hex:  "#5fffff",
		HSL: hsl{
			H: 180,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 95,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   88,
		Name: "DarkRed",
		Hex:  "#870000",
		HSL: hsl{
			H: 0,
			S: 100,
			L: 26,
		},
		RGB: rgb{
			R: 135,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   89,
		Name: "DeepPink4",
		Hex:  "#87005f",
		HSL: hsl{
			H: 317,
			S: 100,
			L: 26,
		},
		RGB: rgb{
			R: 135,
			G: 0,
			B: 95,
		},
	},
	color256{
		Id:   90,
		Name: "DarkMagenta",
		Hex:  "#870087",
		HSL: hsl{
			H: 300,
			S: 100,
			L: 26,
		},
		RGB: rgb{
			R: 135,
			G: 0,
			B: 135,
		},
	},
	color256{
		Id:   91,
		Name: "DarkMagenta",
		Hex:  "#8700af",
		HSL: hsl{
			H: 286,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 135,
			G: 0,
			B: 175,
		},
	},
	color256{
		Id:   92,
		Name: "DarkViolet",
		Hex:  "#8700d7",
		HSL: hsl{
			H: 277,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 135,
			G: 0,
			B: 215,
		},
	},
	color256{
		Id:   93,
		Name: "Purple",
		Hex:  "#8700ff",
		HSL: hsl{
			H: 271,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 135,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   94,
		Name: "Orange4",
		Hex:  "#875f00",
		HSL: hsl{
			H: 42,
			S: 100,
			L: 26,
		},
		RGB: rgb{
			R: 135,
			G: 95,
			B: 0,
		},
	},
	color256{
		Id:   95,
		Name: "LightPink4",
		Hex:  "#875f5f",
		HSL: hsl{
			H: 0,
			S: 17,
			L: 45,
		},
		RGB: rgb{
			R: 135,
			G: 95,
			B: 95,
		},
	},
	color256{
		Id:   96,
		Name: "Plum4",
		Hex:  "#875f87",
		HSL: hsl{
			H: 300,
			S: 17,
			L: 45,
		},
		RGB: rgb{
			R: 135,
			G: 95,
			B: 135,
		},
	},
	color256{
		Id:   97,
		Name: "MediumPurple3",
		Hex:  "#875faf",
		HSL: hsl{
			H: 270,
			S: 33,
			L: 52,
		},
		RGB: rgb{
			R: 135,
			G: 95,
			B: 175,
		},
	},
	color256{
		Id:   98,
		Name: "MediumPurple3",
		Hex:  "#875fd7",
		HSL: hsl{
			H: 260,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 135,
			G: 95,
			B: 215,
		},
	},
	color256{
		Id:   99,
		Name: "SlateBlue1",
		Hex:  "#875fff",
		HSL: hsl{
			H: 255,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 135,
			G: 95,
			B: 255,
		},
	},
	color256{
		Id:   100,
		Name: "Yellow4",
		Hex:  "#878700",
		HSL: hsl{
			H: 60,
			S: 100,
			L: 26,
		},
		RGB: rgb{
			R: 135,
			G: 135,
			B: 0,
		},
	},
	color256{
		Id:   101,
		Name: "Wheat4",
		Hex:  "#87875f",
		HSL: hsl{
			H: 60,
			S: 17,
			L: 45,
		},
		RGB: rgb{
			R: 135,
			G: 135,
			B: 95,
		},
	},
	color256{
		Id:   102,
		Name: "Grey53",
		Hex:  "#878787",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 52,
		},
		RGB: rgb{
			R: 135,
			G: 135,
			B: 135,
		},
	},
	color256{
		Id:   103,
		Name: "LightSlateGrey",
		Hex:  "#8787af",
		HSL: hsl{
			H: 240,
			S: 20,
			L: 60,
		},
		RGB: rgb{
			R: 135,
			G: 135,
			B: 175,
		},
	},
	color256{
		Id:   104,
		Name: "MediumPurple",
		Hex:  "#8787d7",
		HSL: hsl{
			H: 240,
			S: 50,
			L: 68,
		},
		RGB: rgb{
			R: 135,
			G: 135,
			B: 215,
		},
	},
	color256{
		Id:   105,
		Name: "LightSlateBlue",
		Hex:  "#8787ff",
		HSL: hsl{
			H: 240,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 135,
			G: 135,
			B: 255,
		},
	},
	color256{
		Id:   106,
		Name: "Yellow4",
		Hex:  "#87af00",
		HSL: hsl{
			H: 73,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 135,
			G: 175,
			B: 0,
		},
	},
	color256{
		Id:   107,
		Name: "DarkOliveGreen3",
		Hex:  "#87af5f",
		HSL: hsl{
			H: 90,
			S: 33,
			L: 52,
		},
		RGB: rgb{
			R: 135,
			G: 175,
			B: 95,
		},
	},
	color256{
		Id:   108,
		Name: "DarkSeaGreen",
		Hex:  "#87af87",
		HSL: hsl{
			H: 120,
			S: 20,
			L: 60,
		},
		RGB: rgb{
			R: 135,
			G: 175,
			B: 135,
		},
	},
	color256{
		Id:   109,
		Name: "LightSkyBlue3",
		Hex:  "#87afaf",
		HSL: hsl{
			H: 180,
			S: 20,
			L: 60,
		},
		RGB: rgb{
			R: 135,
			G: 175,
			B: 175,
		},
	},
	color256{
		Id:   110,
		Name: "LightSkyBlue3",
		Hex:  "#87afd7",
		HSL: hsl{
			H: 210,
			S: 50,
			L: 68,
		},
		RGB: rgb{
			R: 135,
			G: 175,
			B: 215,
		},
	},
	color256{
		Id:   111,
		Name: "SkyBlue2",
		Hex:  "#87afff",
		HSL: hsl{
			H: 220,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 135,
			G: 175,
			B: 255,
		},
	},
	color256{
		Id:   112,
		Name: "Chartreuse2",
		Hex:  "#87d700",
		HSL: hsl{
			H: 82,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 135,
			G: 215,
			B: 0,
		},
	},
	color256{
		Id:   113,
		Name: "DarkOliveGreen3",
		Hex:  "#87d75f",
		HSL: hsl{
			H: 100,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 135,
			G: 215,
			B: 95,
		},
	},
	color256{
		Id:   114,
		Name: "PaleGreen3",
		Hex:  "#87d787",
		HSL: hsl{
			H: 120,
			S: 50,
			L: 68,
		},
		RGB: rgb{
			R: 135,
			G: 215,
			B: 135,
		},
	},
	color256{
		Id:   115,
		Name: "DarkSeaGreen3",
		Hex:  "#87d7af",
		HSL: hsl{
			H: 150,
			S: 50,
			L: 68,
		},
		RGB: rgb{
			R: 135,
			G: 215,
			B: 175,
		},
	},
	color256{
		Id:   116,
		Name: "DarkSlateGray3",
		Hex:  "#87d7d7",
		HSL: hsl{
			H: 180,
			S: 50,
			L: 68,
		},
		RGB: rgb{
			R: 135,
			G: 215,
			B: 215,
		},
	},
	color256{
		Id:   117,
		Name: "SkyBlue1",
		Hex:  "#87d7ff",
		HSL: hsl{
			H: 200,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 135,
			G: 215,
			B: 255,
		},
	},
	color256{
		Id:   118,
		Name: "Chartreuse1",
		Hex:  "#87ff00",
		HSL: hsl{
			H: 88,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 135,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   119,
		Name: "LightGreen",
		Hex:  "#87ff5f",
		HSL: hsl{
			H: 105,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 135,
			G: 255,
			B: 95,
		},
	},
	color256{
		Id:   120,
		Name: "LightGreen",
		Hex:  "#87ff87",
		HSL: hsl{
			H: 120,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 135,
			G: 255,
			B: 135,
		},
	},
	color256{
		Id:   121,
		Name: "PaleGreen1",
		Hex:  "#87ffaf",
		HSL: hsl{
			H: 140,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 135,
			G: 255,
			B: 175,
		},
	},
	color256{
		Id:   122,
		Name: "Aquamarine1",
		Hex:  "#87ffd7",
		HSL: hsl{
			H: 160,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 135,
			G: 255,
			B: 215,
		},
	},
	color256{
		Id:   123,
		Name: "DarkSlateGray1",
		Hex:  "#87ffff",
		HSL: hsl{
			H: 180,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 135,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   124,
		Name: "Red3",
		Hex:  "#af0000",
		HSL: hsl{
			H: 0,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 175,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   125,
		Name: "DeepPink4",
		Hex:  "#af005f",
		HSL: hsl{
			H: 327,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 175,
			G: 0,
			B: 95,
		},
	},
	color256{
		Id:   126,
		Name: "MediumVioletRed",
		Hex:  "#af0087",
		HSL: hsl{
			H: 313,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 175,
			G: 0,
			B: 135,
		},
	},
	color256{
		Id:   127,
		Name: "Magenta3",
		Hex:  "#af00af",
		HSL: hsl{
			H: 300,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 175,
			G: 0,
			B: 175,
		},
	},
	color256{
		Id:   128,
		Name: "DarkViolet",
		Hex:  "#af00d7",
		HSL: hsl{
			H: 288,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 175,
			G: 0,
			B: 215,
		},
	},
	color256{
		Id:   129,
		Name: "Purple",
		Hex:  "#af00ff",
		HSL: hsl{
			H: 281,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 175,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   130,
		Name: "DarkOrange3",
		Hex:  "#af5f00",
		HSL: hsl{
			H: 32,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 175,
			G: 95,
			B: 0,
		},
	},
	color256{
		Id:   131,
		Name: "IndianRed",
		Hex:  "#af5f5f",
		HSL: hsl{
			H: 0,
			S: 33,
			L: 52,
		},
		RGB: rgb{
			R: 175,
			G: 95,
			B: 95,
		},
	},
	color256{
		Id:   132,
		Name: "HotPink3",
		Hex:  "#af5f87",
		HSL: hsl{
			H: 330,
			S: 33,
			L: 52,
		},
		RGB: rgb{
			R: 175,
			G: 95,
			B: 135,
		},
	},
	color256{
		Id:   133,
		Name: "MediumOrchId3",
		Hex:  "#af5faf",
		HSL: hsl{
			H: 300,
			S: 33,
			L: 52,
		},
		RGB: rgb{
			R: 175,
			G: 95,
			B: 175,
		},
	},
	color256{
		Id:   134,
		Name: "MediumOrchId",
		Hex:  "#af5fd7",
		HSL: hsl{
			H: 280,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 175,
			G: 95,
			B: 215,
		},
	},
	color256{
		Id:   135,
		Name: "MediumPurple2",
		Hex:  "#af5fff",
		HSL: hsl{
			H: 270,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 175,
			G: 95,
			B: 255,
		},
	},
	color256{
		Id:   136,
		Name: "DarkGoldenrod",
		Hex:  "#af8700",
		HSL: hsl{
			H: 46,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 175,
			G: 135,
			B: 0,
		},
	},
	color256{
		Id:   137,
		Name: "LightSalmon3",
		Hex:  "#af875f",
		HSL: hsl{
			H: 30,
			S: 33,
			L: 52,
		},
		RGB: rgb{
			R: 175,
			G: 135,
			B: 95,
		},
	},
	color256{
		Id:   138,
		Name: "RosyBrown",
		Hex:  "#af8787",
		HSL: hsl{
			H: 0,
			S: 20,
			L: 60,
		},
		RGB: rgb{
			R: 175,
			G: 135,
			B: 135,
		},
	},
	color256{
		Id:   139,
		Name: "Grey63",
		Hex:  "#af87af",
		HSL: hsl{
			H: 300,
			S: 20,
			L: 60,
		},
		RGB: rgb{
			R: 175,
			G: 135,
			B: 175,
		},
	},
	color256{
		Id:   140,
		Name: "MediumPurple2",
		Hex:  "#af87d7",
		HSL: hsl{
			H: 270,
			S: 50,
			L: 68,
		},
		RGB: rgb{
			R: 175,
			G: 135,
			B: 215,
		},
	},
	color256{
		Id:   141,
		Name: "MediumPurple1",
		Hex:  "#af87ff",
		HSL: hsl{
			H: 260,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 175,
			G: 135,
			B: 255,
		},
	},
	color256{
		Id:   142,
		Name: "Gold3",
		Hex:  "#afaf00",
		HSL: hsl{
			H: 60,
			S: 100,
			L: 34,
		},
		RGB: rgb{
			R: 175,
			G: 175,
			B: 0,
		},
	},
	color256{
		Id:   143,
		Name: "DarkKhaki",
		Hex:  "#afaf5f",
		HSL: hsl{
			H: 60,
			S: 33,
			L: 52,
		},
		RGB: rgb{
			R: 175,
			G: 175,
			B: 95,
		},
	},
	color256{
		Id:   144,
		Name: "NavajoWhite3",
		Hex:  "#afaf87",
		HSL: hsl{
			H: 60,
			S: 20,
			L: 60,
		},
		RGB: rgb{
			R: 175,
			G: 175,
			B: 135,
		},
	},
	color256{
		Id:   145,
		Name: "Grey69",
		Hex:  "#afafaf",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 68,
		},
		RGB: rgb{
			R: 175,
			G: 175,
			B: 175,
		},
	},
	color256{
		Id:   146,
		Name: "LightSteelBlue3",
		Hex:  "#afafd7",
		HSL: hsl{
			H: 240,
			S: 33,
			L: 76,
		},
		RGB: rgb{
			R: 175,
			G: 175,
			B: 215,
		},
	},
	color256{
		Id:   147,
		Name: "LightSteelBlue",
		Hex:  "#afafff",
		HSL: hsl{
			H: 240,
			S: 100,
			L: 84,
		},
		RGB: rgb{
			R: 175,
			G: 175,
			B: 255,
		},
	},
	color256{
		Id:   148,
		Name: "Yellow3",
		Hex:  "#afd700",
		HSL: hsl{
			H: 71,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 175,
			G: 215,
			B: 0,
		},
	},
	color256{
		Id:   149,
		Name: "DarkOliveGreen3",
		Hex:  "#afd75f",
		HSL: hsl{
			H: 80,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 175,
			G: 215,
			B: 95,
		},
	},
	color256{
		Id:   150,
		Name: "DarkSeaGreen3",
		Hex:  "#afd787",
		HSL: hsl{
			H: 90,
			S: 50,
			L: 68,
		},
		RGB: rgb{
			R: 175,
			G: 215,
			B: 135,
		},
	},
	color256{
		Id:   151,
		Name: "DarkSeaGreen2",
		Hex:  "#afd7af",
		HSL: hsl{
			H: 120,
			S: 33,
			L: 76,
		},
		RGB: rgb{
			R: 175,
			G: 215,
			B: 175,
		},
	},
	color256{
		Id:   152,
		Name: "LightCyan3",
		Hex:  "#afd7d7",
		HSL: hsl{
			H: 180,
			S: 33,
			L: 76,
		},
		RGB: rgb{
			R: 175,
			G: 215,
			B: 215,
		},
	},
	color256{
		Id:   153,
		Name: "LightSkyBlue1",
		Hex:  "#afd7ff",
		HSL: hsl{
			H: 210,
			S: 100,
			L: 84,
		},
		RGB: rgb{
			R: 175,
			G: 215,
			B: 255,
		},
	},
	color256{
		Id:   154,
		Name: "GreenYellow",
		Hex:  "#afff00",
		HSL: hsl{
			H: 78,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 175,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   155,
		Name: "DarkOliveGreen2",
		Hex:  "#afff5f",
		HSL: hsl{
			H: 90,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 175,
			G: 255,
			B: 95,
		},
	},
	color256{
		Id:   156,
		Name: "PaleGreen1",
		Hex:  "#afff87",
		HSL: hsl{
			H: 100,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 175,
			G: 255,
			B: 135,
		},
	},
	color256{
		Id:   157,
		Name: "DarkSeaGreen2",
		Hex:  "#afffaf",
		HSL: hsl{
			H: 120,
			S: 100,
			L: 84,
		},
		RGB: rgb{
			R: 175,
			G: 255,
			B: 175,
		},
	},
	color256{
		Id:   158,
		Name: "DarkSeaGreen1",
		Hex:  "#afffd7",
		HSL: hsl{
			H: 150,
			S: 100,
			L: 84,
		},
		RGB: rgb{
			R: 175,
			G: 255,
			B: 215,
		},
	},
	color256{
		Id:   159,
		Name: "PaleTurquoise1",
		Hex:  "#afffff",
		HSL: hsl{
			H: 180,
			S: 100,
			L: 84,
		},
		RGB: rgb{
			R: 175,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   160,
		Name: "Red3",
		Hex:  "#d70000",
		HSL: hsl{
			H: 0,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 215,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   161,
		Name: "DeepPink3",
		Hex:  "#d7005f",
		HSL: hsl{
			H: 333,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 215,
			G: 0,
			B: 95,
		},
	},
	color256{
		Id:   162,
		Name: "DeepPink3",
		Hex:  "#d70087",
		HSL: hsl{
			H: 322,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 215,
			G: 0,
			B: 135,
		},
	},
	color256{
		Id:   163,
		Name: "Magenta3",
		Hex:  "#d700af",
		HSL: hsl{
			H: 311,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 215,
			G: 0,
			B: 175,
		},
	},
	color256{
		Id:   164,
		Name: "Magenta3",
		Hex:  "#d700d7",
		HSL: hsl{
			H: 300,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 215,
			G: 0,
			B: 215,
		},
	},
	color256{
		Id:   165,
		Name: "Magenta2",
		Hex:  "#d700ff",
		HSL: hsl{
			H: 290,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 215,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   166,
		Name: "DarkOrange3",
		Hex:  "#d75f00",
		HSL: hsl{
			H: 26,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 215,
			G: 95,
			B: 0,
		},
	},
	color256{
		Id:   167,
		Name: "IndianRed",
		Hex:  "#d75f5f",
		HSL: hsl{
			H: 0,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 215,
			G: 95,
			B: 95,
		},
	},
	color256{
		Id:   168,
		Name: "HotPink3",
		Hex:  "#d75f87",
		HSL: hsl{
			H: 340,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 215,
			G: 95,
			B: 135,
		},
	},
	color256{
		Id:   169,
		Name: "HotPink2",
		Hex:  "#d75faf",
		HSL: hsl{
			H: 320,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 215,
			G: 95,
			B: 175,
		},
	},
	color256{
		Id:   170,
		Name: "OrchId",
		Hex:  "#d75fd7",
		HSL: hsl{
			H: 300,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 215,
			G: 95,
			B: 215,
		},
	},
	color256{
		Id:   171,
		Name: "MediumOrchId1",
		Hex:  "#d75fff",
		HSL: hsl{
			H: 285,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 215,
			G: 95,
			B: 255,
		},
	},
	color256{
		Id:   172,
		Name: "Orange3",
		Hex:  "#d78700",
		HSL: hsl{
			H: 37,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 215,
			G: 135,
			B: 0,
		},
	},
	color256{
		Id:   173,
		Name: "LightSalmon3",
		Hex:  "#d7875f",
		HSL: hsl{
			H: 20,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 215,
			G: 135,
			B: 95,
		},
	},
	color256{
		Id:   174,
		Name: "LightPink3",
		Hex:  "#d78787",
		HSL: hsl{
			H: 0,
			S: 50,
			L: 68,
		},
		RGB: rgb{
			R: 215,
			G: 135,
			B: 135,
		},
	},
	color256{
		Id:   175,
		Name: "Pink3",
		Hex:  "#d787af",
		HSL: hsl{
			H: 330,
			S: 50,
			L: 68,
		},
		RGB: rgb{
			R: 215,
			G: 135,
			B: 175,
		},
	},
	color256{
		Id:   176,
		Name: "Plum3",
		Hex:  "#d787d7",
		HSL: hsl{
			H: 300,
			S: 50,
			L: 68,
		},
		RGB: rgb{
			R: 215,
			G: 135,
			B: 215,
		},
	},
	color256{
		Id:   177,
		Name: "Violet",
		Hex:  "#d787ff",
		HSL: hsl{
			H: 280,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 215,
			G: 135,
			B: 255,
		},
	},
	color256{
		Id:   178,
		Name: "Gold3",
		Hex:  "#d7af00",
		HSL: hsl{
			H: 48,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 215,
			G: 175,
			B: 0,
		},
	},
	color256{
		Id:   179,
		Name: "LightGoldenrod3",
		Hex:  "#d7af5f",
		HSL: hsl{
			H: 40,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 215,
			G: 175,
			B: 95,
		},
	},
	color256{
		Id:   180,
		Name: "Tan",
		Hex:  "#d7af87",
		HSL: hsl{
			H: 30,
			S: 50,
			L: 68,
		},
		RGB: rgb{
			R: 215,
			G: 175,
			B: 135,
		},
	},
	color256{
		Id:   181,
		Name: "MistyRose3",
		Hex:  "#d7afaf",
		HSL: hsl{
			H: 0,
			S: 33,
			L: 76,
		},
		RGB: rgb{
			R: 215,
			G: 175,
			B: 175,
		},
	},
	color256{
		Id:   182,
		Name: "Thistle3",
		Hex:  "#d7afd7",
		HSL: hsl{
			H: 300,
			S: 33,
			L: 76,
		},
		RGB: rgb{
			R: 215,
			G: 175,
			B: 215,
		},
	},
	color256{
		Id:   183,
		Name: "Plum2",
		Hex:  "#d7afff",
		HSL: hsl{
			H: 270,
			S: 100,
			L: 84,
		},
		RGB: rgb{
			R: 215,
			G: 175,
			B: 255,
		},
	},
	color256{
		Id:   184,
		Name: "Yellow3",
		Hex:  "#d7d700",
		HSL: hsl{
			H: 60,
			S: 100,
			L: 42,
		},
		RGB: rgb{
			R: 215,
			G: 215,
			B: 0,
		},
	},
	color256{
		Id:   185,
		Name: "Khaki3",
		Hex:  "#d7d75f",
		HSL: hsl{
			H: 60,
			S: 60,
			L: 60,
		},
		RGB: rgb{
			R: 215,
			G: 215,
			B: 95,
		},
	},
	color256{
		Id:   186,
		Name: "LightGoldenrod2",
		Hex:  "#d7d787",
		HSL: hsl{
			H: 60,
			S: 50,
			L: 68,
		},
		RGB: rgb{
			R: 215,
			G: 215,
			B: 135,
		},
	},
	color256{
		Id:   187,
		Name: "LightYellow3",
		Hex:  "#d7d7af",
		HSL: hsl{
			H: 60,
			S: 33,
			L: 76,
		},
		RGB: rgb{
			R: 215,
			G: 215,
			B: 175,
		},
	},
	color256{
		Id:   188,
		Name: "Grey84",
		Hex:  "#d7d7d7",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 84,
		},
		RGB: rgb{
			R: 215,
			G: 215,
			B: 215,
		},
	},
	color256{
		Id:   189,
		Name: "LightSteelBlue1",
		Hex:  "#d7d7ff",
		HSL: hsl{
			H: 240,
			S: 100,
			L: 92,
		},
		RGB: rgb{
			R: 215,
			G: 215,
			B: 255,
		},
	},
	color256{
		Id:   190,
		Name: "Yellow2",
		Hex:  "#d7ff00",
		HSL: hsl{
			H: 69,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 215,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   191,
		Name: "DarkOliveGreen1",
		Hex:  "#d7ff5f",
		HSL: hsl{
			H: 75,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 215,
			G: 255,
			B: 95,
		},
	},
	color256{
		Id:   192,
		Name: "DarkOliveGreen1",
		Hex:  "#d7ff87",
		HSL: hsl{
			H: 80,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 215,
			G: 255,
			B: 135,
		},
	},
	color256{
		Id:   193,
		Name: "DarkSeaGreen1",
		Hex:  "#d7ffaf",
		HSL: hsl{
			H: 90,
			S: 100,
			L: 84,
		},
		RGB: rgb{
			R: 215,
			G: 255,
			B: 175,
		},
	},
	color256{
		Id:   194,
		Name: "Honeydew2",
		Hex:  "#d7ffd7",
		HSL: hsl{
			H: 120,
			S: 100,
			L: 92,
		},
		RGB: rgb{
			R: 215,
			G: 255,
			B: 215,
		},
	},
	color256{
		Id:   195,
		Name: "LightCyan1",
		Hex:  "#d7ffff",
		HSL: hsl{
			H: 180,
			S: 100,
			L: 92,
		},
		RGB: rgb{
			R: 215,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   196,
		Name: "Red1",
		Hex:  "#ff0000",
		HSL: hsl{
			H: 0,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 255,
			G: 0,
			B: 0,
		},
	},
	color256{
		Id:   197,
		Name: "DeepPink2",
		Hex:  "#ff005f",
		HSL: hsl{
			H: 337,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 255,
			G: 0,
			B: 95,
		},
	},
	color256{
		Id:   198,
		Name: "DeepPink1",
		Hex:  "#ff0087",
		HSL: hsl{
			H: 328,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 255,
			G: 0,
			B: 135,
		},
	},
	color256{
		Id:   199,
		Name: "DeepPink1",
		Hex:  "#ff00af",
		HSL: hsl{
			H: 318,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 255,
			G: 0,
			B: 175,
		},
	},
	color256{
		Id:   200,
		Name: "Magenta2",
		Hex:  "#ff00d7",
		HSL: hsl{
			H: 309,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 255,
			G: 0,
			B: 215,
		},
	},
	color256{
		Id:   201,
		Name: "Magenta1",
		Hex:  "#ff00ff",
		HSL: hsl{
			H: 300,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 255,
			G: 0,
			B: 255,
		},
	},
	color256{
		Id:   202,
		Name: "OrangeRed1",
		Hex:  "#ff5f00",
		HSL: hsl{
			H: 22,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 255,
			G: 95,
			B: 0,
		},
	},
	color256{
		Id:   203,
		Name: "IndianRed1",
		Hex:  "#ff5f5f",
		HSL: hsl{
			H: 0,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 255,
			G: 95,
			B: 95,
		},
	},
	color256{
		Id:   204,
		Name: "IndianRed1",
		Hex:  "#ff5f87",
		HSL: hsl{
			H: 345,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 255,
			G: 95,
			B: 135,
		},
	},
	color256{
		Id:   205,
		Name: "HotPink",
		Hex:  "#ff5faf",
		HSL: hsl{
			H: 330,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 255,
			G: 95,
			B: 175,
		},
	},
	color256{
		Id:   206,
		Name: "HotPink",
		Hex:  "#ff5fd7",
		HSL: hsl{
			H: 315,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 255,
			G: 95,
			B: 215,
		},
	},
	color256{
		Id:   207,
		Name: "MediumOrchId1",
		Hex:  "#ff5fff",
		HSL: hsl{
			H: 300,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 255,
			G: 95,
			B: 255,
		},
	},
	color256{
		Id:   208,
		Name: "DarkOrange",
		Hex:  "#ff8700",
		HSL: hsl{
			H: 31,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 255,
			G: 135,
			B: 0,
		},
	},
	color256{
		Id:   209,
		Name: "Salmon1",
		Hex:  "#ff875f",
		HSL: hsl{
			H: 15,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 255,
			G: 135,
			B: 95,
		},
	},
	color256{
		Id:   210,
		Name: "LightCoral",
		Hex:  "#ff8787",
		HSL: hsl{
			H: 0,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 255,
			G: 135,
			B: 135,
		},
	},
	color256{
		Id:   211,
		Name: "PaleVioletRed1",
		Hex:  "#ff87af",
		HSL: hsl{
			H: 340,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 255,
			G: 135,
			B: 175,
		},
	},
	color256{
		Id:   212,
		Name: "OrchId2",
		Hex:  "#ff87d7",
		HSL: hsl{
			H: 320,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 255,
			G: 135,
			B: 215,
		},
	},
	color256{
		Id:   213,
		Name: "OrchId1",
		Hex:  "#ff87ff",
		HSL: hsl{
			H: 300,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 255,
			G: 135,
			B: 255,
		},
	},
	color256{
		Id:   214,
		Name: "Orange1",
		Hex:  "#ffaf00",
		HSL: hsl{
			H: 41,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 255,
			G: 175,
			B: 0,
		},
	},
	color256{
		Id:   215,
		Name: "SandyBrown",
		Hex:  "#ffaf5f",
		HSL: hsl{
			H: 30,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 255,
			G: 175,
			B: 95,
		},
	},
	color256{
		Id:   216,
		Name: "LightSalmon1",
		Hex:  "#ffaf87",
		HSL: hsl{
			H: 20,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 255,
			G: 175,
			B: 135,
		},
	},
	color256{
		Id:   217,
		Name: "LightPink1",
		Hex:  "#ffafaf",
		HSL: hsl{
			H: 0,
			S: 100,
			L: 84,
		},
		RGB: rgb{
			R: 255,
			G: 175,
			B: 175,
		},
	},
	color256{
		Id:   218,
		Name: "Pink1",
		Hex:  "#ffafd7",
		HSL: hsl{
			H: 330,
			S: 100,
			L: 84,
		},
		RGB: rgb{
			R: 255,
			G: 175,
			B: 215,
		},
	},
	color256{
		Id:   219,
		Name: "Plum1",
		Hex:  "#ffafff",
		HSL: hsl{
			H: 300,
			S: 100,
			L: 84,
		},
		RGB: rgb{
			R: 255,
			G: 175,
			B: 255,
		},
	},
	color256{
		Id:   220,
		Name: "Gold1",
		Hex:  "#ffd700",
		HSL: hsl{
			H: 50,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 255,
			G: 215,
			B: 0,
		},
	},
	color256{
		Id:   221,
		Name: "LightGoldenrod2",
		Hex:  "#ffd75f",
		HSL: hsl{
			H: 45,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 255,
			G: 215,
			B: 95,
		},
	},
	color256{
		Id:   222,
		Name: "LightGoldenrod2",
		Hex:  "#ffd787",
		HSL: hsl{
			H: 40,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 255,
			G: 215,
			B: 135,
		},
	},
	color256{
		Id:   223,
		Name: "NavajoWhite1",
		Hex:  "#ffd7af",
		HSL: hsl{
			H: 30,
			S: 100,
			L: 84,
		},
		RGB: rgb{
			R: 255,
			G: 215,
			B: 175,
		},
	},
	color256{
		Id:   224,
		Name: "MistyRose1",
		Hex:  "#ffd7d7",
		HSL: hsl{
			H: 0,
			S: 100,
			L: 92,
		},
		RGB: rgb{
			R: 255,
			G: 215,
			B: 215,
		},
	},
	color256{
		Id:   225,
		Name: "Thistle1",
		Hex:  "#ffd7ff",
		HSL: hsl{
			H: 300,
			S: 100,
			L: 92,
		},
		RGB: rgb{
			R: 255,
			G: 215,
			B: 255,
		},
	},
	color256{
		Id:   226,
		Name: "Yellow1",
		Hex:  "#ffff00",
		HSL: hsl{
			H: 60,
			S: 100,
			L: 50,
		},
		RGB: rgb{
			R: 255,
			G: 255,
			B: 0,
		},
	},
	color256{
		Id:   227,
		Name: "LightGoldenrod1",
		Hex:  "#ffff5f",
		HSL: hsl{
			H: 60,
			S: 100,
			L: 68,
		},
		RGB: rgb{
			R: 255,
			G: 255,
			B: 95,
		},
	},
	color256{
		Id:   228,
		Name: "Khaki1",
		Hex:  "#ffff87",
		HSL: hsl{
			H: 60,
			S: 100,
			L: 76,
		},
		RGB: rgb{
			R: 255,
			G: 255,
			B: 135,
		},
	},
	color256{
		Id:   229,
		Name: "Wheat1",
		Hex:  "#ffffaf",
		HSL: hsl{
			H: 60,
			S: 100,
			L: 84,
		},
		RGB: rgb{
			R: 255,
			G: 255,
			B: 175,
		},
	},
	color256{
		Id:   230,
		Name: "Cornsilk1",
		Hex:  "#ffffd7",
		HSL: hsl{
			H: 60,
			S: 100,
			L: 92,
		},
		RGB: rgb{
			R: 255,
			G: 255,
			B: 215,
		},
	},
	color256{
		Id:   231,
		Name: "Grey100",
		Hex:  "#ffffff",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 100,
		},
		RGB: rgb{
			R: 255,
			G: 255,
			B: 255,
		},
	},
	color256{
		Id:   232,
		Name: "Grey3",
		Hex:  "#080808",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 3,
		},
		RGB: rgb{
			R: 8,
			G: 8,
			B: 8,
		},
	},
	color256{
		Id:   233,
		Name: "Grey7",
		Hex:  "#121212",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 7,
		},
		RGB: rgb{
			R: 18,
			G: 18,
			B: 18,
		},
	},
	color256{
		Id:   234,
		Name: "Grey11",
		Hex:  "#1c1c1c",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 10,
		},
		RGB: rgb{
			R: 28,
			G: 28,
			B: 28,
		},
	},
	color256{
		Id:   235,
		Name: "Grey15",
		Hex:  "#262626",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 14,
		},
		RGB: rgb{
			R: 38,
			G: 38,
			B: 38,
		},
	},
	color256{
		Id:   236,
		Name: "Grey19",
		Hex:  "#303030",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 18,
		},
		RGB: rgb{
			R: 48,
			G: 48,
			B: 48,
		},
	},
	color256{
		Id:   237,
		Name: "Grey23",
		Hex:  "#3a3a3a",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 22,
		},
		RGB: rgb{
			R: 58,
			G: 58,
			B: 58,
		},
	},
	color256{
		Id:   238,
		Name: "Grey27",
		Hex:  "#444444",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 26,
		},
		RGB: rgb{
			R: 68,
			G: 68,
			B: 68,
		},
	},
	color256{
		Id:   239,
		Name: "Grey30",
		Hex:  "#4e4e4e",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 30,
		},
		RGB: rgb{
			R: 78,
			G: 78,
			B: 78,
		},
	},
	color256{
		Id:   240,
		Name: "Grey35",
		Hex:  "#585858",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 34,
		},
		RGB: rgb{
			R: 88,
			G: 88,
			B: 88,
		},
	},
	color256{
		Id:   241,
		Name: "Grey39",
		Hex:  "#626262",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 37,
		},
		RGB: rgb{
			R: 98,
			G: 98,
			B: 98,
		},
	},
	color256{
		Id:   242,
		Name: "Grey42",
		Hex:  "#6c6c6c",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 40,
		},
		RGB: rgb{
			R: 108,
			G: 108,
			B: 108,
		},
	},
	color256{
		Id:   243,
		Name: "Grey46",
		Hex:  "#767676",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 46,
		},
		RGB: rgb{
			R: 118,
			G: 118,
			B: 118,
		},
	},
	color256{
		Id:   244,
		Name: "Grey50",
		Hex:  "#808080",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 50,
		},
		RGB: rgb{
			R: 128,
			G: 128,
			B: 128,
		},
	},
	color256{
		Id:   245,
		Name: "Grey54",
		Hex:  "#8a8a8a",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 54,
		},
		RGB: rgb{
			R: 138,
			G: 138,
			B: 138,
		},
	},
	color256{
		Id:   246,
		Name: "Grey58",
		Hex:  "#949494",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 58,
		},
		RGB: rgb{
			R: 148,
			G: 148,
			B: 148,
		},
	},
	color256{
		Id:   247,
		Name: "Grey62",
		Hex:  "#9e9e9e",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 61,
		},
		RGB: rgb{
			R: 158,
			G: 158,
			B: 158,
		},
	},
	color256{
		Id:   248,
		Name: "Grey66",
		Hex:  "#a8a8a8",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 65,
		},
		RGB: rgb{
			R: 168,
			G: 168,
			B: 168,
		},
	},
	color256{
		Id:   249,
		Name: "Grey70",
		Hex:  "#b2b2b2",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 69,
		},
		RGB: rgb{
			R: 178,
			G: 178,
			B: 178,
		},
	},
	color256{
		Id:   250,
		Name: "Grey74",
		Hex:  "#bcbcbc",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 73,
		},
		RGB: rgb{
			R: 188,
			G: 188,
			B: 188,
		},
	},
	color256{
		Id:   251,
		Name: "Grey78",
		Hex:  "#c6c6c6",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 77,
		},
		RGB: rgb{
			R: 198,
			G: 198,
			B: 198,
		},
	},
	color256{
		Id:   252,
		Name: "Grey82",
		Hex:  "#d0d0d0",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 81,
		},
		RGB: rgb{
			R: 208,
			G: 208,
			B: 208,
		},
	},
	color256{
		Id:   253,
		Name: "Grey85",
		Hex:  "#dadada",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 85,
		},
		RGB: rgb{
			R: 218,
			G: 218,
			B: 218,
		},
	},
	color256{
		Id:   254,
		Name: "Grey89",
		Hex:  "#e4e4e4",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 89,
		},
		RGB: rgb{
			R: 228,
			G: 228,
			B: 228,
		},
	},
	color256{
		Id:   255,
		Name: "Grey93",
		Hex:  "#eeeeee",
		HSL: hsl{
			H: 0,
			S: 0,
			L: 93,
		},
		RGB: rgb{
			R: 238,
			G: 238,
			B: 238,
		},
	},
}

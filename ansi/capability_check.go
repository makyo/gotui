package ansi

import (
	"fmt"
	"strings"
)

// capabilityCheck builds an object testing out every ANSI capability that
// gotui knows about.
func capabilityCheck() check {
	test := map[string][]string{
		"8 Color":          make([]string, 8),
		"8 Color - Bright": make([]string, 8),
		"256 Color":        make([]string, 256),
		"24-bit Color":     make([]string, 1),
	}

	for name, num := range Colors8 {
		c, _ := Colors8.Find(name)
		if num < 60 {
			test["8 Color"][num] = fmt.Sprintf("%-30s\t%-30s", c.FG(name), c.BG(name))
		} else {
			test["8 Color - Bright"][num-60] = fmt.Sprintf("%-30s\t%-30s", c.FG(name), c.BG(name))
		}
	}

	for i, c := range Colors256 {
		test["256 Color"][i] = fmt.Sprintf("%-30s\t%-30s", c.FG(c.Name), c.BG(c.Name))
	}

	test["Attribute"] = []string{
		Bold.Apply("Bold"),
		Faint.Apply("Faint"),
		Italic.Apply("Italic"),
		Underline.Apply("Underline"),
		Blink.Apply("Blink"),
		Flash.Apply("Flash"),
		Reverse.Apply("Reverse"),
		Conceal.Apply("Conceal"),
		CrossedOut.Apply("CrossedOut"),
		AltFont1.Apply("AltFont1"),
		AltFont2.Apply("AltFont2"),
		AltFont3.Apply("AltFont3"),
		AltFont4.Apply("AltFont4"),
		AltFont5.Apply("AltFont5"),
		AltFont6.Apply("AltFont6"),
		AltFont7.Apply("AltFont7"),
		AltFont8.Apply("AltFont8"),
		AltFont9.Apply("AltFont9"),
		Fraktur.Apply("Fraktur"),
		DoubleUnderline.Apply("DoubleUnderline"),
		Framed.Apply("Framed"),
		Encircled.Apply("Encircled"),
		Overlined.Apply("Overlined"),
		IdeogramUnderline.Apply("IdeogramUnderline"),
		IdeogramDoubleUnderline.Apply("IdeogramDoubleUnderline"),
		IdeogramOverline.Apply("IdeogramOverline"),
		IdeogramDoubleOverline.Apply("IdeogramDoubleOverline"),
		IdeogramStressMarking.Apply("IdeogramStressMarking"),
	}

	hex, _ := Colors24bit.Find("#422670")
	hsl, _ := Colors24bit.Find("hsl(166, 47%, 75%)")
	rgb, _ := Colors24bit.Find("rgb(69, 35, 116)")
	test["24-bit Color"][0] = fmt.Sprintf("%s%s", hex.BG(hsl.FG("This BG is two ")), rgb.BG(hsl.FG(" different colors")))

	return test
}

type check map[string][]string

func (c check) String() string {
	var result strings.Builder
	fmt.Fprintln(&result, "8 Color")
	fmt.Fprintln(&result, "=======")
	fmt.Fprint(&result, "\n")
	for _, col := range c["8 Color"] {
		fmt.Fprintln(&result, col)
	}
	fmt.Fprint(&result, "\n\n")

	fmt.Fprintln(&result, "8 Color - Bright")
	fmt.Fprintln(&result, "================")
	fmt.Fprint(&result, "\n")
	for _, col := range c["8 Color - Bright"] {
		fmt.Fprintln(&result, col)
	}
	fmt.Fprint(&result, "\n\n")

	fmt.Fprintln(&result, "256 Color")
	fmt.Fprintln(&result, "=========")
	fmt.Fprint(&result, "\n")
	for _, col := range c["256 Color"] {
		fmt.Fprintln(&result, col)
	}
	fmt.Fprint(&result, "\n\n")

	fmt.Fprintln(&result, "Attribute")
	fmt.Fprintln(&result, "=========")
	fmt.Fprint(&result, "\n")
	for _, col := range c["Attribute"] {
		fmt.Fprintln(&result, col)
	}
	fmt.Fprint(&result, "\n\n")

	fmt.Fprintln(&result, "24-bit Color")
	fmt.Fprintln(&result, "============")
	fmt.Fprint(&result, "\n")
	fmt.Fprintln(&result, c["24-bit Color"][0])

	return result.String()
}

// CapabilityCheck is a map of all ANSI codes that gotui knows about. It is
// intended to be used as a reference to see what capabilities one's terminal
// supports.
var CapabilityCheck = capabilityCheck()

// Copyright 2014 The gotui Authors. All rights reserved.
// Use of this source code is governed by an MIT license.
// The license can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/makyo/gotui"
)

func nextView(g *gotui.Gui, v *gotui.View) error {
	if v == nil || v.Name() == "side" {
		_, err := g.SetCurrentView("main")
		return err
	}
	_, err := g.SetCurrentView("side")
	return err
}

func cursorDown(g *gotui.Gui, v *gotui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gotui.Gui, v *gotui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func getLine(g *gotui.Gui, v *gotui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, l)
		if _, err := g.SetCurrentView("msg"); err != nil {
			return err
		}
	}
	return nil
}

func delMsg(g *gotui.Gui, v *gotui.View) error {
	if err := g.DeleteView("msg"); err != nil {
		return err
	}
	if _, err := g.SetCurrentView("side"); err != nil {
		return err
	}
	return nil
}

func quit(g *gotui.Gui, v *gotui.View) error {
	return gotui.ErrQuit
}

func keybindings(g *gotui.Gui) error {
	if err := g.SetKeybinding("side", gotui.KeyCtrlSpace, gotui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gotui.KeyCtrlSpace, gotui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("side", gotui.KeyArrowDown, gotui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("side", gotui.KeyArrowUp, gotui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gotui.KeyCtrlC, gotui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("side", gotui.KeyEnter, gotui.ModNone, getLine); err != nil {
		return err
	}
	if err := g.SetKeybinding("msg", gotui.KeyEnter, gotui.ModNone, delMsg); err != nil {
		return err
	}

	if err := g.SetKeybinding("main", gotui.KeyCtrlS, gotui.ModNone, saveMain); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gotui.KeyCtrlW, gotui.ModNone, saveVisualMain); err != nil {
		return err
	}
	return nil
}

func saveMain(g *gotui.Gui, v *gotui.View) error {
	f, err := ioutil.TempFile("", "gotui_demo_")
	if err != nil {
		return err
	}
	defer f.Close()

	p := make([]byte, 5)
	v.Rewind()
	for {
		n, err := v.Read(p)
		if n > 0 {
			if _, err := f.Write(p[:n]); err != nil {
				return err
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func saveVisualMain(g *gotui.Gui, v *gotui.View) error {
	f, err := ioutil.TempFile("", "gotui_demo_")
	if err != nil {
		return err
	}
	defer f.Close()

	vb := v.ViewBuffer()
	if _, err := io.Copy(f, strings.NewReader(vb)); err != nil {
		return err
	}
	return nil
}

func layout(g *gotui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("side", -1, -1, 30, maxY); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gotui.ColorGreen
		v.SelFgColor = gotui.ColorBlack
		fmt.Fprintln(v, "Item 1")
		fmt.Fprintln(v, "Item 2")
		fmt.Fprintln(v, "Item 3")
		fmt.Fprint(v, "\rWill be")
		fmt.Fprint(v, "deleted\rItem 4\nItem 5")
	}
	if v, err := g.SetView("main", 30, -1, maxX, maxY); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		b, err := ioutil.ReadFile("Mark.Twain-Tom.Sawyer.txt")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(v, "%s", b)
		v.Editable = true
		v.Wrap = true
		if _, err := g.SetCurrentView("main"); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	g, err := gotui.NewGui(gotui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gotui.ErrQuit {
		log.Panicln(err)
	}
}

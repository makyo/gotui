// Copyright 2014 The gotui Authors. All rights reserved.
// Use of this source code is governed by an MIT license.
// The license can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	"github.com/makyo/gotui"
)

var (
	viewArr = []string{"v1", "v2", "v3", "v4"}
	active  = 0
)

func setCurrentViewOnTop(g *gotui.Gui, name string) (*gotui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func nextView(g *gotui.Gui, v *gotui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	out, err := g.View("v2")
	if err != nil {
		return err
	}
	fmt.Fprintln(out, "Going from view "+v.Name()+" to "+name)

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	if nextIndex == 0 || nextIndex == 3 {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	active = nextIndex
	return nil
}

func layout(g *gotui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("v1", 0, 0, maxX/2-1, maxY/2-1); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		v.Title = "v1 (editable)"
		v.Editable = true
		v.Wrap = true

		if _, err = setCurrentViewOnTop(g, "v1"); err != nil {
			return err
		}
	}

	if v, err := g.SetView("v2", maxX/2-1, 0, maxX-1, maxY/2-1); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		v.Title = "v2"
		v.Wrap = true
		v.Autoscroll = true
	}
	if v, err := g.SetView("v3", 0, maxY/2-1, maxX/2-1, maxY-1); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		v.Title = "v3"
		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprint(v, "Press TAB to change current view")
	}
	if v, err := g.SetView("v4", maxX/2, maxY/2, maxX-1, maxY-1); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		v.Title = "v4 (editable)"
		v.Editable = true
	}
	return nil
}

func quit(g *gotui.Gui, v *gotui.View) error {
	return gotui.ErrQuit
}

func main() {
	g, err := gotui.NewGui(gotui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gotui.ColorGreen

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gotui.KeyCtrlC, gotui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gotui.KeyTab, gotui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gotui.ErrQuit {
		log.Panicln(err)
	}
}

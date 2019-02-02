// Copyright 2014 The gotui Authors. All rights reserved.
// Use of this source code is governed by an MIT license.
// The license can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	"github.com/makyo/gotui"
)

func main() {
	g, err := gotui.NewGui(gotui.OutputNormal)
	if err != nil {
		log.Fatalln(err)
	}
	defer g.Close()

	g.Cursor = true

	g.SetManagerFunc(layout)

	if err := initKeybindings(g); err != nil {
		log.Fatalln(err)
	}

	if err := g.MainLoop(); err != nil && err != gotui.ErrQuit {
		log.Fatalln(err)
	}
}

func layout(g *gotui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("help", maxX-23, 0, maxX-1, 3); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		v.Title = "Keybindings"
		fmt.Fprintln(v, "^a: Set mask")
		fmt.Fprintln(v, "^c: Exit")
	}

	if v, err := g.SetView("input", 0, 0, maxX-24, maxY-1); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		if _, err := g.SetCurrentView("input"); err != nil {
			return err
		}
		v.Editable = true
		v.Wrap = true
	}

	return nil
}

func initKeybindings(g *gotui.Gui) error {
	if err := g.SetKeybinding("", gotui.KeyCtrlC, gotui.ModNone,
		func(g *gotui.Gui, v *gotui.View) error {
			return gotui.ErrQuit
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("input", gotui.KeyCtrlA, gotui.ModNone,
		func(g *gotui.Gui, v *gotui.View) error {
			v.Mask ^= '*'
			return nil
		}); err != nil {
		return err
	}
	return nil
}

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
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gotui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gotui.Gui) error {
	if v, err := g.SetView("v1", 10, 2, 30, 6); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "View #1")
	}
	if v, err := g.SetView("v2", 20, 4, 40, 8); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "View #2")
	}
	if v, err := g.SetView("v3", 30, 6, 50, 10); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "View #3")
	}

	return nil
}

func keybindings(g *gotui.Gui) error {
	err := g.SetKeybinding("", gotui.KeyCtrlC, gotui.ModNone, func(g *gotui.Gui, v *gotui.View) error {
		return gotui.ErrQuit
	})
	if err != nil {
		return err
	}

	err = g.SetKeybinding("", '1', gotui.ModNone, func(g *gotui.Gui, v *gotui.View) error {
		_, err := g.SetViewOnTop("v1")
		return err
	})
	if err != nil {
		return err
	}

	err = g.SetKeybinding("", '2', gotui.ModNone, func(g *gotui.Gui, v *gotui.View) error {
		_, err := g.SetViewOnTop("v2")
		return err
	})
	if err != nil {
		return err
	}

	err = g.SetKeybinding("", '3', gotui.ModNone, func(g *gotui.Gui, v *gotui.View) error {
		_, err := g.SetViewOnTop("v3")
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

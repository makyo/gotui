// Copyright 2014 The gotui Authors. All rights reserved.
// Use of this source code is governed by an MIT license.
// The license can be found in the LICENSE file.

package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

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
	maxX, _ := g.Size()

	if v, err := g.SetView("help", maxX-23, 0, maxX-1, 5); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "KEYBINDINGS")
		fmt.Fprintln(v, "↑ ↓: Seek input")
		fmt.Fprintln(v, "a: Enable autoscroll")
		fmt.Fprintln(v, "^C: Exit")
	}

	if v, err := g.SetView("stdin", 0, 0, 80, 35); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		if _, err := g.SetCurrentView("stdin"); err != nil {
			return err
		}
		dumper := hex.Dumper(v)
		if _, err := io.Copy(dumper, os.Stdin); err != nil {
			return err
		}
		v.Wrap = true
	}

	return nil
}

func initKeybindings(g *gotui.Gui) error {
	if err := g.SetKeybinding("", gotui.KeyCtrlC, gotui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("stdin", 'a', gotui.ModNone, autoscroll); err != nil {
		return err
	}
	if err := g.SetKeybinding("stdin", gotui.KeyArrowUp, gotui.ModNone,
		func(g *gotui.Gui, v *gotui.View) error {
			scrollView(v, -1)
			return nil
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("stdin", gotui.KeyArrowDown, gotui.ModNone,
		func(g *gotui.Gui, v *gotui.View) error {
			scrollView(v, 1)
			return nil
		}); err != nil {
		return err
	}
	return nil
}

func quit(g *gotui.Gui, v *gotui.View) error {
	return gotui.ErrQuit
}

func autoscroll(g *gotui.Gui, v *gotui.View) error {
	v.Autoscroll = true
	return nil
}

func scrollView(v *gotui.View, dy int) error {
	if v != nil {
		v.Autoscroll = false
		ox, oy := v.Origin()
		if err := v.SetOrigin(ox, oy+dy); err != nil {
			return err
		}
	}
	return nil
}

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

	g.Cursor = true
	g.Mouse = true

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gotui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gotui.Gui) error {
	if v, err := g.SetView("but1", 2, 2, 22, 7); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gotui.ColorGreen
		v.SelFgColor = gotui.ColorBlack
		fmt.Fprintln(v, "Button 1 - line 1")
		fmt.Fprintln(v, "Button 1 - line 2")
		fmt.Fprintln(v, "Button 1 - line 3")
		fmt.Fprintln(v, "Button 1 - line 4")
	}
	if v, err := g.SetView("but2", 24, 2, 44, 4); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gotui.ColorGreen
		v.SelFgColor = gotui.ColorBlack
		fmt.Fprintln(v, "Button 2 - line 1")
	}
	return nil
}

func keybindings(g *gotui.Gui) error {
	if err := g.SetKeybinding("", gotui.KeyCtrlC, gotui.ModNone, quit); err != nil {
		return err
	}
	for _, n := range []string{"but1", "but2"} {
		if err := g.SetKeybinding(n, gotui.MouseLeft, gotui.ModNone, showMsg); err != nil {
			return err
		}
	}
	if err := g.SetKeybinding("msg", gotui.MouseLeft, gotui.ModNone, delMsg); err != nil {
		return err
	}
	return nil
}

func quit(g *gotui.Gui, v *gotui.View) error {
	return gotui.ErrQuit
}

func showMsg(g *gotui.Gui, v *gotui.View) error {
	var l string
	var err error

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX/2-10, maxY/2, maxX/2+10, maxY/2+2); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, l)
	}
	return nil
}

func delMsg(g *gotui.Gui, v *gotui.View) error {
	if err := g.DeleteView("msg"); err != nil {
		return err
	}
	return nil
}

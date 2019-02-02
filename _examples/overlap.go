// Copyright 2014 The gotui Authors. All rights reserved.
// Use of this source code is governed by an MIT license.
// The license can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/makyo/gotui"
)

func layout(g *gotui.Gui) error {
	maxX, maxY := g.Size()
	if _, err := g.SetView("v1", -1, -1, 10, 10); err != nil &&
		err != gotui.ErrUnknownView {
		return err
	}
	if _, err := g.SetView("v2", maxX-10, -1, maxX, 10); err != nil &&
		err != gotui.ErrUnknownView {
		return err
	}
	if _, err := g.SetView("v3", maxX/2-5, -1, maxX/2+5, 10); err != nil &&
		err != gotui.ErrUnknownView {
		return err
	}
	if _, err := g.SetView("v4", -1, maxY/2-5, 10, maxY/2+5); err != nil &&
		err != gotui.ErrUnknownView {
		return err
	}
	if _, err := g.SetView("v5", maxX-10, maxY/2-5, maxX, maxY/2+5); err != nil &&
		err != gotui.ErrUnknownView {
		return err
	}
	if _, err := g.SetView("v6", -1, maxY-10, 10, maxY); err != nil &&
		err != gotui.ErrUnknownView {
		return err
	}
	if _, err := g.SetView("v7", maxX-10, maxY-10, maxX, maxY); err != nil &&
		err != gotui.ErrUnknownView {
		return err
	}
	if _, err := g.SetView("v8", maxX/2-5, maxY-10, maxX/2+5, maxY); err != nil &&
		err != gotui.ErrUnknownView {
		return err
	}
	if _, err := g.SetView("v9", maxX/2-5, maxY/2-5, maxX/2+5, maxY/2+5); err != nil &&
		err != gotui.ErrUnknownView {
		return err
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

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gotui.KeyCtrlC, gotui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gotui.ErrQuit {
		log.Panicln(err)
	}
}

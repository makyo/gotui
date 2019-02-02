// Copyright 2014 The gotui Authors. All rights reserved.
// Use of this source code is governed by an MIT license.
// The license can be found in the LICENSE file.

// WARNING: tricky code just for testing purposes, do not use as reference.

package main

import (
	"fmt"
	"log"

	"github.com/makyo/gotui"
)

var vbuf, buf string

func quit(g *gotui.Gui, v *gotui.View) error {
	vbuf = v.ViewBuffer()
	buf = v.Buffer()
	return gotui.ErrQuit
}

func overwrite(g *gotui.Gui, v *gotui.View) error {
	v.Overwrite = !v.Overwrite
	return nil
}

func layout(g *gotui.Gui) error {
	_, maxY := g.Size()
	if v, err := g.SetView("main", 0, 0, 20, maxY-1); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
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

	g.Cursor = true
	g.Mouse = true

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("main", gotui.KeyCtrlC, gotui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("main", gotui.KeyCtrlI, gotui.ModNone, overwrite); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gotui.ErrQuit {
		log.Panicln(err)
	}

	g.Close()

	fmt.Printf("VBUF:\n%s\n", vbuf)
	fmt.Printf("BUF:\n%s\n", buf)
}

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
	g.SetResizeFunc(onresize)

	if err := g.SetKeybinding("", gotui.KeyCtrlC, gotui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gotui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gotui.Gui) error {
	maxX, maxY := g.Size()
	_, err := g.SetView("size", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2)
	if err != nil && err != gotui.ErrUnknownView {
		return err
	}
	return nil
}

func onresize(g *gocui.Gui, x, y int) error {
	v, err := g.View("size")
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		return nil
	}
	v.Clear()
	fmt.Fprintf(v, "%d, %d", x, y)
	return nil
}

func quit(g *gotui.Gui, v *gotui.View) error {
	return gotui.ErrQuit
}

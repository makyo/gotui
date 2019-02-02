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

	if err := g.SetKeybinding("", gotui.KeyCtrlC, gotui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gotui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gotui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("colors", maxX/2-7, maxY/2-12, maxX/2+7, maxY/2+13); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		for i := 0; i <= 7; i++ {
			for _, j := range []int{1, 4, 7} {
				fmt.Fprintf(v, "Hello \033[3%d;%dmcolors!\033[0m\n", i, j)
			}
		}
	}
	return nil
}

func quit(g *gotui.Gui, v *gotui.View) error {
	return gotui.ErrQuit
}

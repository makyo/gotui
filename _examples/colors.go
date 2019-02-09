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
	if v, err := g.SetView("colors", maxX/2-9, maxY/2-4, maxX/2+9, maxY/2+5); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		for i := 0; i <= 7; i++ {
			for _, j := range []int{8, 1, 4, 7} {
				fmt.Fprintf(v, " \033[3%d;%dm%d:%d\033[0m", i, j, i, j)
			}
			fmt.Fprintf(v, "\n")
		}
	}
	return nil
}

func quit(g *gotui.Gui, v *gotui.View) error {
	return gotui.ErrQuit
}

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
	g, err := gotui.NewGui(gotui.Output256)

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
	if v, err := g.SetView("colors", -1, -1, maxX, maxY); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}

		// 256-colors escape codes
		for i := 0; i < 256; i++ {
			str := fmt.Sprintf("\x1b[48;5;%dm\x1b[30m%3d\x1b[0m ", i, i)
			str += fmt.Sprintf("\x1b[38;5;%dm%3d\x1b[0m ", i, i)

			if (i+1)%10 == 0 {
				str += "\n"
			}

			fmt.Fprint(v, str)
		}

		fmt.Fprint(v, "\n\n")

		// 8-colors escape codes
		ctr := 0
		for i := 0; i <= 7; i++ {
			for _, j := range []int{1, 4, 7} {
				str := fmt.Sprintf("\x1b[3%d;%dm%d:%d\x1b[0m ", i, j, i, j)
				if (ctr+1)%20 == 0 {
					str += "\n"
				}

				fmt.Fprint(v, str)

				ctr++
			}
		}
	}
	return nil
}

func quit(g *gotui.Gui, v *gotui.View) error {
	return gotui.ErrQuit
}

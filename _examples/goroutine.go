// Copyright 2014 The gotui Authors. All rights reserved.
// Use of this source code is governed by an MIT license.
// The license can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/makyo/gotui"
)

const NumGoroutines = 10

var (
	done = make(chan struct{})
	wg   sync.WaitGroup

	mu  sync.Mutex // protects ctr
	ctr = 0
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

	for i := 0; i < NumGoroutines; i++ {
		wg.Add(1)
		go counter(g)
	}

	if err := g.MainLoop(); err != nil && err != gotui.ErrQuit {
		log.Panicln(err)
	}

	wg.Wait()
}

func layout(g *gotui.Gui) error {
	if v, err := g.SetView("ctr", 2, 2, 12, 4); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "0")
	}
	return nil
}

func keybindings(g *gotui.Gui) error {
	if err := g.SetKeybinding("", gotui.KeyCtrlC, gotui.ModNone, quit); err != nil {
		return err
	}
	return nil
}

func quit(g *gotui.Gui, v *gotui.View) error {
	close(done)
	return gotui.ErrQuit
}

func counter(g *gotui.Gui) {
	defer wg.Done()

	for {
		select {
		case <-done:
			return
		case <-time.After(500 * time.Millisecond):
			mu.Lock()
			n := ctr
			ctr++
			mu.Unlock()

			g.Update(func(g *gotui.Gui) error {
				v, err := g.View("ctr")
				if err != nil {
					return err
				}
				v.Clear()
				fmt.Fprintln(v, n)
				return nil
			})
		}
	}
}

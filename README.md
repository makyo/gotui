# GOTUI - Go Terminal User Interface

[![GoDoc](https://godoc.org/github.com/makyo/gotui?status.svg)](https://godoc.org/github.com/makyo/gotui)

Lightweight Go package for creating terminal user interfaces based on [gocui](https://github.com/jroimartin/gocui)

## Features

* Minimalist API.
* Views (the "windows" in the GUI) implement the interface io.ReadWriter.
* Support for overlapping views.
* The GUI can be modified at runtime (concurrent-safe).
* Global and view-level keybindings.
* Mouse support.
* Colored text.
* Customizable edition mode.
* Easy to build reusable widgets, complex layouts...

## Installation

Execute:

```
$ go get github.com/makyo/gotui
```

## Documentation

Execute:

```
$ go doc github.com/makyo/gotui
```

Or visit [godoc.org](https://godoc.org/github.com/makyo/gotui) to read it
online.

## Example

```go
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
	if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
		if err != gotui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
	}
	return nil
}

func quit(g *gotui.Gui, v *gotui.View) error {
	return gotui.ErrQuit
}
```

# More information

For more information, please see the lovely [gocui](https://github.com/jroimartin/gocui) package this was based on.

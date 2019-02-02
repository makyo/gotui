// Copyright 2014 The gotui Authors. All rights reserved.
// Use of this source code is governed by an MIT license.
// The license can be found in the LICENSE file.

/*
Package gotui allows to create console user interfaces.

Create a new GUI:

	g, err := gotui.NewGui(gotui.OutputNormal)
	if err != nil {
		// handle error
	}
	defer g.Close()

	// Set GUI managers and key bindings
	// ...

	if err := g.MainLoop(); err != nil && err != gotui.ErrQuit {
		// handle error
	}

Set GUI managers:

	g.SetManager(mgr1, mgr2)

Managers are in charge of GUI's layout and can be used to build widgets. On
each iteration of the GUI's main loop, the Layout function of each configured
manager is executed. Managers are used to set-up and update the application's
main views, being possible to freely change them during execution. Also, it is
important to mention that a main loop iteration is executed on each reported
event (key-press, mouse event, window resize, etc).

GUIs are composed by Views, you can think of it as buffers. Views implement the
io.ReadWriter interface, so you can just write to them if you want to modify
their content. The same is valid for reading.

Create and initialize a view with absolute coordinates:

	if v, err := g.SetView("viewname", 2, 2, 22, 7); err != nil {
		if err != gotui.ErrUnknownView {
			// handle error
		}
		fmt.Fprintln(v, "This is a new view")
		// ...
	}

Views can also be created using relative coordinates:

	maxX, maxY := g.Size()
	if v, err := g.SetView("viewname", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		// ...
	}

Configure keybindings:

	if err := g.SetKeybinding("viewname", gotui.KeyEnter, gotui.ModNone, fcn); err != nil {
		// handle error
	}

gotui implements full mouse support that can be enabled with:

	g.Mouse = true

Mouse events are handled like any other keybinding:

	if err := g.SetKeybinding("viewname", gotui.MouseLeft, gotui.ModNone, fcn); err != nil {
		// handle error
	}

IMPORTANT: Views can only be created, destroyed or updated in three ways: from
the Layout function within managers, from keybinding callbacks or via
*Gui.Update(). The reason for this is that it allows gotui to be
concurrent-safe. So, if you want to update your GUI from a goroutine, you must
use *Gui.Update(). For example:

	g.Update(func(g *gotui.Gui) error {
		v, err := g.View("viewname")
		if err != nil {
			// handle error
		}
		v.Clear()
		fmt.Fprintln(v, "Writing from different goroutines")
		return nil
	})

By default, gotui provides a basic edition mode. This mode can be extended
and customized creating a new Editor and assigning it to *View.Editor:

	type Editor interface {
		Edit(v *View, key Key, ch rune, mod Modifier)
	}

DefaultEditor can be taken as example to create your own custom Editor:

	var DefaultEditor Editor = EditorFunc(simpleEditor)

	func simpleEditor(v *View, key Key, ch rune, mod Modifier) {
		switch {
		case ch != 0 && mod == 0:
			v.EditWrite(ch)
		case key == KeySpace:
			v.EditWrite(' ')
		case key == KeyBackspace || key == KeyBackspace2:
			v.EditDelete(true)
		// ...
		}
	}

Colored text:

Views allow to add colored text using ANSI colors. For example:

	fmt.Fprintln(v, "\x1b[0;31mHello world")

For more information, see the examples in folder "_examples/".
*/
package gotui

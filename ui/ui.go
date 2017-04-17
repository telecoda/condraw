package ui

import (
	"fmt"

	cui "github.com/jroimartin/gocui"
)

const (
	// Drawing
	Drawing = "Drawing"
)

var (
	curView = -1
	idxView = 0
)

func Init() (*cui.Gui, error) {
	g, err := cui.NewGui(cui.OutputNormal)
	if err != nil {
		return nil, err
	}
	g.Mouse = true
	g.Highlight = true
	g.SelFgColor = cui.ColorRed
	g.SelBgColor = cui.ColorWhite
	g.FgColor = cui.ColorBlack
	g.BgColor = cui.ColorWhite

	g.SetManagerFunc(layout)

	// init menu handlers
	if err := initMenuHandlers(g); err != nil {
		return nil, err
	}

	// init key bindings
	if err := initKeyBindings(g); err != nil {
		return nil, err
	}

	// init drawing
	if err := layoutDrawing(g); err != nil {
		return nil, err
	}
	return g, nil

}

func layout(g *cui.Gui) error {

	// init top menubar
	menuX := 0
	for _, menu := range menus {
		// layout menubar items
		if err := layoutMenuBar(g, menu, menuX); err != nil {
			return err
		}
		menuX += len(menu) + 2
	}
	return nil
}

func layoutMenuBar(g *cui.Gui, name string, xCoord int) error {

	if v, err := g.SetView(name, xCoord, -1, xCoord+len(name)+2, 1); err != nil {
		if err != cui.ErrUnknownView {
			return err
		}
		// init view content
		fmt.Fprintln(v, name)
		v.FgColor = cui.ColorBlack
		v.BgColor = cui.ColorWhite
		v.Frame = false
		g.SetViewOnTop(name)
	}
	return nil
}

func layoutDrawing(g *cui.Gui) error {

	width, height := g.Size()

	if v, err := g.SetView(Drawing, 0, 1, width, height-1); err != nil {
		if err != cui.ErrUnknownView {
			return err
		}
		// init view content
		v.FgColor = cui.ColorWhite
		v.BgColor = cui.ColorBlue
		v.Frame = true
		v.Clear()
	}

	return nil
}

func initKeyBindings(g *cui.Gui) error {
	// global quit key
	if err := g.SetKeybinding("", cui.KeyCtrlC, cui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", cui.KeyTab, cui.ModNone,
		func(g *cui.Gui, v *cui.View) error {
			return nextView(g, true)
		}); err != nil {
		return err
	}
	return nil
}

func nextView(g *cui.Gui, disableCurrent bool) error {
	next := curView + 1
	vs := g.Views()
	views := make([]string, len(vs))
	for i, view := range vs {
		views[i] = view.Name()
	}
	if next > len(views)-1 {
		next = 0
	}

	if _, err := g.SetCurrentView(views[next]); err != nil {
		return err
	}

	curView = next
	return nil
}

func initMenuHandlers(g *cui.Gui) error {

	// file menu
	if err := initMenu(g, FileMenu, fileMenu, fileMenuItems, 0, 0); err != nil {
		return err
	}
	// edit menu
	if err := initMenu(g, EditMenu, editMenu, editMenuItems, 6, 0); err != nil {
		return err
	}
	// about menu
	if err := initMenu(g, AboutMenu, aboutMenu, aboutMenuItems, 12, 0); err != nil {
		return err
	}

	return nil
}

func quit(g *cui.Gui, v *cui.View) error {
	return cui.ErrQuit
}

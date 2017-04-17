package ui

import (
	"fmt"

	"strings"

	cui "github.com/jroimartin/gocui"
)

// this package contains all the event handlers for key and mouse events

type Handler func(*cui.Gui, *cui.View) error

type MenuItem struct {
	Handler
}

const (
	// menu bar
	FileMenu  = "File"
	EditMenu  = "Edit"
	AboutMenu = "About"
)

var menus = []string{FileMenu, EditMenu, AboutMenu}

func initMenu(g *cui.Gui, menuName string, itemNames []string, itemMap map[string]MenuItem, x, y int) error {
	// init menu handler func
	menuBarClickHandler := func(g *cui.Gui, v *cui.View) error {

		if isModalDisplayed() {
			return nil
		}
		// size menu based upon longest description
		maxWidth := 0
		for _, item := range itemNames {
			width := len(item) + 2
			if width > maxWidth {
				maxWidth = width
			}
		}
		// draw menu items relative to original menu
		menuItems := v.Name() + "MenuItems"
		if v, err := g.SetView(menuItems, x+1, y+1, x+maxWidth, y+len(itemNames)+2); err != nil {
			if err != cui.ErrUnknownView {
				return err
			}
			for _, item := range itemNames {
				fmt.Fprintln(v, item)
			}
		}

		g.SetViewOnTop(v.Name())
		g.SetViewOnTop(menuItems)

		views := g.Views()
		// hide all other menuitem views
		for _, view := range views {
			if strings.HasSuffix(view.Name(), "MenuItems") && view.Name() != menuItems {
				// hide other menus
				g.DeleteView(view.Name())
			}
		}
		return nil
	}

	// init menu bar click handler
	if err := g.SetKeybinding(menuName, cui.MouseLeft, cui.ModNone, menuBarClickHandler); err != nil {
		return err
	}

	// init menu item handler
	menuItemClickHandler := func(g *cui.Gui, v *cui.View) error {

		if _, err := g.SetCurrentView(v.Name()); err != nil {
			return err
		}

		if _, err := g.SetViewOnTop(v.Name()); err != nil {
			return err
		}

		var l string
		var err error
		_, cy := v.Cursor()
		if l, err = v.Line(cy); err != nil {
			l = ""
		}

		// l should contain name of menu item
		// use this as a key to find handler
		if menuItem, ok := itemMap[l]; ok {
			return menuItem.Handler(g, v)
		} else {
			return fmt.Errorf("No handler found for menu item: %s", l)
		}
	}

	// menu item mouse clicks
	if err := g.SetKeybinding(menuName+"MenuItems", cui.MouseLeft, cui.ModNone, menuItemClickHandler); err != nil {
		return err
	}

	return nil
}

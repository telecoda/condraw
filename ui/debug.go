package ui

import (
	cui "github.com/jroimartin/gocui"
	"github.com/y0ssar1an/q"
)

func listViews(from string, g *cui.Gui) {
	q.Q("Current func:", from)
	views := g.Views()
	viewNames := make([]string, 0, 0)
	for _, view := range views {
		viewNames = append(viewNames, view.Name())
	}
	q.Q(viewNames)
}

func listCurrentView(from string, g *cui.Gui, v *cui.View) {
	q.Q("Current func:", from)
	q.Q("Current view:", v.Name())
	listViews(from, g)
}

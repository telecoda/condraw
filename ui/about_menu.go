package ui

import (
	"fmt"

	cui "github.com/jroimartin/gocui"
)

const (
	// menu items
	AboutItem = "About"
)

var Version = "v0.1"

var aboutMenu = []string{
	AboutItem,
}

var aboutMenuItems = map[string]MenuItem{
	AboutItem: MenuItem{Handler: aboutHandler},
}

var aboutDialogText = fmt.Sprintf(
	` Condraw by @telecoda
 Version: %s

 Console based drawing package
 developed using Golang
 https://github.com/telecoda/condraw
`, Version)

func aboutHandler(g *cui.Gui, v *cui.View) error {
	// delete menuItem view
	g.DeleteView(v.Name())

	width, height := g.Size()
	hWidth := width / 2
	hHeight := height / 2
	dWidth := 26
	dHeight := 6
	if v, err := g.SetView("aboutDialog", hWidth-dWidth/2-1, hHeight-dHeight/2-1, hWidth+dWidth+1, hHeight+dHeight+1); err != nil {
		if err != cui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, aboutDialogText)
		// add an ok button
		okButton := NewButtonWidget("OK", hWidth+4, hHeight+4, " OK ", aboutOKHandler)
		okButton.Layout(g)

		modalDisplayed()
	}

	return nil
}

func aboutOKHandler(g *cui.Gui, v *cui.View) error {
	// delete menuItem view
	g.DeleteView(v.Name())
	g.DeleteView("aboutDialog")

	modalClosed()

	return nil
}

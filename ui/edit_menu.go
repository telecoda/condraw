package ui

import (
	cui "github.com/jroimartin/gocui"
)

const (
	// menu items
	EditUndoItem  = "Undo"
	EditRedoItem  = "Redo"
	EditCutItem   = "Cut"
	EditCopyItem  = "Copy"
	EditPasteItem = "Paste"
)

var editMenu = []string{
	EditUndoItem,
	EditRedoItem,
	EditCutItem,
	EditCopyItem,
	EditPasteItem,
}

var editMenuItems = map[string]MenuItem{
	EditUndoItem:  MenuItem{Handler: undoHandler},
	EditRedoItem:  MenuItem{Handler: redoHandler},
	EditCutItem:   MenuItem{Handler: cutHandler},
	EditCopyItem:  MenuItem{Handler: copyHandler},
	EditPasteItem: MenuItem{Handler: pasteHandler},
}

func undoHandler(g *cui.Gui, v *cui.View) error {
	return nil
}

func redoHandler(g *cui.Gui, v *cui.View) error {
	return nil
}

func cutHandler(g *cui.Gui, v *cui.View) error {
	return nil
}

func copyHandler(g *cui.Gui, v *cui.View) error {
	return nil
}

func pasteHandler(g *cui.Gui, v *cui.View) error {
	return nil
}

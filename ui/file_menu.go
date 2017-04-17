package ui

import (
	cui "github.com/jroimartin/gocui"
)

const (
	// menu items
	FileNewItem    = "New"
	FileLoadItem   = "Load"
	FileSaveItem   = "Save"
	FileSaveAsItem = "Save as"
	FileExitItem   = "Exit"
)

var fileMenu = []string{
	FileNewItem,
	FileLoadItem,
	FileSaveItem,
	FileSaveAsItem,
	FileExitItem,
}

var fileMenuItems = map[string]MenuItem{
	FileNewItem:    MenuItem{Handler: showNewFileDialog},
	FileLoadItem:   MenuItem{Handler: showLoadFileDialog},
	FileSaveItem:   MenuItem{Handler: showSaveFileDialog},
	FileSaveAsItem: MenuItem{Handler: showSaveAsFileDialog},
	FileExitItem:   MenuItem{Handler: quit},
}

func showNewFileDialog(g *cui.Gui, v *cui.View) error {
	return nil
}

func showLoadFileDialog(g *cui.Gui, v *cui.View) error {
	return nil
}

func showSaveFileDialog(g *cui.Gui, v *cui.View) error {
	return nil
}

func showSaveAsFileDialog(g *cui.Gui, v *cui.View) error {
	return nil
}

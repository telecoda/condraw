package ui

import (
	"bufio"
	"io/ioutil"

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
	// get drawing view
	dv, err := g.View(DrawingView)
	if err != nil {
		return err
	}
	// read view state from file
	data, err := ioutil.ReadFile("./drawing.txt")
	if err != nil {
		return err
	}
	_, err = dv.Write(data)
	return err
}

func showSaveFileDialog(g *cui.Gui, v *cui.View) error {
	// get drawing view
	dv, err := g.View(DrawingView)
	if err != nil {
		return err
	}
	// write view state to file
	buf := bufio.NewReader(dv)
	data, err := ioutil.ReadAll(buf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("./drawing.txt", data, 0644)
}

func showSaveAsFileDialog(g *cui.Gui, v *cui.View) error {
	return nil
}

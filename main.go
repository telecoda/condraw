package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/telecoda/condraw/ui"
)

func main() {
	gui, err := ui.Init()
	if err != nil {
		log.Panicln(err)
	}
	defer gui.Close()

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

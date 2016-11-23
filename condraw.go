package main

import "github.com/nsf/termbox-go"

var appState state
var drawing *Drawing
var statusBar *StatusBar
var palette *PaletteDialog

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	Init()

	eventLoop()

}

func Init() {
	// init termbox settings
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.SetOutputMode(termbox.Output256)
	termWidth, termHeight := termbox.Size()
	// drawing is 1 line less than terminal to allow for status bar
	drawing = NewDrawing(termWidth, termHeight-1)
	// init UI
	statusBar = InitStatusBar()
	palette = newPaletteDialog()

	// default brush
	brush = NewBrush(defaultBrushChar, defaultBrushFg, defaultBrushBg, 4, 2)

	// register functions
	registerEventHandlers()
	registerRenderers()
	setState(drawState)

}

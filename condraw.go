package main

import "github.com/nsf/termbox-go"

var appState state
var drawing *Drawing
var statusBar *StatusBar
var paletteDialog Dialog
var modeDialog Dialog
var brushDialog Dialog
var mode termbox.OutputMode

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
	//mode = termbox.OutputGrayscale
	mode = termbox.Output256
	//mode = termbox.Output216
	termbox.SetOutputMode(mode)
	//termbox.SetOutputMode(termbox.OutputGrayscale)
	termWidth, termHeight := termbox.Size()
	// drawing is 1 line less than terminal to allow for status bar
	drawing = NewDrawing(termWidth, termHeight-1, mode)
	// init UI
	statusBar = InitStatusBar()
	brushDialog = newBrushDialog()
	modeDialog = newModeDialog()
	paletteDialog = newPaletteDialog()

	// default brush
	brush = NewBrush(defaultBrushChar, defaultBrushFg, defaultBrushBg, 1, 1)

	// register functions
	registerEventHandlers()
	registerRenderers()
	setState(drawState)

}

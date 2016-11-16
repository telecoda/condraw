package main

import "github.com/nsf/termbox-go"

var appState state
var drawing *Drawing

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
	registerEventHandlers()
	registerRenderers()
	setState(drawState)
	termWidth, termHeight := termbox.Size()
	// drawing is 1 line less than terminal to allow for status bar
	drawing = NewDrawing(termWidth, termHeight-2)

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	reallocBackBuffer(termbox.Size())
}

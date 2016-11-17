package main

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
)

var renderers map[state]renderer
var currentRenderer renderer

func registerRenderers() {
	renderers = make(map[state]renderer)

	renderers[drawState] = drawingRenderer
	renderers[menuBarState] = menuBarRenderer
	renderers[menuState] = menuRenderer
	renderers[dialogState] = dialogRenderer
}

func drawingRenderer() {

	drawing.render()

	statusBarRenderer()

	cursorRenderer()
}

func cursorRenderer() {
	// render cursor
	termbox.SetCell(cursorX, cursorY, cursorRune, cursorFg, cursorBg)

}

func statusBarRenderer() {
	// bottom of screen
	_, height := termbox.Size()
	renderText(fmt.Sprintf("Cursor: %03d,%03d | Size: %03d, %03d | Event: %#v",
		cursorX, cursorY, drawing.width, drawing.height, lastMouseEvent), 0, height-1, statusBarFg, statusBarBg)
}

func menuBarRenderer() {
	// top of screen
	renderMenuBar()
}

func menuRenderer() {
}

func dialogRenderer() {
}

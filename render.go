package main

import termbox "github.com/nsf/termbox-go"

var renderers map[state]renderer
var currentRenderer renderer

func registerRenderers() {
	renderers = make(map[state]renderer)

	renderers[drawState] = drawingRenderer
	renderers[paletteState] = paletteRenderer
	renderers[modeState] = modeRenderer
	renderers[brushState] = brushRenderer
	//renderers[menuBarState] = menuBarRenderer
	//renderers[menuState] = menuRenderer
	//renderers[dialogState] = dialogRenderer
}

func drawingRenderer() {

	drawing.render()

	//statusBarRenderer()

	statusBar.Render()

	cursorRenderer()
}

func cursorRenderer() {
	// render cursor
	// only render when inside drawing bounds
	if drawing.inBounds(cursorX, cursorY) {
		termbox.SetCell(cursorX, cursorY, cursorRune, cursorFg, cursorBg)
	}
}

func brushRenderer() {
	drawingRenderer()

	// render brush dialog
	brushDialog.Render()
}

func paletteRenderer() {
	drawingRenderer()

	// render palette dialog
	paletteDialog.Render()
}

func modeRenderer() {
	drawingRenderer()

	// render mode dialog
	modeDialog.Render()
}

func menuBarRenderer() {
	// top of screen
	renderMenuBar()
}

func menuRenderer() {
}

func dialogRenderer() {
}

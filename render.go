package main

import termbox "github.com/nsf/termbox-go"

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

// func statusBarRenderer() {
// 	// bottom of screen
// 	_, height := termbox.Size()
// 	barY := height - 1
// 	barX := 0
// 	ui.RenderBar('-', barY, statusBarFg, statusBarBg)

// 	ui.RenderText(fmt.Sprintf("Fg: X | Bg: X | Brush: X | Cursor: %03d,%03d | Size: %03d, %03d | Event: %#v",
// 		cursorX, cursorY, drawing.width, drawing.height, lastMouseEvent), barX, barY, statusBarFg, statusBarBg)

// 	// Render brush details
// 	ui.RenderRune(' ', barX+4, barY, brushFg, brushFg)
// 	ui.RenderRune(' ', barX+12, barY, brushBg, brushBg)
// 	ui.RenderRune(brush, barX+23, barY, brushFg, brushBg)

// }

func menuBarRenderer() {
	// top of screen
	renderMenuBar()
}

func menuRenderer() {
}

func dialogRenderer() {
}

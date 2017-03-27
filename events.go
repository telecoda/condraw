package main

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
)

var lastKeyEvent termbox.Event
var lastMouseEvent termbox.Event

type eventHandler func(event termbox.Event)
type renderer func()

var handlers map[state]eventHandler
var currentHandler eventHandler

func registerEventHandlers() {
	handlers = make(map[state]eventHandler)

	handlers[drawState] = drawingHandler
	handlers[paletteState] = paletteHandler
	handlers[modeState] = modeHandler
	handlers[brushState] = brushHandler
	//handlers[menuState] = menuHandler
	//handlers[dialogState] = dialogHandler
}

// setState sets state of app and assigns event handler/renderer
func setState(newState state) {
	if handler, ok := handlers[newState]; !ok {
		panic(fmt.Sprintf("State: %d no supporting event handler", newState))
	} else {
		if renderer, ok := renderers[newState]; !ok {
			panic(fmt.Sprintf("State: %d no supporting renderer", newState))
		} else {
			currentHandler = handler
			currentRenderer = renderer
			appState = newState
		}

	}
}

func eventLoop() {

	currentRenderer()
	termbox.Flush()

mainloop:
	for {

		// global event handling first
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:

			lastKeyEvent = ev
			if ev.Key == termbox.KeyCtrlC {
				break mainloop
			} else {
				// route unhandled key events to appropriate handler
				currentHandler(ev)
			}
		case termbox.EventResize:
			drawing.resizeConsole(ev.Width, ev.Height)
		case termbox.EventMouse:
			currentHandler(ev)
		default:
			// route any other events to appropriate handler
			//currentHandler(ev)
		}

		currentRenderer()
		termbox.Flush()

	}

}

var keyw = rune('w')
var keyW = rune('W')
var keya = rune('a')
var keyA = rune('A')
var keys = rune('s')
var keyS = rune('S')
var keyd = rune('d')
var keyD = rune('D')
var keyg = rune('g')
var keyPlus = rune('=')
var keyMinus = rune('-')

func drawingHandler(ev termbox.Event) {

	switch ev.Type {
	case termbox.EventKey:

		if ev.Key == termbox.KeyTab {
			//setState(menuBarState)
		}
		if ev.Key == termbox.KeySpace {
			brush.paintToDrawing(drawing)
			return
		}
		if ev.Key == termbox.KeyDelete || ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
			brush.eraseFromDrawing(drawing)
			return
		}

		switch ev.Ch {
		case keya:
			cursorLeft()
		case keyA:
			brush.paintToDrawing(drawing)
			cursorLeft()
		case keyd:
			cursorRight()
		case keyD:
			brush.paintToDrawing(drawing)
			cursorRight()
		case keyw:
			cursorUp()
		case keyW:
			brush.paintToDrawing(drawing)
			cursorUp()
		case keys:
			cursorDown()
		case keyS:
			brush.paintToDrawing(drawing)
			cursorDown()
		case keyg:
			drawing.ToggleGrid()
			drawing.render()
		case keyPlus:
			brush.increaseSize()
		case keyMinus:
			brush.decreaseSize()
		}
	case termbox.EventMouse:
		lastMouseEvent = ev
		// check if mouse clicked on status bar
		_, height := termbox.Size()
		if ev.MouseY == height-1 {
			// click on status bar
			statusBar.Handle(ev)
		} else {
			// move cursor
			cursorX, cursorY = ev.MouseX, ev.MouseY
			// click on drawing
			if ev.Key == termbox.MouseLeft {
				brush.paintToDrawing(drawing)
			}
			if ev.Key == termbox.MouseRight {
				brush.eraseFromDrawing(drawing)
			}
		}
	}

}

func menuBarHandler(ev termbox.Event) {
	switch ev.Type {
	case termbox.EventKey:
		if ev.Key == termbox.KeyTab {
			setState(drawState)
		}
	}
}

func menuHandler(ev termbox.Event) {

}

func dialogHandler(event termbox.Event) {
}

func brushHandler(ev termbox.Event) {
	brushDialog.Handle(ev)
}

func paletteHandler(ev termbox.Event) {
	paletteDialog.Handle(ev)
}

func modeHandler(ev termbox.Event) {
	modeDialog.Handle(ev)
}

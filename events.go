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
	handlers[menuBarState] = menuBarHandler
	handlers[menuState] = menuHandler
	handlers[dialogState] = dialogHandler
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

func drawingHandler(ev termbox.Event) {

	switch ev.Type {
	case termbox.EventKey:

		if ev.Key == termbox.KeyTab {
			setState(menuBarState)
		}
		if ev.Key == termbox.KeySpace {
			drawing.paintAtCursor()
			return
		}
		if ev.Key == termbox.KeyDelete || ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
			drawing.eraseAtCursor()
			return
		}

		switch ev.Ch {
		case keya:
			cursorLeft()
		case keyA:
			drawing.paintAtCursor()
			cursorLeft()
		case keyd:
			cursorRight()
		case keyD:
			drawing.paintAtCursor()
			cursorRight()
		case keyw:
			cursorUp()
		case keyW:
			drawing.paintAtCursor()
			cursorUp()
		case keys:
			cursorDown()
		case keyS:
			drawing.paintAtCursor()
			cursorDown()
		}
	case termbox.EventMouse:
		lastMouseEvent = ev
		// move cursor
		cursorX, cursorY = ev.MouseX, ev.MouseY
		if ev.Key == termbox.MouseLeft {
			drawing.paintAtCursor()
		}
		if ev.Key == termbox.MouseRight {
			drawing.eraseAtCursor()
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

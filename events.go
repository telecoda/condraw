package main

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
)

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
			if ev.Key == termbox.KeyCtrlC {
				break mainloop
			} else {
				// route unhandled key events to appropriate handler
				currentHandler(ev)
			}
		case termbox.EventResize:
			reallocBackBuffer(ev.Width, ev.Height)
		default:
			// route any other events to appropriate handler
			//currentHandler(ev)
		}
		currentRenderer()
		termbox.Flush()

	}

}

// switch ev := termbox.PollEvent(); ev.Type {
// case termbox.EventKey:
// 	if ev.Key == termbox.KeyCtrlC {
// 		break mainloop
// 	}
// 	if ev.Key == termbox.KeyEsc {
// 		toggleMode()
// 	}
// 	if ev.Key == termbox.KeyArrowLeft {
// 		cursorX--
// 	}
// 	if ev.Key == termbox.KeyArrowRight {
// 		cursorX++
// 	}
// 	if ev.Key == termbox.KeyArrowUp {
// 		cursorY--
// 	}
// 	if ev.Key == termbox.KeyArrowDown {
// 		cursorY++
// 	}
// case termbox.EventMouse:
// 	if ev.Key == termbox.MouseLeft {
// 		mx, my = ev.MouseX, ev.MouseY
// 	}
// case termbox.EventResize:
// 	reallocBackBuffer(ev.Width, ev.Height)
// }
var keyw = rune('w')
var keyW = rune('W')
var keya = rune('a')
var keyA = rune('A')
var keys = rune('s')
var keyS = rune('S')
var keyd = rune('d')
var keyD = rune('D')

var lastEvent termbox.Event

func drawingHandler(ev termbox.Event) {

	switch ev.Type {
	case termbox.EventKey:

		lastEvent = ev

		if ev.Key == termbox.KeyEsc {
			setState(menuBarState)
		}
		if ev.Key == termbox.KeySpace {
			paintAtCursor()
			return
		}
		if ev.Key == termbox.KeyDelete || ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
			deleteAtCursor()
			return
		}

		switch ev.Ch {
		case keya:
			cursorLeft()
		case keyA:
			paintAtCursor()
			cursorLeft()
		case keyd:
			cursorRight()
		case keyD:
			paintAtCursor()
			cursorRight()
		case keyw:
			cursorUp()
		case keyW:
			paintAtCursor()
			cursorUp()
		case keys:
			cursorDown()
		case keyS:
			paintAtCursor()
			cursorDown()
		}
	}

}

func menuBarHandler(ev termbox.Event) {
	switch ev.Type {
	case termbox.EventKey:
		if ev.Key == termbox.KeyEsc {
			setState(drawState)
		}
	}
}

func menuHandler(ev termbox.Event) {

}

func dialogHandler(event termbox.Event) {

}

package main

import termbox "github.com/nsf/termbox-go"

var cursorX, cursorY int
var brush = '░'

var brushRunes = []rune{' ', '░', '▒', '▓', '█'}

// Colours
var brushFg = termbox.ColorBlack
var brushBg = termbox.ColorWhite
var menuForeground = termbox.ColorBlack
var menuBackground = termbox.ColorWhite
var statusBarFg = termbox.ColorWhite
var statusBarBg = termbox.ColorBlack
var cursorRune = ' '
var cursorFg = termbox.ColorRed
var cursorBg = termbox.ColorRed

var backbuf []termbox.Cell
var bbw, bbh int

type Drawing struct {
	width, height int
	drawBuf       []termbox.Cell
}

func NewDrawing(width, height int) *Drawing {
	drawing := &Drawing{
		width:  width,
		height: height,
	}

	cursorX = 10
	cursorY = 10
	return drawing
}

func (d *Drawing) render() {
	copy(termbox.CellBuffer(), backbuf)
}

func reallocBackBuffer(w, h int) {
	bbw, bbh = w, h
	backbuf = make([]termbox.Cell, w*h)
}

func cursorUp() {
	if cursorY <= 0 {
		return
	}
	cursorY--
}

func cursorDown() {
	if cursorY >= drawing.height {
		return
	}
	cursorY++
}

func cursorLeft() {
	if cursorX <= 0 {
		return
	}
	cursorX--
}

func cursorRight() {
	if cursorX >= drawing.width {
		return
	}
	cursorX++
}

func paintAtCursor() {
	backbuf[bbw*cursorY+cursorX] = termbox.Cell{Ch: brush, Fg: brushFg, Bg: brushBg}
}

func deleteAtCursor() {
	backbuf[bbw*cursorY+cursorX] = termbox.Cell{Ch: ' ', Fg: brushBg, Bg: brushFg}
}

package main

import termbox "github.com/nsf/termbox-go"

var cursorX, cursorY int
var eraser = ' '
var brushRunes = []rune{' ', '░', '▒', '▓', '█'}
var defaultBrushChar = '░'

var brush *Brush

// Colours
var defaultBrushFg = termbox.ColorRed
var defaultBrushBg = termbox.ColorBlue
var menuForeground = termbox.ColorBlack
var menuBackground = termbox.ColorWhite
var statusBarFg = termbox.ColorBlack
var statusBarBg = termbox.ColorWhite
var cursorRune = ' '
var cursorFg = termbox.ColorRed
var cursorBg = termbox.ColorRed

//var backbuf []termbox.Cell
//var bbw, bbh int

type Drawing struct {
	mode          termbox.OutputMode
	width, height int
	drawBuf       []termbox.Cell
}

func NewDrawing(width, height int, mode termbox.OutputMode) *Drawing {
	drawing := &Drawing{
		mode:    mode,
		width:   width,
		height:  height,
		drawBuf: make([]termbox.Cell, width*height),
	}

	cursorX = 10
	cursorY = 10

	drawing.defaultDrawing()

	return drawing
}

func (d *Drawing) defaultDrawing() {

}

func (d *Drawing) render() {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlue)
	// copy from drawing buffer to on screen buffer
	// dimension may differ..
	uiWidth, uiHeight := termbox.Size()

	for x := 0; x <= uiWidth; x++ {
		for y := 0; y <= uiHeight; y++ {
			cell := d.GetCell(x, y)
			if cell != nil {
				termbox.SetCell(x, y, cell.Ch, cell.Fg, cell.Bg)
			}
		}
	}

}

func (d *Drawing) GetCell(x, y int) *termbox.Cell {
	if d.inBounds(x, y) {
		return &d.drawBuf[d.width*y+x]
	}
	return nil
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

func (d *Drawing) inBounds(x, y int) bool {
	// check cursor is on drawing bounds
	if x >= d.width || y >= d.height {
		return false
	}

	if x < 0 || y < 0 {
		return false
	}

	return true
}

func (d *Drawing) resizeConsole(width, height int) {

}

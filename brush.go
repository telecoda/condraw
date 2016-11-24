package main

import termbox "github.com/nsf/termbox-go"

type Brush struct {
	char          rune
	fg            termbox.Attribute
	bg            termbox.Attribute
	width, height int
}

func NewBrush(char rune, fg, bg termbox.Attribute, width, height int) *Brush {
	brush := &Brush{
		char:   char,
		fg:     fg,
		bg:     bg,
		width:  width,
		height: height,
	}

	return brush
}

func (b *Brush) paintToScreen(screenX, screenY int) {
	offsetX := b.width / 2
	offsetY := b.height / 2
	for bx := 0; bx < b.width; bx++ {
		for by := 0; by < b.height; by++ {
			termbox.SetCell(screenX+bx-offsetX, screenY+by-offsetY, b.char, b.fg, b.bg)
		}
	}
}

func (b *Brush) paintToDrawing(d *Drawing) {
	offsetX := b.width / 2
	offsetY := b.height / 2
	if d.inBounds(cursorX, cursorY) {
		for x := 0; x < b.width; x++ {
			for y := 0; y < b.height; y++ {
				d.drawBuf[d.width*(cursorY+y-offsetY)+cursorX+x-offsetX] = termbox.Cell{Ch: b.char, Fg: b.fg, Bg: b.bg}
			}
		}
	}
}

func (b *Brush) eraseFromDrawing(d *Drawing) {
	offsetX := b.width / 2
	offsetY := b.height / 2
	if d.inBounds(cursorX, cursorY) {
		for x := 0; x < b.width; x++ {
			for y := 0; y < b.height; y++ {
				d.drawBuf[d.width*(cursorY+y-offsetY)+cursorX+x-offsetX] = termbox.Cell{Ch: eraser, Fg: b.bg, Bg: b.fg}
			}
		}
	}
}

package main

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/y0ssar1an/q"
)

type Brush struct {
	char          rune
	fg            termbox.Attribute
	bg            termbox.Attribute
	width, height int
	drawing       *Drawing
}

func NewBrush(char rune, fg, bg termbox.Attribute, width, height int) *Brush {
	brush := &Brush{
		char:    char,
		fg:      fg,
		bg:      bg,
		width:   width,
		height:  height,
		drawing: NewDrawing(2, 2, mode),
	}

	brush.initDrawing()

	return brush
}

// temp code to define a brush internally as a drawing
func (b *Brush) initDrawing() {

	brushDrawing := b.drawing
	for x := 0; x <= b.width; x++ {
		for y := 0; y <= b.height; y++ {
			q.Q(x, y)
			brushDrawing.drawBuf[brushDrawing.width*y+x] = termbox.Cell{Ch: rune('B'), Fg: termbox.ColorGreen, Bg: termbox.ColorYellow}
		}
	}
	q.Q(b)
}

func (b *Brush) getBrushCell(x, y int) termbox.Cell {
	// calc offset inside brush
	// brush will be repeated evenly across drawing
	brushX := x % b.drawing.width
	brushY := y % b.drawing.height

	return b.drawing.drawBuf[b.drawing.width*brushY+brushX]
}

func (b *Brush) paintToScreen(screenX, screenY int) {
	offsetX := b.width / 2
	offsetY := b.height / 2
	for bx := 0; bx < b.width; bx++ {
		for by := 0; by < b.height; by++ {
			// paint with brush drawing data
			// termbox.SetCell(screenX+bx-offsetX, screenY+by-offsetY, b.char, b.fg, b.bg)
			brushCell := b.getBrushCell(screenX, screenY)
			termbox.SetCell(screenX+bx-offsetX, screenY+by-offsetY, brushCell.Ch, brushCell.Fg, brushCell.Bg)

		}
	}
}

func (b *Brush) paintToDrawing(d *Drawing) {
	offsetX := b.width / 2
	offsetY := b.height / 2
	if d.inBounds(cursorX, cursorY) {
		for x := 0; x < b.width; x++ {
			for y := 0; y < b.height; y++ {
				// paint with brush drawing data
				//				d.drawBuf[d.width*(cursorY+y-offsetY)+cursorX+x-offsetX] = termbox.Cell{Ch: b.char, Fg: b.fg, Bg: b.bg}
				brushCell := b.getBrushCell(cursorX, cursorY)
				d.drawBuf[d.width*(cursorY+y-offsetY)+cursorX+x-offsetX] = termbox.Cell{Ch: brushCell.Ch, Fg: brushCell.Fg, Bg: brushCell.Bg}
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

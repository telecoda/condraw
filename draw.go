package main

import (
	"fmt"
	"os"

	termbox "github.com/nsf/termbox-go"
)

var cursorX, cursorY int
var eraser = ' '
var brushRunes = []rune{' ', '░', '▒', '▓', '█'}

//var defaultBrushChar = '░'
var defaultBrushChar = rune(' ')
var defaultGridChar = rune(9633)

// Rune: 8414 , ⃞
// Rune: 8419 , ⃣
// Rune: 8987 , ⌛
// Rune: 9600 , ▀
// Rune: 9601 , ▁
// Rune: 9602 , ▂
// Rune: 9603 , ▃
// Rune: 9604 , ▄
// Rune: 9605 , ▅
// Rune: 9606 , ▆
// Rune: 9607 , ▇
// Rune: 9608 , █
// Rune: 9609 , ▉
// Rune: 9610 , ▊
// Rune: 9611 , ▋
// Rune: 9612 , ▌
// Rune: 9613 , ▍
// Rune: 9614 , ▎
// Rune: 9615 , ▏
// Rune: 9616 , ▐
// Rune: 9617 , ░
// Rune: 9618 , ▒
// Rune: 9619 , ▓
// Rune: 9620 , ▔
// Rune: 9621 , ▕
// Rune: 9622 , ▖
// Rune: 9623 , ▗
// Rune: 9624 , ▘
// Rune: 9625 , ▙
// Rune: 9626 , ▚
// Rune: 9627 , ▛
// Rune: 9628 , ▜
// Rune: 9629 , ▝
// Rune: 9630 , ▞
// Rune: 9631 , ▟
// Rune: 9632 , ■
// Rune: 9633 , □
// Rune: 9634 , ▢
// Rune: 9635 , ▣
// Rune: 9636 , ▤
// Rune: 9637 , ▥
// Rune: 9638 , ▦
// Rune: 9639 , ▧
// Rune: 9640 , ▨
// Rune: 9641 , ▩
// Rune: 9698 , ◢
// Rune: 9699 , ◣
// Rune: 9700 , ◤
// Rune: 9701 , ◥
// Rune: 10061 , ❍
// Rune: 10063 , ❏
// Rune: 10064 , ❐
// Rune: 10065 , ❑
// Rune: 10066 , ❒

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
	displayGrid   bool
	drawBuf       []termbox.Cell
}

func NewDrawing(width, height int, mode termbox.OutputMode) *Drawing {
	drawing := &Drawing{
		mode:        mode,
		width:       width,
		height:      height,
		displayGrid: false,
		drawBuf:     make([]termbox.Cell, width*height),
	}

	cursorX = 10
	cursorY = 10

	drawing.defaultDrawing()

	return drawing
}

func (d *Drawing) defaultDrawing() {

}

func (d *Drawing) ToggleGrid() {
	d.displayGrid = !d.displayGrid
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
				if d.displayGrid {
					// if cell is empty draw grid char
					if cell.Bg == 0 && cell.Fg == 0 && cell.Ch == 0 {
						termbox.SetCell(x, y, defaultGridChar, termbox.ColorWhite, termbox.ColorBlack)
					} else {
						termbox.SetCell(x, y, defaultGridChar, cell.Fg, cell.Bg)
					}
				} else {
					termbox.SetCell(x, y, cell.Ch, cell.Fg, cell.Bg)
				}
			}
		}
	}

}

// save drawing to disk
func (d *Drawing) save(filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("TEMP: error opening file: %s\n", err)
		return err
	}

	uiWidth, uiHeight := termbox.Size()

	for x := 0; x <= uiWidth; x++ {
		for y := 0; y <= uiHeight; y++ {
			cell := d.GetCell(x, y)
			if cell != nil {
				file.Write([]byte(string(cell.Fg)))
				file.Write([]byte(string(cell.Bg)))
				file.Write([]byte(string(cell.Ch)))
			}
		}
		// newline
		file.Write([]byte("\n"))
	}

	file.Close()

	return nil
}

// load drawing from disk
func (d *Drawing) load(filename string) error {
	return nil
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

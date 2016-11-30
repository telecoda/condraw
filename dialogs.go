package main

import (
	"image"

	termbox "github.com/nsf/termbox-go"
	"github.com/telecoda/condraw/ui"
)

type Dialog interface {
	Show(mode termbox.OutputMode, title string, callback func(returnValue interface{}))
	Render()
	Handle(ev termbox.Event)
}

type PaletteDialog struct {
	mode termbox.OutputMode
	ui.Component
	Title    string
	callback func(returnValue interface{})
}

type Palette256Dialog struct {
	PaletteDialog
}

type Palette216Dialog struct {
	PaletteDialog
}

type PaletteGrayscaleDialog struct {
	PaletteDialog
}

func newPaletteDialog() Dialog {
	var pal Dialog
	switch mode {
	case termbox.OutputNormal:
		pal = &PaletteDialog{}
	case termbox.Output256:
		pal = &Palette256Dialog{}
	case termbox.Output216:
		pal = &Palette216Dialog{}
	case termbox.OutputGrayscale:
		pal = &PaletteGrayscaleDialog{}
	default:
		pal = &PaletteDialog{}
	}
	return pal
}

func (p *PaletteDialog) content() string {
	return "1234567812345678"
}

func (p *PaletteDialog) Show(mode termbox.OutputMode, title string, callback func(returnValue interface{})) {
	p.mode = mode
	p.Title = title
	p.callback = callback
	setState(paletteState)
}

func (p *PaletteDialog) Render() {
	// 8 colour palette
	// make sure component is centred
	termWidth, termHeight := termbox.Size()
	cx := termWidth / 2
	cy := termHeight / 2
	coloursOnLine := 8
	offsetX := coloursOnLine
	offsetY := coloursOnLine / 2
	p.X = cx - offsetX
	p.Y = cy - offsetY
	for i := termbox.Attribute(0); i < termbox.Attribute(8); i++ {
		px := int(i) % coloursOnLine
		py := int(i) / coloursOnLine
		p.renderColour(p.X+px*2, p.Y+py, i)
	}
	ui.RenderText("----pallette----", p.X, p.Y-1, statusBarFg, statusBarBg)
}

func (p *PaletteDialog) Handle(ev termbox.Event) {
	switch ev.Type {

	case termbox.EventMouse:
		// click on palette
		if ev.Key == termbox.MouseLeft {
			p.selectColour(ev.MouseX, ev.MouseY)
			setState(drawState)
		}
	}

}

func (p *PaletteDialog) selectColour(x, y int) {
	colourX := (x - p.X) / 2
	colourY := y - p.Y
	index := colourX + colourY*16
	if index >= 0 && index < 8 {
		p.callback(termbox.Attribute(index))
	}
}

func (p *Palette256Dialog) Render() {
	// 256 colour palette
	// make sure component is centred
	termWidth, termHeight := termbox.Size()
	cx := termWidth / 2
	cy := termHeight / 2
	coloursOnLine := 16
	offsetX := coloursOnLine
	offsetY := coloursOnLine / 2
	p.X = cx - offsetX
	p.Y = cy - offsetY
	for i := termbox.Attribute(0); i < termbox.Attribute(256); i++ {
		px := int(i) % coloursOnLine
		py := int(i) / coloursOnLine
		p.renderColour(p.X+px*2, p.Y+py, i)
	}
	ui.RenderText("-------------palette------------", p.X, p.Y-1, statusBarFg, statusBarBg)
}

func (p *Palette256Dialog) Handle(ev termbox.Event) {
	switch ev.Type {

	case termbox.EventMouse:
		// click on palette
		if ev.Key == termbox.MouseLeft {
			p.selectColour(ev.MouseX, ev.MouseY)
			setState(drawState)
		}
	}

}

func (p *Palette256Dialog) selectColour(x, y int) {
	colourX := (x - p.X) / 2
	colourY := y - p.Y
	index := colourX + colourY*16
	if index >= 0 && index < 256 {
		p.callback(termbox.Attribute(index))
	}
}

func (p *Palette216Dialog) Render() {
	// make sure component is centred
	// make sure component is centred
	termWidth, termHeight := termbox.Size()
	cx := termWidth / 2
	cy := termHeight / 2
	coloursOnLine := 16
	offsetX := coloursOnLine
	offsetY := coloursOnLine / 2
	p.X = cx - offsetX
	p.Y = cy - offsetY
	for i := termbox.Attribute(0); i < termbox.Attribute(256); i++ {
		px := int(i) % coloursOnLine
		py := int(i) / coloursOnLine
		p.renderColour(p.X+px*2, p.Y+py, i)
	}
	ui.RenderText("-------------palette------------", p.X, p.Y-1, statusBarFg, statusBarBg)
}

func (p *Palette216Dialog) Handle(ev termbox.Event) {
	switch ev.Type {

	case termbox.EventMouse:
		// click on palette
		if ev.Key == termbox.MouseLeft {
			p.selectColour(ev.MouseX, ev.MouseY)
			setState(drawState)
		}
	}

}

func (p *Palette216Dialog) selectColour(x, y int) {
	colourX := (x - p.X) / 2
	colourY := y - p.Y
	index := colourX + colourY*16
	if index >= 0 && index < 216 {
		p.callback(termbox.Attribute(index))
	}
}

func (p *PaletteGrayscaleDialog) Render() {
	// Grayscale colour palette
	// make sure component is centred
	termWidth, termHeight := termbox.Size()
	cx := termWidth / 2
	cy := termHeight / 2
	coloursOnLine := 12
	offsetX := coloursOnLine
	offsetY := coloursOnLine / 2
	p.X = cx - offsetX
	p.Y = cy - offsetY
	for i := termbox.Attribute(0); i < termbox.Attribute(24); i++ {
		px := int(i) % coloursOnLine
		py := int(i) / coloursOnLine
		p.renderColour(p.X+px*2, p.Y+py, i)
	}
	ui.RenderText("---------palette--------", p.X, p.Y-1, statusBarFg, statusBarBg)
}

func (p *PaletteGrayscaleDialog) Handle(ev termbox.Event) {
	switch ev.Type {

	case termbox.EventMouse:
		// click on palette
		if ev.Key == termbox.MouseLeft {
			p.selectColour(ev.MouseX, ev.MouseY)
			setState(drawState)
		}
	}
}

func (p *PaletteGrayscaleDialog) selectColour(x, y int) {
	colourX := (x - p.X) / 2
	colourY := y - p.Y
	index := colourX + colourY*12
	if index >= 0 && index < 24 {
		p.callback(termbox.Attribute(index))
	}
}

func (p *PaletteDialog) renderColour(x, y int, colour termbox.Attribute) {
	rect := image.Rect(x, y, x+2, y+1)
	ui.RenderRect(rect, colour)
}

type ModeDialog struct {
	mode termbox.OutputMode
	ui.Component
	Title    string
	callback func(resultVar interface{})
}

func newModeDialog() Dialog {
	modeDialog := &ModeDialog{}

	return modeDialog
}

func (m *ModeDialog) Render() {
	// make sure component is centred
	termWidth, termHeight := termbox.Size()
	cx := termWidth / 2
	cy := termHeight / 2
	coloursOnLine := 12
	offsetX := coloursOnLine
	offsetY := coloursOnLine / 2
	m.X = cx - offsetX
	m.Y = cy - offsetY
	ui.RenderText("-------mode--------", m.X, m.Y-1, statusBarFg, statusBarBg)
	ui.RenderText("|   8 colour mode |", m.X, m.Y, statusBarFg, statusBarBg)
	ui.RenderText("| 256 colour mode |", m.X, m.Y+1, statusBarFg, statusBarBg)
	ui.RenderText("| 216 colour mode |", m.X, m.Y+2, statusBarFg, statusBarBg)
	ui.RenderText("|  grayscale mode |", m.X, m.Y+3, statusBarFg, statusBarBg)
	ui.RenderText("-------------------", m.X, m.Y+4, statusBarFg, statusBarBg)

}

func (m *ModeDialog) Handle(ev termbox.Event) {
	switch ev.Type {

	case termbox.EventMouse:
		// click on mode dialog
		if ev.Key == termbox.MouseLeft {
			m.selectMode(ev.MouseY)
			setState(drawState)
		}
	}
}

func (m *ModeDialog) selectMode(y int) {
	modeY := y - m.Y + 1
	index := termbox.OutputMode(modeY)
	if index >= termbox.OutputNormal && index <= termbox.OutputGrayscale {
		m.callback(index)
	}
}

func (m *ModeDialog) Show(mode termbox.OutputMode, title string, callback func(resultVar interface{})) {
	m.mode = mode
	m.Title = title
	m.callback = callback
	setState(modeState)
}

type BrushDialog struct {
	brush *Brush
	ui.Component
	Title    string
	callback func(resultVar interface{})
}

func newBrushDialog() Dialog {
	brushDialog := &BrushDialog{}

	return brushDialog
}

var totalRunes = 2048
var runesOnLine = 96

func (b *BrushDialog) Render() {
	// make sure component is centred
	termWidth, termHeight := termbox.Size()
	cx := termWidth / 2
	cy := termHeight / 2

	offsetX := runesOnLine / 2
	offsetY := totalRunes / runesOnLine / 2
	b.X = cx - offsetX
	b.Y = cy - offsetY
	for i := rune(0); i < rune(totalRunes); i++ {
		bx := int(i) % runesOnLine
		by := int(i) / runesOnLine
		ui.RenderRune(i, b.X+bx, b.Y+by, statusBarFg, statusBarBg)
	}
	ui.RenderText("--------------brush-------------", b.X, b.Y-1, statusBarFg, statusBarBg)
}

func (b *BrushDialog) Handle(ev termbox.Event) {
	switch ev.Type {

	case termbox.EventMouse:
		// click on mode dialog
		if ev.Key == termbox.MouseLeft {
			b.selectBrush(ev.MouseX, ev.MouseY)
			setState(drawState)
		}
	}
}

func (b *BrushDialog) selectBrush(x, y int) {
	brushY := y - b.Y
	brushX := x - b.X
	index := rune(brushY*runesOnLine + brushX)

	if index >= rune(0) && index <= rune(totalRunes) {
		b.callback(index)
	}
}

func (b *BrushDialog) Show(mode termbox.OutputMode, title string, callback func(resultVar interface{})) {
	b.brush = brush
	b.Title = title
	b.callback = callback
	setState(brushState)
}

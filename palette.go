package main

import (
	"image"

	termbox "github.com/nsf/termbox-go"
	"github.com/telecoda/condraw/ui"
)

type PaletteDialog struct {
	ui.Component
	Title    string
	callback func(colour termbox.Attribute)
}

func newPaletteDialog() *PaletteDialog {
	comp := &PaletteDialog{}
	comp.Component.SetSize(len(comp.content()), len(comp.content())/2)
	return comp
}

func (p *PaletteDialog) content() string {
	return "12345678123456781234567812345678"
}

func (p *PaletteDialog) Render() {
	// make sure component is centred
	termWidth, termHeight := termbox.Size()
	cx := termWidth / 2
	cy := termHeight / 2
	offsetX := p.Width() / 2
	offsetY := p.Height() / 2
	p.X = cx - offsetX
	p.Y = cy - offsetY
	coloursOnLine := 16
	for i := termbox.Attribute(0); i <= termbox.Attribute(255); i++ {
		px := int(i) % coloursOnLine
		py := int(i) / coloursOnLine
		p.renderColour(p.X+px*2, p.Y+py, i)
	}
	ui.RenderText("------------pallette------------", p.X, p.Y-1, statusBarFg, statusBarBg)
}

func (p *PaletteDialog) renderColour(x, y int, colour termbox.Attribute) {
	rect := image.Rect(x, y, x+2, y+1)
	ui.RenderRect(rect, colour)
}

func (p *PaletteDialog) selectColour(x, y int) {
	colourX := (x - p.X) / 2
	colourY := y - p.Y
	index := colourX + colourY*16
	p.callback(termbox.Attribute(index))
}

func (p *PaletteDialog) Handle(ev termbox.Event) {
	switch ev.Type {

	case termbox.EventMouse:
		// click on palette
		if ev.Key == termbox.MouseLeft {
			// calc which colour was clicked
			p.selectColour(ev.MouseX, ev.MouseY)
			setState(drawState)
		}
	}

}

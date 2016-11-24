package ui

import (
	"image"

	termbox "github.com/nsf/termbox-go"
)

/* This code is slowly lapsing into an entire UI rendering framework
   that is not my objective, so I'm putting on the brakes and getting back
   to the features of the drawing program, for now... */

type UIComponent interface {
	Render()
	SetPosition(x, y int)
	GetPosition() (int, int)
	SetSize(width, height int)
	Width() int
	InBounds(x, y int) bool
	Handle(event termbox.Event)
}

func RenderBar(r rune, y int, fg, bg termbox.Attribute) {
	width, _ := termbox.Size()
	for x := 0; x < width; x++ {
		termbox.SetCell(x, y, r, fg, bg)
	}
}

func RenderText(text string, x, y int, fg, bg termbox.Attribute) {
	for i, rune := range text {
		termbox.SetCell(x+i, y, rune, fg, bg)
	}
}

func RenderRune(r rune, x, y int, fg, bg termbox.Attribute) {
	// probably a pointless wrapper function..
	termbox.SetCell(x, y, r, fg, bg)
}

func RenderRect(rect image.Rectangle, colour termbox.Attribute) {
	min := rect.Min
	max := rect.Max
	for x := min.X; x < max.X; x++ {
		for y := min.Y; y < max.Y; y++ {
			termbox.SetCell(x, y, rune(' '), colour, colour)
		}
	}
}

type Component struct {
	X, Y          int
	width, height int
	//parent        *Component
	children []Component
}

func (c *Component) SetPosition(x, y int) {
	c.X = x
	c.Y = y
}

func (c *Component) GetPosition() (int, int) {
	return c.X, c.Y
}

func (c *Component) SetSize(width, height int) {
	c.width = width
	c.height = height
}

func (c *Component) Width() int {
	return c.width
}
func (c *Component) Height() int {
	return c.height
}

func (c *Component) InBounds(x, y int) bool {
	leftX := c.X
	rightX := c.X + c.width
	topY := c.Y
	bottomY := c.Y + c.height
	if x >= leftX && x < rightX && y >= topY && y < bottomY {
		return true
	}
	return false
}

func (c *Component) Handle(event termbox.Event) {

}

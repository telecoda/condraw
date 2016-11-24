package main

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
	"github.com/telecoda/condraw/ui"
)

type StatusBar struct {
	ui.Component
	children []ui.UIComponent
}

func InitStatusBar() *StatusBar {
	statusBar := &StatusBar{
		children: make([]ui.UIComponent, 0),
	}

	// append components
	cursorStatus := newCursorStatus()
	fgStatus := newForegroundStatus()
	bgStatus := newBackgroundStatus()
	brushStatus := newBrushStatus()
	statusBar.children = append(statusBar.children, cursorStatus, fgStatus, bgStatus, brushStatus)
	statusBar.positionAtBottom()

	return statusBar
}

func (s StatusBar) Render() {
	// render bar
	s.positionAtBottom()
	ui.RenderBar('-', s.Y, statusBarFg, statusBarBg)
	// render all nested components
	for i, _ := range s.children {
		child := s.children[i]
		child.Render()
	}
}

func (s *StatusBar) positionAtBottom() {
	// position components based upon their width
	_, height := termbox.Size()
	s.Y = height - 1
	s.X = 0
	posX := 0
	for i, _ := range s.children {
		s.children[i].SetPosition(posX, s.Y)
		posX += s.children[i].Width() + 1
	}
}

func (s *StatusBar) Handle(ev termbox.Event) {
	switch ev.Type {
	case termbox.EventMouse:
		// click on status bar
		if ev.Key == termbox.MouseLeft {
			// check which component is being clicked on
			for i, _ := range s.children {
				child := s.children[i]
				if child.InBounds(ev.MouseX, ev.MouseY) {
					// forward event to component
					child.Handle(ev)
					return
				}
			}
		}
	}

}

type CursorStatus struct {
	ui.Component
}

func newCursorStatus() ui.UIComponent {
	comp := &CursorStatus{}
	// calc rendered Size
	comp.Component.SetSize(len(comp.content(0, 0)), 1)
	return comp
}

func (c CursorStatus) content(x, y int) string {
	return fmt.Sprintf("| Cursor: %03d,%03d |", x, y)
}

func (c CursorStatus) Render() {
	ui.RenderText(c.content(cursorX, cursorY), c.X, c.Y, statusBarFg, statusBarBg)
}

type ForegroundStatus struct {
	ui.Component
}

func newForegroundStatus() ui.UIComponent {
	comp := &ForegroundStatus{}
	comp.Component.SetSize(len(comp.content()), 1)
	return comp
}

func (f ForegroundStatus) content() string {
	return "| FG |"
}

func (f ForegroundStatus) Render() {
	if brush.fg == termbox.ColorBlack {
		ui.RenderText(f.content(), f.X, f.Y, termbox.ColorWhite, brush.fg)
	} else {
		ui.RenderText(f.content(), f.X, f.Y, termbox.ColorBlack, brush.fg)
	}
}

func (f ForegroundStatus) Handle(ev termbox.Event) {
	switch ev.Type {

	case termbox.EventMouse:
		// click on statusbar component
		if ev.Key == termbox.MouseLeft {
			// setup palette for Foreground
			palette.Title = "Set Foreground"
			palette.callback = f.SelectedCallback
			setState(paletteState)
		}
	}

}

func (f ForegroundStatus) SelectedCallback(colour termbox.Attribute) {
	brush.fg = colour
}

type BackgroundStatus struct {
	ui.Component
}

func newBackgroundStatus() ui.UIComponent {
	comp := &BackgroundStatus{}
	comp.Component.SetSize(len(comp.content()), 1)
	return comp
}

func (b BackgroundStatus) content() string {
	return "| BG |"
}

func (b BackgroundStatus) Render() {
	if brush.bg == termbox.ColorBlack {
		ui.RenderText(b.content(), b.X, b.Y, termbox.ColorWhite, brush.bg)
	} else {
		ui.RenderText(b.content(), b.X, b.Y, termbox.ColorBlack, brush.bg)
	}
}

func (b BackgroundStatus) Handle(ev termbox.Event) {
	switch ev.Type {

	case termbox.EventMouse:
		// click on statusbar component
		if ev.Key == termbox.MouseLeft {
			// setup palette for Background
			palette.Title = "Set Background"
			palette.callback = b.SelectedCallback
			setState(paletteState)
		}
	}

}

func (b BackgroundStatus) SelectedCallback(colour termbox.Attribute) {
	brush.bg = colour
}

type BrushStatus struct {
	ui.Component
}

func newBrushStatus() ui.UIComponent {
	comp := &BrushStatus{}
	comp.Component.SetSize(len(comp.content()), 1)
	return comp
}

func (b BrushStatus) content() string {
	return "| Brush: x |"
}

func (b BrushStatus) Render() {
	ui.RenderText(b.content(), b.X, b.Y, statusBarFg, statusBarBg)
	ui.RenderRune(brush.char, b.X+9, b.Y, brush.fg, brush.bg)
}

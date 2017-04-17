package ui

import (
	"errors"
	"fmt"
	"strings"

	cui "github.com/jroimartin/gocui"
)

const delta = 0.2

type HelpWidget struct {
	name string
	x, y int
	w, h int
	body string
}

func NewHelpWidget(name string, x, y int, body string) *HelpWidget {
	lines := strings.Split(body, "\n")

	w := 0
	for _, l := range lines {
		if len(l) > w {
			w = len(l)
		}
	}
	h := len(lines) + 1
	w = w + 1

	return &HelpWidget{name: name, x: x, y: y, w: w, h: h, body: body}
}

func (w *HelpWidget) Layout(g *cui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil {
		if err != cui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, w.body)
	}
	return nil
}

type StatusbarWidget struct {
	name string
	x, y int
	w    int
	val  float64
}

func NewStatusbarWidget(name string, x, y, w int) *StatusbarWidget {
	return &StatusbarWidget{name: name, x: x, y: y, w: w}
}

func (w *StatusbarWidget) SetVal(val float64) error {
	if val < 0 || val > 1 {
		return errors.New("invalid value")
	}
	w.val = val
	return nil
}

func (w *StatusbarWidget) Val() float64 {
	return w.val
}

func (w *StatusbarWidget) Layout(g *cui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+2)
	if err != nil && err != cui.ErrUnknownView {
		return err
	}
	v.Clear()

	rep := int(w.val * float64(w.w-1))
	fmt.Fprint(v, strings.Repeat("▒", rep))
	return nil
}

type ButtonWidget struct {
	name    string
	x, y    int
	w       int
	label   string
	handler func(g *cui.Gui, v *cui.View) error
}

func NewButtonWidget(name string, x, y int, label string, handler func(g *cui.Gui, v *cui.View) error) *ButtonWidget {
	return &ButtonWidget{name: name, x: x, y: y, w: len(label) + 1, label: label, handler: handler}
}

func (w *ButtonWidget) Layout(g *cui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+2)
	if err != nil {
		if err != cui.ErrUnknownView {
			return err
		}
		if _, err := g.SetCurrentView(w.name); err != nil {
			return err
		}
		if err := g.SetKeybinding(w.name, cui.MouseLeft, cui.ModNone, w.handler); err != nil {
			return err
		}
		fmt.Fprint(v, w.label)
	}
	return nil
}

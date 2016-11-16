package main

import termbox "github.com/nsf/termbox-go"

var Menus []*Menu
var currentMenu int

type Menu struct {
	Title string
}

func NewMenu(title string) *Menu {
	m := &Menu{
		Title: title,
	}

	return m
}

func InitMenus() {
	fileMenu := NewMenu("File")
	editMenu := NewMenu("Edit")
	Menus = []*Menu{fileMenu, editMenu}
}

func renderMenuBar() {
	mx := 3
	my := 0
	// render solid bar first
	renderBar('-', 0, menuForeground, menuBackground)
	// overlap menu text on top
	for idx, menu := range Menus {
		if idx == currentMenu {
			renderText(menu.Title, mx, my, menuBackground, menuForeground)
		} else {
			renderText(menu.Title, mx, my, menuForeground, menuBackground)

		}
		mx += len(menu.Title) + 2
	}
}

func renderBar(r rune, y int, fg, bg termbox.Attribute) {
	width, _ := termbox.Size()
	for x := 0; x < width; x++ {
		termbox.SetCell(x, y, r, fg, bg)
	}
}

func renderText(text string, x, y int, fg, bg termbox.Attribute) {
	for i, rune := range text {
		termbox.SetCell(x+i, y, rune, fg, bg)
	}
}

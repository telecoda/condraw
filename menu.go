package main

import "github.com/telecoda/condraw/ui"

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
	ui.RenderBar('-', 0, menuForeground, menuBackground)
	// overlap menu text on top
	for idx, menu := range Menus {
		if idx == currentMenu {
			ui.RenderText(menu.Title, mx, my, menuBackground, menuForeground)
		} else {
			ui.RenderText(menu.Title, mx, my, menuForeground, menuBackground)

		}
		mx += len(menu.Title) + 2
	}
}

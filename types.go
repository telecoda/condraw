package main

type state int

const (
	drawState state = iota
	menuBarState
	menuState
	dialogState
)

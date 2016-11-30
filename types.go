package main

type state int

const (
	drawState state = iota
	paletteState
	modeState
	brushState
)

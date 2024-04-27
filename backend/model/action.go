package model

type Action struct {
	DisplayTxt string
	Hotkey     rune
	Execute    func()
}

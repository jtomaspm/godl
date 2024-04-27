package components

import (
	"github.com/rivo/tview"
)

const catArt = `
 /^--^\     /^--^\     /^--^\
 \____/     \____/     \____/
 /      \   /      \   /      \
 |        | |        | |        |
 \__  __/   \__  __/   \__  __/
 |^|^|^|^|^|^|^|^|^|^|^|^\ \^|^|^|^/ /^|^|^|^|^\ \^|^|^|^|^|^|^|^|^|^|^|^|
 | | | | | | | | | | | | |\ \| | |/ /| | | | | | \ \ | | | | | | | | | | |
 ########################/ /######\ \###########/ /#######################
 | | | | | | | | | | | | \/| | | | \/| | | | | |\/ | | | | | | | | | | | |
 |_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|
`

type DisplayComponent struct {
	Component *tview.TextView
}

func NewDisplayComponent() *DisplayComponent {
	return &DisplayComponent{
		Component: tview.NewTextView().
			SetText(catArt).
			SetTextAlign(tview.AlignCenter).
			SetDynamicColors(true),
	}
}

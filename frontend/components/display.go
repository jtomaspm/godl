package components

import (
	"github.com/rivo/tview"
)

func DisplayScreen() *tview.TextView {
	return tview.NewTextView().
		SetText("").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
}

package components

import "github.com/rivo/tview"

func StatusBar(val string) *tview.TextView {
	b := tview.NewTextView().
		SetText(val).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
	return b
}

func TopBar(app *tview.Application) *tview.Grid {
	widgets := tview.NewGrid().SetBorders(false)

	// Add placeholders to the grid.
	widgets.AddItem(tview.NewTextView().SetText("[1][2]").SetTextAlign(tview.AlignLeft), 0, 0, 1, 2, 0, 0, false)
	widgets.AddItem(StatusBar("Welcome to godl :)"), 0, 2, 1, 14, 0, 0, false)
	widgets.AddItem(Clock(app).SetTextAlign(tview.AlignRight), 0, 15, 1, 2, 0, 0, false)

	return widgets
}

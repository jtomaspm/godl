package components

import "github.com/rivo/tview"

func WelcomeMenu(w []tview.Primitive) *tview.Grid {
	widgets := tview.NewGrid().SetBorders(false)

	for i := 0; i < len(w); i++ {
		widgets.AddItem(w[i], i, 0, 1, 1, 0, 0, false)
	}

	return widgets
}

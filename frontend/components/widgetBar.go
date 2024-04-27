package components

import "github.com/rivo/tview"

func WidgetBar(w []tview.Primitive, vertical bool) *tview.Grid {
	widgets := tview.NewGrid().SetBorders(false)

	for i := 0; i < len(w); i++ {
		if vertical {
			widgets.AddItem(w[i], i, 0, 1, 1, 0, 0, false)
		} else {
			widgets.AddItem(w[i], 0, i, 1, 1, 0, 0, false)
		}
	}

	return widgets
}

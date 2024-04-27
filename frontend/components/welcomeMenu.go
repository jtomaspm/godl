package components

import (
	"fmt"
	"godl/frontend/model"

	"github.com/rivo/tview"
)

func WelcomeMenu(w []model.Action) *tview.List {
	widgets := tview.NewList()

	for i := 0; i < len(w); i++ {
		widgets.AddItem(fmt.Sprintf("%c %s", w[i].Hotkey, w[i].DisplayTxt), "", 0, w[i].Execute)
	}

	widgets.SetCurrentItem(-1)

	return widgets
}

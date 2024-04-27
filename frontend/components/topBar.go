package components

import (
	"godl/backend/model"

	"github.com/rivo/tview"
)

type TopBarComponent struct {
	Component *tview.Grid
	Clock     *tview.TextView
	app       *tview.Application
}

func NewTopBarComponent(app *tview.Application) *TopBarComponent {
	widgets := tview.NewGrid().SetBorders(false)
	c := Clock(app).SetTextAlign(tview.AlignRight)

	widgets.AddItem(tview.NewTextView().SetText(" [1][2]").SetTextAlign(tview.AlignLeft), 0, 0, 1, 15, 0, 0, false)
	widgets.AddItem(c, 0, 15, 1, 2, 0, 0, false)

	return &TopBarComponent{
		Component: widgets,
		Clock:     c,
		app:       app,
	}
}

func (b *TopBarComponent) GetActions() []model.Action {
	return []model.Action{}
}

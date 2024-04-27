package components

import (
	"godl/backend/model"

	"github.com/rivo/tview"
)

type BottomBarComponent struct {
	Component *tview.Flex
	Status    *StatusBarComponent
	clock     *tview.TextView
	app       *tview.Application
}

func NewBottomBarComponent(status string, app *tview.Application) *BottomBarComponent {
	widgets := tview.NewFlex().SetDirection(tview.FlexColumn)
	s := NewStatusBarComponent(status)
	c := Clock(app)

	widgets.AddItem(NewStatusBarComponent("\n").Component, 0, 1, false)
	widgets.AddItem(s.Component, 0, 8, false)
	widgets.AddItem(c, 0, 1, false)

	return &BottomBarComponent{
		Component: widgets,
		Status:    s,
		clock:     c,
		app:       app,
	}
}

func (b *BottomBarComponent) GetActions() []model.Action {
	return []model.Action{}
}

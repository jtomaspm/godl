package components

import (
	"godl/backend/model"

	"github.com/rivo/tview"
)

type TopBarComponent struct {
	Component *tview.Flex
}

func NewTopBarComponent() *TopBarComponent {
	widgets := tview.NewFlex()

	return &TopBarComponent{
		Component: widgets,
	}
}

func (b *TopBarComponent) GetActions() []model.Action {
	return []model.Action{}
}

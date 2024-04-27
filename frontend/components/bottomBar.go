package components

import (
	"godl/backend/model"

	"github.com/rivo/tview"
)

type BottomBarComponent struct {
	Component *tview.Grid
	menu      *WelcomeMenuComponent
}

func NewBottomBarComponent(actions []model.Action) *BottomBarComponent {
	bottomBar := []tview.Primitive{}
	m := NewWelcomeMenuComponent(actions)

	bottomBar = append(bottomBar, StatusBar(""))
	bottomBar = append(bottomBar, m.Coponent)
	bottomBar = append(bottomBar, StatusBar(""))

	wb := WidgetBar(bottomBar, false)
	return &BottomBarComponent{
		Component: wb,
		menu:      m,
	}
}

func (b *BottomBarComponent) GetActions() []model.Action {
	return b.menu.GetActions()
}

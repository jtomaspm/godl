package components

import (
	"godl/frontend/model"

	"github.com/rivo/tview"
)

type BottomBarComponent struct {
	Component *tview.Grid
}

func NewBottomBarComponent(actions []model.Action) *BottomBarComponent {
	bottomBar := []tview.Primitive{}

	bottomBar = append(bottomBar, StatusBar(""))
	bottomBar = append(bottomBar, WelcomeMenu(actions))
	bottomBar = append(bottomBar, StatusBar(""))

	wb := WidgetBar(bottomBar, false)
	return &BottomBarComponent{
		Component: wb,
	}
}

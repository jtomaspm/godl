package components

import (
	"godl/backend/model"

	"github.com/rivo/tview"
)

type MainMenuComponent struct {
	Component *tview.Grid
	menu      *WelcomeMenuComponent
}

func NewMainMenuComponent(actions []model.Action) *MainMenuComponent {
	mainMenu := []tview.Primitive{}
	m := NewWelcomeMenuComponent(actions)

	mainMenu = append(mainMenu, NewStatusBarComponent("").Component)
	mainMenu = append(mainMenu, m.Coponent)
	mainMenu = append(mainMenu, NewStatusBarComponent("").Component)

	wb := WidgetBar(mainMenu, false)
	return &MainMenuComponent{
		Component: wb,
		menu:      m,
	}
}

func (b *MainMenuComponent) GetActions() []model.Action {
	return b.menu.GetActions()
}

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

func (b *MainMenuComponent) GetActions() []model.Action {
	return b.menu.GetActions()
}

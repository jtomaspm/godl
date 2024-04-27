package components

import (
	"fmt"
	"godl/backend/model"

	"github.com/rivo/tview"
)

type WelcomeMenuComponent struct {
	Coponent          *tview.List
	navigationActions []model.Action
	actions           []model.Action
}

func NewWelcomeMenuComponent(actions []model.Action) *WelcomeMenuComponent {
	widgets := tview.NewList()

	for i := 0; i < len(actions); i++ {
		widgets.AddItem(fmt.Sprintf("(%c) %s", actions[i].Hotkey, actions[i].DisplayTxt), "", 0, actions[i].Execute)
	}

	return &WelcomeMenuComponent{
		Coponent: widgets,
		actions:  actions,
		navigationActions: []model.Action{
			{
				DisplayTxt: "Up",
				Hotkey:     'k',
				Execute: func() {
					widgets.SetCurrentItem(widgets.GetCurrentItem() - 1)
				},
			},
			{
				DisplayTxt: "Down",
				Hotkey:     'j',
				Execute: func() {
					cur := widgets.GetCurrentItem()
					if cur == widgets.GetItemCount()-1 {
						widgets.SetCurrentItem(0)
					} else {
						widgets.SetCurrentItem(cur + 1)
					}
				},
			},
		},
	}
}

func (b *WelcomeMenuComponent) GetActions() []model.Action {
	return append(b.actions, b.navigationActions...)
}

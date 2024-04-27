package components

import (
	"godl/backend/model"

	"github.com/rivo/tview"
)

type StatusBarComponent struct {
	Component *tview.TextView
}

func NewStatusBarComponent(val string) *StatusBarComponent {
	b := tview.NewTextView().
		SetText(val).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	return &StatusBarComponent{
		Component: b,
	}
}

func (b *StatusBarComponent) SetStatus(status string) {
	b.Component.SetText(status)
}

func (b *StatusBarComponent) SetStatusNL(status string) {
	b.Component.SetText("\n" + status)
}

func (b *StatusBarComponent) GetActions() []model.Action {
	return []model.Action{}
}

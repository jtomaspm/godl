package events

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type GenericMediator struct {
	App *tview.Application
}

func NewGenericMediator(a *tview.Application) *GenericMediator {
	return &GenericMediator{
		App: a,
	}
}

func (m *GenericMediator) Handle(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "q":
		m.App.Stop()
	}
	return event
}

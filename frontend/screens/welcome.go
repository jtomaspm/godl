package screens

import (
	"godl/backend/events"
	"godl/backend/model"
	"godl/frontend/components"
	"godl/frontend/widgets"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Welcome_s struct {
	App      *tview.Application
	mediator *events.Mediator
	actions  []model.Action
}

func NewWelcomeScreen(a *tview.Application, m events.Mediator) *Welcome_s {
	return &Welcome_s{
		App:      a,
		mediator: &m,
	}
}

func (s *Welcome_s) Draw() {
	topBar := components.NewTopBarComponent(s.App)
	s.actions = append(s.actions, topBar.GetActions()...)

	display := components.NewDisplayComponent()
	s.actions = append(s.actions, display.GetActions()...)

	bottomBar := components.NewBottomBarComponent(widgets.WelcomeWidgets(topBar))
	s.actions = append(s.actions, bottomBar.GetActions()...)

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(topBar.Component, 0, 1, false).
		AddItem(display.Component, 0, 5, false).
		AddItem(bottomBar.Component, 0, 7, false)

	s.App.SetRoot(layout, true)
	s.mediator.AddHandlers(s.actions)
	s.App.SetInputCapture(s.HandleEvents)
}

func (s *Welcome_s) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	s.mediator.Handle(event.Rune())
	return event
}

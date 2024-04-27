package screens

import (
	"godl/backend/events"
	"godl/frontend/components"
	"godl/frontend/widgets"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Welcome_s struct {
	App      *tview.Application
	mediator *events.Mediator
}

func NewWelcomeScreen(a *tview.Application, m events.Mediator) *Welcome_s {
	return &Welcome_s{
		App:      a,
		mediator: &m,
	}
}

func (s *Welcome_s) Draw() {
	topBar := components.NewTopBarComponent(s.App)

	display := components.NewDisplayComponent()

	bottomBar := components.NewBottomBarComponent(widgets.WelcomeWidgets(topBar))

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(topBar.Component, 0, 1, false).
		AddItem(display.Component, 0, 5, false).
		AddItem(bottomBar.Component, 0, 7, false)

	s.App.SetRoot(layout, true)

	s.App.SetInputCapture(s.HandleEvents)
}

func (s *Welcome_s) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	s.mediator.Handle(event.Rune())
	return event
}

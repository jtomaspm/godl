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
	app      *tview.Application
	mediator *events.Mediator
	actions  []model.Action
}

func NewWelcomeScreen(a *tview.Application, m events.Mediator) *Welcome_s {
	return &Welcome_s{
		app:      a,
		mediator: &m,
	}
}

func (s *Welcome_s) Draw() {
	topBar := components.NewTopBarComponent()
	s.actions = append(s.actions, topBar.GetActions()...)

	display := components.NewDisplayComponent()
	s.actions = append(s.actions, display.GetActions()...)

	bottomBar := components.NewBottomBarComponent("\nWelcome to godl :)", s.app)
	s.actions = append(s.actions, bottomBar.GetActions()...)

	mainMenu := components.NewMainMenuComponent(widgets.WelcomeWidgets(bottomBar.Status))
	s.actions = append(s.actions, mainMenu.GetActions()...)

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(topBar.Component, 0, 2, false).
		AddItem(display.Component, 0, 10, false).
		AddItem(mainMenu.Component, 0, 12, false).
		AddItem(bottomBar.Component, 0, 1, false)

	s.app.SetRoot(layout, true)
	s.mediator.AddHandlers(s.actions)
	s.app.SetInputCapture(s.HandleEvents)
}

func (s *Welcome_s) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	s.mediator.Handle(event.Rune())
	return event
}

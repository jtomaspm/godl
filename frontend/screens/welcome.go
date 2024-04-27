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

	statusBar := components.NewStatusBarComponent("Welcome to godl :)")
	s.actions = append(s.actions, statusBar.GetActions()...)

	mainMenu := components.NewMainMenuComponent(widgets.WelcomeWidgets(statusBar))
	s.actions = append(s.actions, mainMenu.GetActions()...)

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(topBar.Component, 0, 1, false).
		AddItem(display.Component, 0, 8, false).
		AddItem(mainMenu.Component, 0, 9, false).
		AddItem(statusBar.Component, 0, 1, false)

	s.App.SetRoot(layout, true)
	s.mediator.AddHandlers(s.actions)
	s.App.SetInputCapture(s.HandleEvents)
}

func (s *Welcome_s) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	s.mediator.Handle(event.Rune())
	return event
}

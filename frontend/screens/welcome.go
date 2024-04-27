package screens

import (
	"godl/frontend/components"
	"godl/frontend/events"
	"godl/frontend/widgets"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Welcome_s struct {
	App *tview.Application
}

func NewWelcomeScreen(a *tview.Application) *Welcome_s {
	return &Welcome_s{
		App: a,
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
	m := events.NewGenericMediator(s.App)
	return m.Handle(event)
}

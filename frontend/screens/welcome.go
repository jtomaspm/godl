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

	topBar := components.TopBar(s.App)

	display := components.DisplayScreen()

	bottomBar := []tview.Primitive{}

	bottomBar = append(bottomBar, components.StatusBar("w1"))
	bottomBar = append(bottomBar, components.WelcomeMenu(widgets.WelcomeWidgets()))
	bottomBar = append(bottomBar, components.StatusBar("w3"))

	wb := components.WidgetBar(bottomBar, false)

	// Add the top bar and grid to the application layout.
	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(topBar, 0, 1, false).
		AddItem(display, 0, 4, false).
		AddItem(wb, 0, 8, false)

	s.App.SetRoot(layout, true)

	s.App.SetInputCapture(s.HandleEvents)
}

func (s *Welcome_s) HandleEvents(event *tcell.EventKey) *tcell.EventKey {
	m := events.NewGenericMediator(s.App)
	return m.Handle(event)
}

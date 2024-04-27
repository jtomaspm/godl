package app

import (
	"godl/backend/events"
	"godl/backend/model"
	"godl/frontend/screens"

	"github.com/rivo/tview"
)

type AppContext struct {
	App          *tview.Application
	CurrentSreen screens.Screen
}

func (ac *AppContext) initApp() {
	ac.CurrentSreen.Draw()
}

func NewAppContext() *AppContext {
	a := tview.NewApplication()
	ac := &AppContext{
		App:          a,
		CurrentSreen: screens.NewWelcomeScreen(a, *baseMediator(a)),
	}
	ac.initApp()
	return ac
}

func baseMediator(app *tview.Application) *events.Mediator {
	m := events.NewMediator()

	m.AddHandler(model.Action{
		DisplayTxt: "Quit",
		Hotkey:     'q',
		Execute: func() {
			app.Stop()
		},
	})

	return m
}

func (ac *AppContext) Run() error {
	return ac.App.Run()
}

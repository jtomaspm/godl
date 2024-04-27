package app

import (
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
		CurrentSreen: screens.NewWelcomeScreen(a),
	}
	ac.initApp()
	return ac
}

func NewAppContextFromApp(a *tview.Application) *AppContext {
	ac := &AppContext{
		App:          a,
		CurrentSreen: screens.NewWelcomeScreen(a),
	}
	ac.initApp()
	return ac
}

func (ac *AppContext) Run() error {
	return ac.App.Run()
}

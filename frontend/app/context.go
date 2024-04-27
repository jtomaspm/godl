package app

import (
	"godl/backend/events"
	"godl/backend/model"
	"godl/frontend/router"
	"godl/frontend/screens"

	"github.com/rivo/tview"
)

type AppContext struct {
	app    *tview.Application
	router router.Router
}

func (ac *AppContext) initApp() {
	ac.router.Draw("welcome")
}

func (ac *AppContext) addRoutes() {
	ac.router.AddRoute("welcome", screens.NewWelcomeScreen(ac.app, *baseMediator(ac.app)))
}

func NewAppContext() *AppContext {
	a := tview.NewApplication()
	ac := &AppContext{
		app:    a,
		router: *router.NewRouter(),
	}
	ac.addRoutes()
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
	return ac.app.Run()
}

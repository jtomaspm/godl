package app

import (
	"godl/backend/events"
	"godl/backend/model"
	"godl/frontend/components"

	"github.com/rivo/tview"
)

type AppContext struct {
	app    *tview.Application
	navBar *components.NavBarComponent
}

func (ac *AppContext) initApp() {
}

func (ac *AppContext) addRoutes() {
}

func NewAppContext() *AppContext {
	a := tview.NewApplication()
	ac := &AppContext{
		app:    a,
		navBar: components.NewNavBarComponent(),
	}
	ac.addRoutes()
	ac.setupNavBar()
	ac.initApp()
	return ac
}

func baseMediator(app *tview.Application) *events.Mediator {
	m := events.NewMediator()

	m.AddHandler(&model.Action{
		DisplayTxt: "Quit",
		Hotkey:     'q',
		Execute: func() {
			app.Stop()
		},
	})

	return m
}

func (ac *AppContext) setupNavBar() {
	ac.navBar.SetActive("welcome")
}

func (ac *AppContext) Run() error {
	return ac.app.Run()
}

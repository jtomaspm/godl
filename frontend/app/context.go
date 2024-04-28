package app

import (
	"godl/backend/events"
	"godl/backend/model"
	"godl/frontend/components"
	"godl/frontend/router"
	"godl/frontend/screens"

	"github.com/rivo/tview"
)

type AppContext struct {
	app    *tview.Application
	router router.Router
	navBar *components.NavBarComponent
}

func (ac *AppContext) initApp() {
	ac.router.Draw("welcome")
}

func (ac *AppContext) addRoutes() {
	ac.router.AddRoute("welcome", screens.NewWelcomeScreen(ac.app, *baseMediator(ac.app), ac.navBar))
}

func NewAppContext() *AppContext {
	a := tview.NewApplication()
	ac := &AppContext{
		app:    a,
		router: *router.NewRouter(),
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
	ac.navBar.AddTab(*components.NewTabComponent("welcome", ac.router.GetResolution("welcome")))
	ac.navBar.AddTab(*components.NewTabComponent("settings", func() {}))
	ac.navBar.SetActive("welcome")
}

func (ac *AppContext) Run() error {
	return ac.app.Run()
}

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
	navBar *components.NavBarComponent
	router *router.Router
}

func NewAppContext() *AppContext {
	a := tview.NewApplication()
	r := router.NewRouter()
	ac := &AppContext{
		app:    a,
		navBar: components.NewNavBarComponent(r),
		router: r,
	}
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

func (ac *AppContext) initApp() {

	ac.navBar.AddTab("welcome", screens.NewWelcomeScreen(ac.app, *baseMediator(ac.app), ac.navBar).Draw)
	ac.navBar.AddTab("settings", screens.NewWelcomeScreen(ac.app, *baseMediator(ac.app), ac.navBar).Draw)

	if err := ac.navBar.SetCurrent(0); err != nil {
		panic(err)
	}
}

func (ac *AppContext) Run() error {
	return ac.app.Run()
}

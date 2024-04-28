package app

import (
	"godl/backend/events"
	"godl/backend/model"
	"godl/frontend/components"
	"godl/frontend/router"

	"github.com/rivo/tview"
)

type AppContext struct {
	app    *tview.Application
	navBar *components.NavBarComponent
	router *router.Router
}

func (ac *AppContext) initApp() {
}

func (ac *AppContext) addRoutes() {
}

func NewAppContext() *AppContext {
	a := tview.NewApplication()
	r := router.NewRouter()
	ac := &AppContext{
		app:    a,
		navBar: components.NewNavBarComponent(r),
		router: r,
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

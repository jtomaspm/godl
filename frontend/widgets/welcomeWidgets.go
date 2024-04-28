package widgets

import (
	"godl/backend/model"
	"godl/frontend/components"
)

func WelcomeWidgets(statusBar *components.StatusBarComponent) []model.Action {
	return []model.Action{
		{
			DisplayTxt: "Change working directory",
			Hotkey:     'f',
			Execute: func() {
				statusBar.SetStatusNL("Test 1")
			},
		},
		{
			DisplayTxt: "Test 2",
			Hotkey:     'z',
			Execute: func() {
				statusBar.SetStatusNL("Test 2")
			},
		},
		{
			DisplayTxt: "Test 3",
			Hotkey:     'x',
			Execute: func() {
				statusBar.SetStatusNL("Test 3")
			},
		},
		{
			DisplayTxt: "Test 4",
			Hotkey:     'c',
			Execute: func() {
				statusBar.SetStatusNL("Test 4")
			},
		},
		{
			DisplayTxt: "Test 5",
			Hotkey:     'v',
			Execute: func() {
				statusBar.SetStatusNL("Test 5")
			},
		},
	}
}

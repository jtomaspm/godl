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
			Hotkey:     '2',
			Execute: func() {
				statusBar.SetStatusNL("Test 2")
			},
		},
		{
			DisplayTxt: "Test 3",
			Hotkey:     '3',
			Execute: func() {
				statusBar.SetStatusNL("Test 3")
			},
		},
		{
			DisplayTxt: "Test 4",
			Hotkey:     '4',
			Execute: func() {
				statusBar.SetStatusNL("Test 4")
			},
		},
		{
			DisplayTxt: "Test 5",
			Hotkey:     '5',
			Execute: func() {
				statusBar.SetStatusNL("Test 5")
			},
		},
	}
}

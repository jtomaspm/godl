package widgets

import (
	"godl/frontend/components"
	"godl/frontend/model"
)

func WelcomeWidgets(topBar *components.TopBarComponent) []model.Action {
	return []model.Action{
		{
			DisplayTxt: "Change working directory",
			Hotkey:     '1',
			Execute: func() {
				topBar.SetStatus("Test 1")
			},
		},
		{
			DisplayTxt: "Test 2",
			Hotkey:     '2',
			Execute: func() {
				topBar.SetStatus("Test 2")
			},
		},
		{
			DisplayTxt: "Test 3",
			Hotkey:     '3',
			Execute: func() {
				topBar.SetStatus("Test 3")
			},
		},
		{
			DisplayTxt: "Test 4",
			Hotkey:     '4',
			Execute: func() {
				topBar.SetStatus("Test 4")
			},
		},
		{
			DisplayTxt: "Test 5",
			Hotkey:     '5',
			Execute: func() {
				topBar.SetStatus("Test 5")
			},
		},
	}
}

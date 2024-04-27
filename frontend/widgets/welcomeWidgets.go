package widgets

import (
	"godl/frontend/components"

	"github.com/rivo/tview"
)

func WelcomeWidgets() []tview.Primitive {
	return []tview.Primitive{
		components.StatusBar("w1"),
		components.StatusBar("w2"),
		components.StatusBar("w3"),
		components.StatusBar("w4"),
		components.StatusBar("w5"),
		components.StatusBar("w6"),
	}
}

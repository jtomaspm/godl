package components

import "github.com/rivo/tview"

func StatusBar(val string) *tview.TextView {
	b := tview.NewTextView().
		SetText(val).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
	return b
}

type TopBarComponent struct {
	Component *tview.Grid
	StatusBar *tview.TextView
	Clock     *tview.TextView
	app       *tview.Application
}

func NewTopBarComponent(app *tview.Application) *TopBarComponent {
	widgets := tview.NewGrid().SetBorders(false)
	sb := StatusBar("Welcome to godl :)")
	c := Clock(app).SetTextAlign(tview.AlignRight)

	widgets.AddItem(tview.NewTextView().SetText("[1][2]").SetTextAlign(tview.AlignLeft), 0, 0, 1, 2, 0, 0, false)
	widgets.AddItem(sb, 0, 2, 1, 14, 0, 0, false)
	widgets.AddItem(c, 0, 15, 1, 2, 0, 0, false)

	return &TopBarComponent{
		Component: widgets,
		StatusBar: sb,
		Clock:     c,
		app:       app,
	}
}

func (b *TopBarComponent) SetStatus(status string) {
	b.app.QueueUpdateDraw(func() {
		b.StatusBar.SetText(status)
	})
}

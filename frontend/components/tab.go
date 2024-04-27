package components

import "github.com/rivo/tview"

type TabComponent struct {
	Component *tview.TextView
	label     string
}

func NewTabComponent(label string) *TabComponent {
	b := tview.NewTextView().
		SetText(label).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	c := TabComponent{
		Component: b,
		label:     label,
	}

	c.SetInactive()

	return &c
}

func (t *TabComponent) SetActive() {
	t.Component.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)
	t.Component.SetTextColor(tview.Styles.ContrastBackgroundColor)
}

func (t *TabComponent) SetInactive() {
	t.Component.SetBackgroundColor(tview.Styles.ContrastBackgroundColor)
	t.Component.SetTextColor(tview.Styles.PrimaryTextColor)
}

func PlaceholderTab() *TabComponent {
	b := tview.NewTextView().
		SetText("").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	c := TabComponent{
		Component: b,
		label:     "",
	}

	c.SetInactive()

	return &c
}

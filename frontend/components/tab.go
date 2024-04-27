package components

import "github.com/rivo/tview"

type TabComponent struct {
	Component *tview.TextView
	label     string
	callback  func()
}

func NewTabComponent(label string, callback func()) *TabComponent {
	b := tview.NewTextView().
		SetText(label).
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true)

	c := TabComponent{
		Component: b,
		label:     label,
		callback:  callback,
	}

	c.SetInactive()

	return &c
}

func (t *TabComponent) SetActive() {
	t.Component.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)
	t.Component.SetTextColor(tview.Styles.ContrastBackgroundColor)
}

func (t *TabComponent) SetInactive() {
	t.Component.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)
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

func (t *TabComponent) Activate() {
	t.callback()
}

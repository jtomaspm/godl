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

	SetTabInactive(c.Component)

	return &c
}

func SetTabActive(t *tview.TextView) {
	t.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)
	t.SetTextColor(tview.Styles.ContrastBackgroundColor)
}

func SetTabInactive(t *tview.TextView) {
	t.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)
	t.SetTextColor(tview.Styles.PrimaryTextColor)
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

	SetTabInactive(c.Component)

	return &c
}

func (t *TabComponent) Activate() {
	t.callback()
}

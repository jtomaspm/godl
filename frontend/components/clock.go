package components

import (
	"time"

	"github.com/rivo/tview"
)

func Clock(app *tview.Application) *tview.TextView {
	b := tview.NewTextView().
		SetTextAlign(tview.AlignRight).
		SetDynamicColors(true)
	go func() {
		for {
			now := time.Now()
			currentTime := now.Format("\n🕑 15:04 ")
			app.QueueUpdateDraw(func() {
				b.SetText(currentTime)
			})
			time.Sleep(time.Second)
		}
	}()
	return b
}

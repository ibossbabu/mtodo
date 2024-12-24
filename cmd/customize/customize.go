package customize

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func TextView() *tview.TextView {
	text := tview.NewTextView().
		SetTextColor(tcell.ColorHotPink).
		SetTextAlign(tview.AlignCenter)
	text.SetBackgroundColor(tcell.ColorBlack)
	text.SetBorder(true)
	text.SetTitle("TODO LIST")
	text.SetBorderColor(tcell.ColorRoyalBlue)

	return text
}

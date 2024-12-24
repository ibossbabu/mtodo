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
	text.SetBorderColor(tcell.ColorRoyalBlue)
  	text.SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
		// Draw a horizontal line 
		centerY := y + height/4
		for cx := x + 1; cx < x+width-1; cx++ {
			screen.SetContent(cx, centerY, tview.BoxDrawingsLightHorizontal, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
		}
		tview.Print(screen, "Mode", x+2, centerY-1, width-2, tview.AlignLeft, tcell.ColorYellow)

		return x + 1, centerY + 1, width - 2, height - (centerY + 1 - y)
	})
	return text
}

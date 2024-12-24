package table

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Table() *tview.Table {
	tbl := tview.NewTable().
		SetSelectable(true, false).
		SetSelectedStyle(
			tcell.StyleDefault.
				Background(tcell.ColorRoyalBlue).
				Foreground(tcell.ColorWhite),
		)
	tbl.SetBorder(true)
	tbl.SetBorderColor(tcell.ColorRoyalBlue)
	tbl.SetTitle("TODO LIST")

	tbl.SetSelectedStyle(
		tcell.StyleDefault.
			Background(tcell.ColorDarkSlateGrey).
			Foreground(tcell.ColorWhite),
	)

	return tbl
}

package table

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Table() *tview.Table {
	tbl := tview.NewTable().
		SetSelectable(true, false)
	tbl.SetBorder(true)
	tbl.SetBorderColor(tcell.ColorRoyalBlue)
	tbl.SetTitle("TODO LIST")
  tbl.SetTitleColor(tcell.ColorLime)
	tbl.SetSelectedStyle(
		tcell.StyleDefault.
			Background(tcell.ColorDarkSlateGrey). //  highlight color
	    Foreground(tcell.ColorPaleGoldenrod),    //highlight text color 
	)
	return tbl
}

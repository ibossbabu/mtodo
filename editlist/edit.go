package editlist

import 	("github.com/rivo/tview"
  "github.com/gdamore/tcell/v2"
)

func Edit() *tview.InputField {
  edit := tview.NewInputField().
  		SetLabel("Rename: ").
		SetFieldBackgroundColor(tcell.ColorRoyalBlue)
  return edit
}

package main

import (
	"todo/customize"
	"todo/table"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func input() *tview.InputField {
	inputfield := tview.NewInputField().
		SetLabel("Add: "). 
    SetFieldBackgroundColor(tcell.ColorRoyalBlue)
	return inputfield
}

func main() {
	text := customize.TextView()
	app := tview.NewApplication()
	inputfield := input()
	tbl := table.Table()
	row := 0 
  updateMode := func(mode string) {
        text.SetText(mode)
    }
	inputfield.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			newText := inputfield.GetText()
			if newText != "" {
				cell := tview.NewTableCell(newText).
					SetExpansion(1).          // This makes the cell expand to fill space
					SetAlign(tview.AlignLeft) 

				tbl.SetCell(row, 0, cell)
				row++
				inputfield.SetText("") // Clear the input field
				// Move selection to new row
				tbl.Select(row-1, 0)
			}
		}
	})

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tbl, 0, 6, false).
		AddItem(text, 0, 1, false).
		AddItem(inputfield, 1, 1, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			if event.Rune() == 'i' && app.GetFocus() != inputfield {
				app.SetFocus(inputfield)
        updateMode("INSERT")
				return nil
			}
      if event.Rune() == 'q' && app.GetFocus() != inputfield {
				app.Stop()
				return nil
      }
      if event.Rune() == 'v' && app.GetFocus() != inputfield{
        app.SetFocus(tbl)
        updateMode("VISUAL")
        return nil
      }
		case tcell.KeyEsc:
			app.SetFocus(nil)
      updateMode("NORMAL")
			// Ensure there's at least one row to select
			if row > 0 {
				tbl.Select(tbl.GetSelection())
			}
		}
		return event
	})

	if err := app.SetRoot(flex, true).SetFocus(nil).Run(); err != nil {
		panic(err)
	}
}

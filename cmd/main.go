package main

import (
	"fmt"
	"os"
	"todo/customize"
	"todo/table"
  "todo/editlist"
  "todo/storage"
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
  edit := editlist.Edit()
	tbl := table.Table()
	tasks := []storage.Task{}
	row := 0
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tbl, 0, 6, false).
		AddItem(text, 0, 1, false).
		AddItem(inputfield, 0, 0, false).
		AddItem(edit, 0, 0, false)

// Load existing tasks
	loadedTasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
	} else {
		tasks = loadedTasks
		// Populate table with loaded tasks
		for _, task := range tasks {
			checkbox := " "
			if task.Checked {
				checkbox = "󰄳 "
			}
			tbl.SetCell(row, 0, tview.NewTableCell(checkbox).
				SetAlign(tview.AlignCenter).
				SetSelectable(true))
			tbl.SetCell(row, 2, tview.NewTableCell(task.Text).
				SetExpansion(1).
				SetAlign(tview.AlignLeft))
			row++
		}
	}
  
	updateMode := func(mode string) {
		text.SetText(mode)
      switch mode {
    case "RENAME":
        flex.ResizeItem(inputfield, 0, 0)
        flex.ResizeItem(edit, 1, 1)
    case "INSERT":
        flex.ResizeItem(edit, 0, 0)
        flex.ResizeItem(inputfield, 1, 1)
    case "NORMAL":
        flex.ResizeItem(edit, 0, 0)
        flex.ResizeItem(inputfield, 0, 0)
    default:
        flex.ResizeItem(inputfield, 0, 0)
        flex.ResizeItem(edit, 1, 1)
    }
	}

	checkboxCell := func() *tview.TableCell {
		checkbox := " "
		return tview.NewTableCell(checkbox).
			SetAlign(tview.AlignCenter).
			SetSelectable(true)
	}

	inputfield.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			newText := inputfield.GetText()
			if newText != "" {
				tbl.SetCell(row, 0, checkboxCell())
				cell := tview.NewTableCell(newText).
					SetExpansion(1).
					SetAlign(tview.AlignLeft)
				tbl.SetCell(row, 2, cell)
				// Add new task to tasks slice
				tasks = append(tasks, storage.Task{
					Text:    newText,
					Checked: false,
				})
				// Save tasks to file
				if err := storage.SaveTasks(tasks); err != nil {
					fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
				}
				row++
				inputfield.SetText("")
				tbl.Select(row-1, 0)
			}
		}
	})

	edit.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			newText := edit.GetText()
			selectedRow, _ := tbl.GetSelection()
			
			if newText != "" && selectedRow < len(tasks) {
				// Update the task text
				tasks[selectedRow].Text = newText
				// Update the table cell
				cell := tview.NewTableCell(newText).
					SetExpansion(1).
					SetAlign(tview.AlignLeft)
				tbl.SetCell(selectedRow, 2, cell)
				// Save tasks to file
				if err := storage.SaveTasks(tasks); err != nil {
					fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
				}
			}
			edit.SetText("")
			app.SetFocus(tbl)
			updateMode("VISUAL")
		}
		if key == tcell.KeyEsc {
			// Cancel rename operation
			edit.SetText("")
			app.SetFocus(tbl)
			updateMode("VISUAL")
		}
	})
	tbl.SetSelectedFunc(func(row, col int) {
		if col == 0 {
			cell := tbl.GetCell(row, 0)
			isChecked := cell.Text == " "
			if isChecked {
				cell.SetText("󰄳 ")
			} else {
				cell.SetText(" ")
			}
			tbl.SetCell(row, 0, cell)
			// Update task checked status
			tasks[row].Checked = isChecked
			// Save updated tasks
			if err := storage.SaveTasks(tasks); err != nil {
				fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
			}
		}
	})

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			  if event.Rune() == 'i' && text.GetText(true) != "EDIT" && app.GetFocus() != inputfield {
				app.SetFocus(inputfield)
				updateMode("INSERT")
				return nil
			}
			if event.Rune() == 'q' && app.GetFocus() != inputfield {
				app.Stop()
				return nil
			}
			if event.Rune() == 'v' && app.GetFocus() != inputfield {
				app.SetFocus(tbl)
				updateMode("VISUAL")
				return nil
			}
			if event.Rune() == 'r' && text.GetText(true) == "VISUAL" {
				selectedRow, _ := tbl.GetSelection()
				if selectedRow >= 0 && selectedRow < len(tasks) {
					// Set initial text to current task text
					edit.SetText(tasks[selectedRow].Text)
					app.SetFocus(edit)
					updateMode("EDIT")
					return nil
				}
			}
		case tcell.KeyEsc:
			app.SetFocus(nil)
			updateMode("NORMAL")
			if row > 0 {
				tbl.Select(tbl.GetSelection())
			}
		}
		return event
	})

	if err := app.SetRoot(flex, true).SetFocus(tbl).Run(); err != nil {
		panic(err)
	}
}

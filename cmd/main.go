package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
  "todo/cmd/customize"
)

func input() *tview.InputField {
	inputfield := tview.NewInputField().
		SetLabel("Add: ")
	return inputfield
}


func main() {
  text := customize.TextView()
	app := tview.NewApplication()
	inputfield := input()
	text.SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
		// Draw a horizontal line across the middle of the box
		centerY := y + height/5
		for cx := x + 1; cx < x+width-1; cx++ {
			screen.SetContent(cx, centerY, tview.BoxDrawingsLightHorizontal, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
		}
		tview.Print(screen, " Center Line ", x+1, centerY-1, width-2, tview.AlignLeft, tcell.ColorYellow)

		return x + 1, centerY + 2, width - 2, height - (centerY + 1 - y)
	})

	inputfield.SetDoneFunc(func(key tcell.Key) {
 if key == tcell.KeyEnter {
            currentText := text.GetText(false) // Get the current text without regions
            newText := currentText + "\n" + "\n"+ inputfield.GetText()
            text.SetText(newText) // Update the TextView with the updated text
            inputfield.SetText("") // Clear the input field
        }
	})

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(text, 0, 1, false).
		AddItem(inputfield, 1, 1, true) // Set inputfield to be in focus initially

	if err := app.SetRoot(flex, true).SetFocus(inputfield).Run(); err != nil {
		panic(err)
	}
}

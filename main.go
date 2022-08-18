package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// App struct holds the application state
type App struct {
	output *widget.Label
}

// makeUI defines the widgets that will used to build the application's interface
func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button, *widget.Button) {
	output := widget.NewLabel("Hello Fyne!")

	entry := widget.NewEntry()

	actionBtn := widget.NewButton("Set Text", func() {
		app.output.SetText(entry.Text)
	})
	actionBtn.Importance = widget.HighImportance

	clearBtn := widget.NewButton("Reset", func() {
		entry.SetText("")
		app.output.SetText("Hello Fyne!")
	})

	app.output = output

	return output, entry, actionBtn, clearBtn
}

var myApp App

func main() {
	a := app.New()
	w := a.NewWindow("Hello Fyne!")

	output, entry, actionBtn, clearBtn := myApp.makeUI()

	buttons := container.NewHBox(actionBtn, clearBtn)
	windowContent := container.NewVBox(output, entry, buttons)

	w.SetContent(windowContent)
	w.Resize(fyne.Size{Width: 500, Height: 500})
	w.ShowAndRun()
}

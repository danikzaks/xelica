package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/danikzaks/xelica/internal/ui"
)

func main() {
	myApp := app.New()

	window := myApp.NewWindow("Xelica Tool")

	ui.SetupUI(window)

	window.ShowAndRun()
}

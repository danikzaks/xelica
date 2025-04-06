package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/danikzaks/xelica/internal/ui"
)

func main() {
	a := app.New()

	window := a.NewWindow("Xelica Tool")
	
	ui.SetupUI(window)

	window.ShowAndRun()
}

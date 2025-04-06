package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/danikzaks/xelica/internal/ui"
)

func main() {
	a := app.New()

	w := a.NewWindow("Xelica Tool")

	w.SetMainMenu(ui.MakeMenu(a, w))
	ui.SetupUI(w)

	w.ShowAndRun()
}

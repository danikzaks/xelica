package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/danikzaks/xelica/internal/ui"
)

func main() {
	a := app.NewWithID("com.danikzaks.xelica")

	w := a.NewWindow("Xelica Tool")
	w.SetMainMenu(ui.MakeMenu(a, w))
	w.SetMaster()

	topNav := ui.CreateNavBar()

	ui.SetupUI(w)
	ui.MakeTray(a)

	content := container.NewVBox(
		topNav,
		container.NewVBox(widget.NewLabel("Home")),
	)

	w.SetContent(content)
	w.ShowAndRun()
}

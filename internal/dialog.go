package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
)

func dialogScreen(win fyne.Window) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(layout.NewVBoxLayout())
}

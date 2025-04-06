package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
)

func containerScreen(_ fyne.Window) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(layout.NewVBoxLayout())
}

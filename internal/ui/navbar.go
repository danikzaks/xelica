package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateNavBar() *fyne.Container {
	homeButton := widget.NewButton("Home", func() {
		println("Home clicked")
	})

	settingsButton := widget.NewButton("Settings", func() {
		println("Settings clicked")
	})

	nav := container.NewHBox(
		homeButton,
		settingsButton,
	)

	return nav
}

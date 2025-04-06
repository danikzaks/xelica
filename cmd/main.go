package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()

	myWindow := myApp.NewWindow("Xelica")

	hi := widget.NewLabel("Hi")

	myWindow.SetContent(hi)
	myWindow.Resize(fyne.NewSize(400, 400))

	myWindow.ShowAndRun()
}

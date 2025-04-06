package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/danikzaks/xelica/internal/converter"
)

func SetupUI(window fyne.Window) {
	inputField := widget.NewMultiLineEntry()
	inputField.SetPlaceHolder("Введите JSON или число для конвертации...")

	outputField := widget.NewMultiLineEntry()
	outputField.SetPlaceHolder("Результат будет здесь...")

	formatButton := widget.NewButton("Форматировать JSON", func() {
		result, err := converter.FormatJSON(inputField.Text)
		if err != nil {
			outputField.SetText("Ошибка: " + err.Error())
		} else {
			outputField.SetText(result)
		}
	})

	convertButton := widget.NewButton("Конвертировать число", func() {
		result := converter.ConvertNumber(inputField.Text)
		outputField.SetText(result)
	})

	window.SetContent(container.NewVBox(
		inputField,
		formatButton,
		convertButton,
		outputField,
	))
}

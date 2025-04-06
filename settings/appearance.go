package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const (
	themeNameDark        = "dark"
	themeNameLight       = "light"
	themeNameSystem      = ""
	themeNameSystemLabel = "system default"
)

type Settings struct {
	fyneSettings app.SettingsSchema

	preview *fyne.Container
	colors  []fyne.CanvasObject

	userTheme fyne.Theme
}

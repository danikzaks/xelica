package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"log"
)

func MakeTray(a fyne.App) {
	if desk, ok := a.(desktop.App); ok {
		helloItem := fyne.NewMenuItem("Hello", func() {
			log.Println("Hello clicked!")
		})
		helloItem.Icon = theme.HomeIcon()

		openAppItem := fyne.NewMenuItem("Open App", func() {
			log.Println("Opening main window...")
		})

		settingsItem := fyne.NewMenuItem("Settings", func() {
			log.Println("Opening settings...")
		})

		separator := fyne.NewMenuItemSeparator()

		quitItem := fyne.NewMenuItem("Quit", func() {
			log.Println("Exiting application...")
			a.Quit()
		})

		menu := fyne.NewMenu("Xelica Tool", helloItem, openAppItem, settingsItem, separator, quitItem)

		desk.SetSystemTrayMenu(menu)

		trayIcon := theme.HomeIcon()
		desk.SetSystemTrayIcon(trayIcon)
	}
}

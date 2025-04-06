package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type iconInfo struct {
	name string
	icon fyne.Resource
}

type browser struct {
	current int
	icons   []iconInfo

	name *widget.Select
	icon *widget.Icon
}

func (b *browser) setIcon(index int) {
	if index < 0 || index > len(b.icons)-1 {
		return
	}
	b.current = index

	b.name.SetSelected(b.icons[index].name)
	b.icon.SetResource(b.icons[index].icon)
}

func iconScreen(_ fyne.Window) fyne.CanvasObject {
	return nil
}

package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func ContainerWithSize(w, h float32, objects ...fyne.CanvasObject) *fyne.Container {
	return container.NewHBox(container.NewGridWrap(fyne.NewSize(w, h), objects...))
}

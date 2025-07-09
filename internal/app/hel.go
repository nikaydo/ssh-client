package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func ContainerWithSize(w, h float32, objects ...fyne.CanvasObject) *fyne.Container {
	return container.New(layout.NewHBoxLayout(), container.NewGridWrap(fyne.NewSize(w, h), objects...))
}

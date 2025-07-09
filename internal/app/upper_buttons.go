package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MenuButtons struct {
	Connect  *widget.Button
	Settings *widget.Button
	Consoles *widget.Button
}

type UpMenu struct {
	MenuButtons
	ConnectE ConnectEntrys
}

func InitUpMenu() (b UpMenu) {
	b.Connect = widget.NewButton("Подключиться", func() {})
	b.Settings = widget.NewButton("Настройки", func() {})
	b.Consoles = widget.NewButton("Консоль", func() {})
	return
}

func (u *UpMenu) GetContainer() *fyne.Container {
	return container.New(
		layout.NewHBoxLayout(),
		ContainerWithSize(120, 36, u.Connect),
		ContainerWithSize(100, 36, u.Settings),
		ContainerWithSize(80, 36, u.Consoles),
	)
}

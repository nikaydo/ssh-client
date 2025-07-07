package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type UpButtons struct {
	Connect  *widget.Button
	ConnectE ConnectEntrys
	Settings *widget.Button
}

func InitUpButtons() (b UpButtons) {
	b.Connect = widget.NewButton("Подключиться", func() {})
	b.Settings = widget.NewButton("Настройки", func() {})
	return
}

func (u *UpButtons) GetContainer() *fyne.Container {
	return container.NewHBox(
		ContainerWithSize(120, 36, u.Connect),
		ContainerWithSize(100, 36, u.Settings),
	)
}

func (u *UpButtons) ConnectButton(f func()) {
	u.Connect.OnTapped = f
}

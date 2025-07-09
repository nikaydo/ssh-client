package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type ConnectionsField interface {
	FillStruct()
}

type ConnectPage struct {
	ConnectEntrys
	Container *fyne.Container
}

type ConnectEntrys struct {
	Login     *widget.Entry
	Password  *widget.Entry
	Ip        *widget.Entry
	LoginC    *fyne.Container
	PasswordC *fyne.Container
	IpC       *fyne.Container
	Cbutton   *widget.Button
}

func (c *ConnectEntrys) FillStruct(w, h float32) {
	c.Login = widget.NewEntry()
	c.Password = widget.NewEntry()
	c.Ip = widget.NewEntry()
	c.LoginC = ContainerWithSize(w, h, c.Login)
	c.PasswordC = ContainerWithSize(w, h, c.Password)
	c.IpC = ContainerWithSize(w, h, c.Ip)
}

func (c *ConnectEntrys) MakeContainer() *fyne.Container {
	c.Cbutton = widget.NewButton("Подключиться", nil)
	B := container.New(layout.NewHBoxLayout(), container.NewGridWrap(fyne.NewSize(200, 36), c.Cbutton))
	return container.New(
		layout.NewVBoxLayout(),
		c.LoginC,
		c.PasswordC,
		c.IpC,
		B,
	)
}

func SetLables(t []fyne.CanvasObject, names []string) *fyne.Container {
	for _, i := range names {
		t = append(t, widget.NewLabel(i))
	}
	objs := make([]fyne.CanvasObject, len(t))
	copy(objs, t)
	n := container.NewVBox(
		objs...,
	)
	return n
}

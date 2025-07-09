package app

import (
	myssh "github.com/nikaydo/ssh-client/internal/ssh"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"golang.org/x/crypto/ssh"
)

type App struct {
	App    fyne.App
	Window fyne.Window
}

func RunApp() App {
	return App{App: app.New()}
}

func (a *App) MakeWindow() {
	a.Window = a.App.NewWindow("ssh client")
	Tabs := InitTabs()
	UpButtons := InitUpMenu()
	UpButtons.ConnectE.FillStruct(200, 36)
	labelsConnect := SetLables([]fyne.CanvasObject{}, []string{"Логин", "Пороль", "ip"})
	form := container.NewBorder(
		UpButtons.GetContainer(),
		nil,
		container.New(layout.NewVBoxLayout(), container.New(layout.NewHBoxLayout(), labelsConnect, UpButtons.ConnectE.MakeContainer())),
		nil,
	)

	UpButtons.Connect.OnTapped = func() {
		UpButtons.Connect.Enable()
		a.Window.SetContent(form)
		UpButtons.Connect.Disable()
		UpButtons.Consoles.Enable()
	}

	UpButtons.Consoles.OnTapped = func() {
		UpButtons.Consoles.Enable()
		a.Window.SetContent(container.NewBorder(
			UpButtons.GetContainer(),
			nil,
			nil,
			nil,
			Tabs.DocTabs))
		UpButtons.Consoles.Disable()
		UpButtons.Connect.Enable()
	}

	UpButtons.Connect.Disable()

	Tabs.DocTabs.OnClosed = func(ti *container.TabItem) {
		if len(Tabs.DocTabs.Items) == 0 {
			a.Window.SetContent(form)
		}
	}

	UpButtons.ConnectE.Cbutton.OnTapped = func() {
		TabList := InitTabs()
		TabList.DocTabs = Tabs.DocTabs
		var session myssh.Session
		id, grid, scroll := TabList.Add("Подключение...", &session)
		TabList.AppendDocTab()
		s := myssh.SetConfig(UpButtons.ConnectE.Login.Text, UpButtons.ConnectE.Password.Text)
		UpButtons.Connect.Enable()
		UpButtons.Consoles.Disable()
		TabList.ShowAll()
		go func() {
			if err := s.Dial(UpButtons.ConnectE.Ip.Text); err != nil {
				fyne.Do(func() {
					TabList.Items[id].Tab.Text = "Ошибка"
					TabList.DocTabs.Refresh()
					grid.Append(err.Error())
				})
				return
			}
			session, _ = myssh.InitSession(s, session, ssh.TerminalModes{})

			if err := session.Shell(); err != nil {
				fyne.Do(func() {
					TabList.Items[id].Tab.Text = "Ошибка"
					TabList.DocTabs.Refresh()
					grid.Append(err.Error())
				})
				return
			}
			TabList.Items[id].Tab.Text = UpButtons.ConnectE.Ip.Text
			item := TabList.Items[id]
			item.Session = session
			TabList.Items[id] = item
			fyne.Do(func() {
				TabList.DocTabs.Refresh()
			})
			myssh.StartListening(session, grid, scroll)

		}()

		form := container.NewBorder(
			UpButtons.GetContainer(),
			nil,
			nil,
			nil,
			TabList.DocTabs)
		a.Window.SetContent(form)
		scroll.ScrollToBottom()
	}

	a.Window.SetContent(form)
	a.Window.Resize(fyne.NewSize(1000, 800))
	a.Window.ShowAndRun()
}

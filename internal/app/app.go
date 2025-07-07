package app

import (
	"bufio"
	"log"

	myssh "github.com/nikaydo/ssh-client/internal/ssh"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

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
	UpButtons := InitUpButtons()
	UpButtons.ConnectE.FillStruct(200, 36)

	command := widget.NewEntry()
	grid := widget.NewTextGrid()
	grid.Hide()
	command.Hide()

	scroll := container.NewScroll(container.NewStack(grid))

	labelsConnect := SetLables([]fyne.CanvasObject{}, []string{"Логин", "Пороль", "ip"})
	session_cmd := container.NewBorder(nil, command, nil, nil, scroll)
	form := container.NewBorder(
		UpButtons.GetContainer(),
		nil,
		container.NewVBox(container.NewHBox(labelsConnect, UpButtons.ConnectE.MakeContainer())),
		nil,
	)
	UpButtons.ConnectButton(func() {

		a.Window.SetContent(form)
	})

	a.Window.SetContent(form)

	Tabs.Tabs.OnClosed = func(ti *container.TabItem) {
		if len(Tabs.Tabs.Items) == 0 {
			a.Window.SetContent(form)
		}
	}

	var session myssh.Session
	UpButtons.ConnectE.Cbutton.OnTapped = func() {

		Tabs.Add(UpButtons.ConnectE.Ip.Text, session_cmd)

		s := myssh.SetConfig(UpButtons.ConnectE.Login.Text, UpButtons.ConnectE.Password.Text)

		if err := s.Dial(UpButtons.ConnectE.Ip.Text); err != nil {
			log.Println(err)
		}

		session, _ = myssh.InitSession(s, ssh.TerminalModes{})
		if err := session.Shell(); err != nil {
			log.Println(err)
		}

		go func() {
			scanner := bufio.NewScanner(session.StdoutPipe)
			for scanner.Scan() {
				line := scanner.Text()

				fyne.Do(func() {
					grid.Append(line)
					scroll.ScrollToBottom()
				})
			}
		}()

		grid.Show()
		command.Show()

		a.Window.SetContent(Tabs.Tabs)
		scroll.ScrollToBottom()
	}

	command.OnSubmitted = func(str string) {
		_, err := session.StdinPipe.Write([]byte(str + "\n"))
		if err != nil {
			log.Println(err)
		}
		command.SetText("")
		scroll.ScrollToBottom()
	}
	a.Window.SetContent(form)
	a.Window.Resize(fyne.NewSize(1000, 800))
	a.Window.ShowAndRun()
}

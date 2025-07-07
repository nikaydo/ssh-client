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
	var ConnectE ConnectEntrys
	ConnectE.FillStruct(200, 36)
	c := ConnectE.MakeContainer()

	command := widget.NewEntry()
	send := ContainerWithSize(400, 36, command)

	grid := widget.NewTextGrid()
	grid.Hide()
	command.Hide()
	scroll := container.NewScroll(container.NewStack(grid))

	labelsConnect := SetLables([]fyne.CanvasObject{}, []string{"Логин", "Пороль", "ip"})
	form := container.NewHBox(
		labelsConnect,
		c,
	)
	session_cmd := container.NewBorder(nil, send, nil, nil, scroll)
	var session myssh.Session
	ConnectE.Cbutton.OnTapped = func() {
		s := myssh.SetConfig(ConnectE.Login.Text, ConnectE.Password.Text)
		if err := s.Dial(ConnectE.Ip.Text); err != nil {
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
		a.Window.SetContent(session_cmd)
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

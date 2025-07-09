package ssh

import (
	"bufio"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/crypto/ssh"
)

type Ssh struct {
	Config *ssh.ClientConfig
	Client *ssh.Client
}

type Session struct {
	*ssh.Session
	StdinPipe  io.WriteCloser
	StdoutPipe io.Reader
	StderrPipe io.Reader
}

func (s *Ssh) Dial(host string) (err error) {
	client, err := ssh.Dial("tcp", host, s.Config)
	if err != nil {
		return
	}
	s.Client = client
	return
}

func InitSession(ssh_ Ssh, s Session, termmodes ssh.TerminalModes) (session Session, err error) {
	session = s
	session.Session, err = ssh_.Client.NewSession()
	if err != nil {
		return
	}
	session.StdinPipe, err = session.Session.StdinPipe()
	if err != nil {
		return
	}
	session.StderrPipe, err = session.Session.StderrPipe()
	if err != nil {
		return
	}
	session.StdoutPipe, err = session.Session.StdoutPipe()
	if err != nil {
		return
	}
	err = session.RequestPty("xterm", 40, 80, termmodes)
	if err != nil {
		return
	}
	return
}

func SetConfig(User, Password string) (s Ssh) {
	s.Config = &ssh.ClientConfig{
		User: User,
		Auth: []ssh.AuthMethod{
			ssh.Password(Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return
}

func StartListening(session Session, grid *widget.TextGrid, scroll *container.Scroll) {
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
}

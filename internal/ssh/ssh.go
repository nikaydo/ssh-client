package ssh

import (
	"io"

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

func InitSession(ssh_ Ssh, termmodes ssh.TerminalModes) (Session, error) {
	var session Session
	var err error
	session.Session, err = ssh_.Client.NewSession()
	if err != nil {
		return session, err
	}
	session.StdinPipe, err = session.Session.StdinPipe()
	if err != nil {
		return session, err
	}
	session.StderrPipe, err = session.Session.StderrPipe()
	if err != nil {
		return session, err
	}
	session.StdoutPipe, err = session.Session.StdoutPipe()
	if err != nil {
		return session, err
	}
	err = session.RequestPty("xterm", 40, 80, termmodes)
	if err != nil {
		return session, err
	}
	return session, nil
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

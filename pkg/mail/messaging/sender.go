package messaging

import (
	"fmt"
	"net/smtp"
	"wire_test/pkg/mail/conn"
)

type Sender struct {
	client *conn.SmtpClient
	auth   Auther
}

type Auther interface {
	Auth() (smtp.Auth, string, error)
}

func NewSender(client *conn.SmtpClient, auth Auther) *Sender {
	return &Sender{
		client: client,
		auth:   auth,
	}
}

func (s *Sender) Send(to, subject string, contents []byte) error {
	auth, sender, err := s.auth.Auth()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	err = s.client.Auth(auth)
	if err != nil {
		return err
	}
	err = s.client.Mail(sender)
	if err != nil {
		return err
	}
	err = s.client.Rcpt(to)
	if err != nil {
		return err
	}

	w, err := s.client.Data()
	if err != nil {
		return err
	}
	subjectHeader := []byte(fmt.Sprintf("Subject: %s\n", subject))
	wholeEmail := append(subjectHeader, contents...)
	_, err = w.Write(wholeEmail)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return s.client.Quit()
}

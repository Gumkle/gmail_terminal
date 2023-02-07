package auth

import (
	"fmt"
	"net/smtp"
)

type Auther interface {
	Auth(auth smtp.Auth) error
}

type addresser interface {
	Host() string
}

type AutherAddresser interface {
	Auther
	addresser
}

type storer interface {
	Store(login, password string) error
}

type purger interface {
	Purge() error
}

type reader interface {
	Read() (string, string, error)
}

type StorerPurgerReader interface {
	storer
	purger
	reader
}

type Authenticator struct {
	client     AutherAddresser
	repository StorerPurgerReader
}

func NewAuthenticator(client AutherAddresser, repository StorerPurgerReader) *Authenticator {
	return &Authenticator{
		client:     client,
		repository: repository,
	}
}

func (a *Authenticator) Login(login, password string) error {
	auth := smtp.PlainAuth("", login, password, "smtp.gmail.com")
	err := a.client.Auth(auth)
	if err != nil {
		return err
	}

	// if no error has occured, that means the login and passwords are valid
	// let's now save them for further use
	err = a.repository.Store(login, password)
	if err != nil {
		return fmt.Errorf("failed to log in: %w", err)
	}
	return nil
}

func (a *Authenticator) Logout() error {
	err := a.repository.Purge()
	if err != nil {
		return fmt.Errorf("failed to clear credentials: %w", err)
	}
	return nil
}

func (a *Authenticator) Auth() (smtp.Auth, string, error) {
	login, password, err := a.repository.Read()
	if err != nil {
		return nil, "", fmt.Errorf("auth failed: %w", err)
	}
	return smtp.PlainAuth("", login, password, "smtp.gmail.com"), login, nil
}

package auth

import (
	"fmt"
	"net/smtp"
)

type Auther interface {
	Auth(auth smtp.Auth) error
}

type Addresser interface {
	Address() string
}

type AutherAddresser interface {
	Auther
	Addresser
}

type Storer interface {
	Store(login, password string) error
}

type Purger interface {
	Purge() error
}

type StorerPurger interface {
	Storer
	Purger
}

type Authenticator struct {
	client     AutherAddresser
	repository StorerPurger
}

func NewAuthenticator(client AutherAddresser, repository StorerPurger) *Authenticator {
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

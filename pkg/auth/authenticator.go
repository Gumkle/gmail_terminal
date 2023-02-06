package auth

import (
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

type Authenticator struct {
	client     AutherAddresser
	repository Storer
}

func NewAuthenticator(client AutherAddresser, repository Storer) *Authenticator {
	return &Authenticator{
		client:     client,
		repository: repository,
	}
}

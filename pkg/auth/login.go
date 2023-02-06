package auth

import (
	"fmt"
	"net/smtp"
)

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

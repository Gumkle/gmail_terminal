// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"wire_test/pkg/filesystem"
	"wire_test/pkg/mail/auth"
	"wire_test/pkg/mail/conn"
	"wire_test/pkg/mail/messaging"
	"wire_test/pkg/storage"
)

// Injectors from wire.go:

func InitializeSmtpClient(address conn.SmtpAddress) (*conn.SmtpClient, func(), error) {
	smtpClient, cleanup, err := conn.NewSmtpClient(address)
	if err != nil {
		return nil, nil, err
	}
	return smtpClient, func() {
		cleanup()
	}, nil
}

func InitializeAuthenticator(address conn.SmtpAddress, appDataRoot filesystem.PathString) (*auth.Authenticator, func(), error) {
	smtpClient, cleanup, err := InitializeSmtpClient(address)
	if err != nil {
		return nil, nil, err
	}
	fileSystem, err := filesystem.NewFileSystem(appDataRoot)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	authRepository := storage.NewAuthRepository(fileSystem)
	authenticator := auth.NewAuthenticator(smtpClient, authRepository)
	return authenticator, func() {
		cleanup()
	}, nil
}

func InitializeSender(address conn.SmtpAddress, appDataRoot filesystem.PathString) (*messaging.Sender, func(), error) {
	smtpClient, cleanup, err := InitializeSmtpClient(address)
	if err != nil {
		return nil, nil, err
	}
	authenticator, cleanup2, err := InitializeAuthenticator(address, appDataRoot)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	sender := messaging.NewSender(smtpClient, authenticator)
	return sender, func() {
		cleanup2()
		cleanup()
	}, nil
}

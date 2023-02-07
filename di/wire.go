//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"wire_test/pkg/filesystem"
	"wire_test/pkg/mail/auth"
	"wire_test/pkg/mail/conn"
	"wire_test/pkg/mail/messaging"
	"wire_test/pkg/storage"
)

func InitializeSmtpClient(address conn.SmtpAddress) (*conn.SmtpClient, func(), error) {
	wire.Build(conn.NewSmtpClient)
	return &conn.SmtpClient{}, nil, nil
}

func InitializeAuthenticator(address conn.SmtpAddress, appDataRoot filesystem.PathString) (*auth.Authenticator, func(), error) {
	wire.Build(
		auth.NewAuthenticator,
		wire.Bind(new(auth.AutherAddresser), new(*conn.SmtpClient)),
		InitializeSmtpClient,
		wire.Bind(new(auth.StorerPurgerReader), new(*storage.AuthRepository)),
		storage.NewAuthRepository,
		wire.Bind(new(storage.SaverRemoverReader), new(*filesystem.FileSystem)),
		filesystem.NewFileSystem,
	)
	return &auth.Authenticator{}, nil, nil
}

func InitializeSender(address conn.SmtpAddress, appDataRoot filesystem.PathString) (*messaging.Sender, func(), error) {
	wire.Build(
		wire.Bind(new(messaging.Auther), new(*auth.Authenticator)),
		messaging.NewSender,
		InitializeAuthenticator,
		InitializeSmtpClient,
	)
	return &messaging.Sender{}, nil, nil
}

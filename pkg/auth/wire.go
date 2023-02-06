//go:build wireinject
// +build wireinject

package auth

import (
	"github.com/google/wire"
	"wire_test/pkg/auth/internal"
	"wire_test/pkg/filesystem"
	"wire_test/pkg/storage"
)

func InitializeSmtpClient(address internal.SmtpAddress) (*internal.SmtpClient, error) {
	wire.Build(internal.NewSmtpClient)
	return &internal.SmtpClient{}, nil
}

func InitializeAuthenticator(address internal.SmtpAddress, appDataRoot filesystem.PathString) (*Authenticator, error) {
	wire.Build(
		NewAuthenticator,
		wire.Bind(new(AutherAddresser), new(*internal.SmtpClient)),
		InitializeSmtpClient,
		wire.Bind(new(Storer), new(*storage.AuthRepository)),
		storage.NewAuthRepository,
		wire.Bind(new(storage.Saver), new(*filesystem.FileSystem)),
		filesystem.NewFileSystem,
	)
	return &Authenticator{}, nil
}

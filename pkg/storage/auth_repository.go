package storage

import (
	"fmt"
	"strings"
)

type AuthRepository struct {
	Database SaverRemoverReader
}

type saver interface {
	Save(destination, contents string) error
}

type remover interface {
	Remove(destination string) error
}

type reader interface {
	Read(source string) (string, error)
}

type SaverRemoverReader interface {
	saver
	remover
	reader
}

func NewAuthRepository(saver SaverRemoverReader) *AuthRepository {
	return &AuthRepository{
		Database: saver,
	}
}

func (a *AuthRepository) Store(login, password string) error {
	err := a.Database.Save("credentials", fmt.Sprintf("%s:%s", login, password))
	if err != nil {
		return fmt.Errorf("failed to store credentials: %w", err)
	}
	return nil
}

func (a *AuthRepository) Purge() error {
	err := a.Database.Remove("credentials")
	if err != nil {
		return fmt.Errorf("failed to purge credentials: %w", err)
	}
	return nil
}

func (a *AuthRepository) Read() (string, string, error) {
	contents, err := a.Database.Read("credentials")
	if err != nil {
		return "", "", fmt.Errorf("failed to read credentials: %v", err)
	}
	splitted := strings.Split(contents, ":")
	return splitted[0], splitted[1], nil
}

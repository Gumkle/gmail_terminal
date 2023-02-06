package storage

import "fmt"

type AuthRepository struct {
	Database SaverRemover
}

type Saver interface {
	Save(destination, contents string) error
}

type Remover interface {
	Remove(destination string) error
}

type SaverRemover interface {
	Saver
	Remover
}

func NewAuthRepository(saver SaverRemover) *AuthRepository {
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

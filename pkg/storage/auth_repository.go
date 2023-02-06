package storage

import "fmt"

type AuthRepository struct {
	Database Saver
}

type Saver interface {
	Save(destination, contents string) error
}

func NewAuthRepository(saver Saver) *AuthRepository {
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

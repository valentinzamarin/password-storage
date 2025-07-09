package services

import (
	"password-storage/internal/domain/entities"
	"password-storage/internal/domain/repositories"
)

type PasswordService struct {
	passwordRepo repositories.PasswordRepo
}

func NewPasswordService(passwordRepo repositories.PasswordRepo) *PasswordService {
	return &PasswordService{passwordRepo: passwordRepo}
}

func (ps *PasswordService) AddNewPassword(url, login, password, description string) error {
	newPassword, err := entities.NewPassword(url, login, password, description)
	if err != nil {
		return err
	}

	if err := newPassword.Validate(); err != nil {
		return err
	}

	return ps.passwordRepo.AddPassword(*newPassword)
}

func (ps *PasswordService) GetPasswords() ([]*entities.Password, error) {
	passwords, err := ps.passwordRepo.GetAllPasswords()
	if err != nil {
		return nil, err
	}

	return passwords, nil
}

func (ps *PasswordService) DeletePassword(id int) {
	ps.passwordRepo.DeletePasswordById(uint(id))
}

package repositories

import "password-storage/internal/domain/entities"

type PasswordRepo interface {
	AddPassword(password entities.Password) error
	GetAllPasswords() ([]*entities.Password, error)
	DeletePasswordById(id uint) error
}

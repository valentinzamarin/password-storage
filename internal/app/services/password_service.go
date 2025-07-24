package services

import (
	"password-storage/internal/app/events"
	"password-storage/internal/domain/entities"
	domainevents "password-storage/internal/domain/events"
	"password-storage/internal/domain/repositories"
)

type PasswordService struct {
	passwordRepo repositories.PasswordRepo
	eventBus     *events.EventBus
}

func NewPasswordService(passwordRepo repositories.PasswordRepo, eventBus *events.EventBus) *PasswordService {
	return &PasswordService{
		passwordRepo: passwordRepo,
		eventBus:     eventBus,
	}
}

func (ps *PasswordService) AddNewPassword(url, login, password, description string) error {
	newPassword, err := entities.NewPassword(url, login, password, description)
	if err != nil {
		return err
	}

	if err := newPassword.Validate(); err != nil {
		return err
	}

	ps.passwordRepo.AddPassword(*newPassword)

	ps.eventBus.Publish(domainevents.PasswordTopic, domainevents.AddedPasswordEvent{
		URL:         newPassword.URL,
		Login:       newPassword.Login,
		Password:    newPassword.Password,
		Description: newPassword.Description,
	})

	return nil
}

func (ps *PasswordService) GetPasswords() ([]*entities.Password, error) {
	passwords, err := ps.passwordRepo.GetAllPasswords()
	if err != nil {
		return nil, err
	}

	return passwords, nil
}

func (ps *PasswordService) DeletePassword(id uint) {

	ps.passwordRepo.DeletePasswordById(id)

	ps.eventBus.Publish(domainevents.PasswordTopic, domainevents.RemovedPasswordEvent{
		ID: id,
	})
}

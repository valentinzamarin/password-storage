package sqlite

import "password-storage/internal/domain/entities"

func toDBPassword(password entities.Password) *PasswordModel {
	return &PasswordModel{
		ID:          uint(password.ID),
		URL:         password.URL,
		Login:       password.Login,
		Password:    password.Password,
		Description: password.Description,
	}
}

func fromDBPassword(model *PasswordModel) *entities.Password {
	return &entities.Password{
		ID:          int(model.ID),
		URL:         model.URL,
		Login:       model.Login,
		Password:    model.Password,
		Description: model.Description,
	}
}

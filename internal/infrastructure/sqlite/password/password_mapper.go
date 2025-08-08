package passwords

import "password-storage/internal/domain/entities"

func toDBPassword(password entities.Password, encryptedPassword []byte) *PasswordModel {

	return &PasswordModel{
		ID:                uint(password.ID),
		URL:               password.URL,
		Login:             password.Login,
		EncryptedPassword: encryptedPassword,
		Description:       password.Description,
	}
}

func fromDBPassword(model *PasswordModel, decryptedPasswordBytes string) *entities.Password {
	return &entities.Password{
		ID:          model.ID,
		URL:         model.URL,
		Login:       model.Login,
		Password:    decryptedPasswordBytes,
		Description: model.Description,
	}
}

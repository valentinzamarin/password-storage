package passwords

import (
	"fmt"
	"password-storage/internal/app/encrypt"
	"password-storage/internal/domain/entities"
	"password-storage/internal/domain/repositories"

	"gorm.io/gorm"
)

type GormPasswordRepository struct {
	db             *gorm.DB
	encryptService *encrypt.PasswordEncrypt
}

func NewGormPasswordRepository(db *gorm.DB, encryptService *encrypt.PasswordEncrypt) repositories.PasswordRepo {
	return &GormPasswordRepository{
		db:             db,
		encryptService: encryptService,
	}
}

func (p *GormPasswordRepository) AddPassword(password entities.Password) error {
	encryptedPassword, err := p.encryptService.Encrypt([]byte(password.Password))
	if err != nil {
		fmt.Printf("Encryption failed: %v\n", err)
		return err
	}

	dbPassword := toDBPassword(password, encryptedPassword)

	err = p.db.Create(dbPassword).Error
	if err != nil {
		fmt.Printf("DB save failed: %v\n", err)
		return err
	}
	return nil
}

func (p *GormPasswordRepository) GetAllPasswords() ([]*entities.Password, error) {
	var dbPasswords []PasswordModel
	if err := p.db.Find(&dbPasswords).Error; err != nil {
		return nil, err
	}

	passwords := make([]*entities.Password, len(dbPasswords))

	for i, dbPassword := range dbPasswords {

		decryptedPasswordBytes, err := p.encryptService.Decrypt(dbPassword.EncryptedPassword)
		if err != nil {
			return nil, fmt.Errorf("could not decrypt password for entry %d: %w", dbPassword.ID, err)
		}
		passwords[i] = fromDBPassword(&dbPassword, string(decryptedPasswordBytes))
	}

	return passwords, nil
}

func (p *GormPasswordRepository) DeletePasswordById(id uint) error {
	result := p.db.Delete(&PasswordModel{}, id)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *GormPasswordRepository) UpdatePassword(id uint, description string) error {
	result := p.db.Model(&PasswordModel{}).Where("id = ?", id).Update("description", description)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *GormPasswordRepository) SearchPasswordByURL(title string) []*entities.Password {
	var passwords []*entities.Password
	p.db.Where("url LIKE ?", "%"+title+"%").Find(&passwords)
	return passwords
}

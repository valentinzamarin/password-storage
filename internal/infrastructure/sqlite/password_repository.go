package sqlite

import (
	"password-storage/internal/domain/entities"
	"password-storage/internal/domain/repositories"

	"gorm.io/gorm"
)

type GormPasswordRepository struct {
	db *gorm.DB
}

func NewGormPasswordRepository(db *gorm.DB) repositories.PasswordRepo {
	return &GormPasswordRepository{db: db}
}

func (p *GormPasswordRepository) AddPassword(password entities.Password) error {
	dbPassword := toDBPassword(password)

	if err := p.db.Create(dbPassword).Error; err != nil {
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
		passwords[i] = fromDBPassword(&dbPassword)
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

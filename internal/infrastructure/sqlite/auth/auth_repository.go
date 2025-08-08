package auth

import "gorm.io/gorm"

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) GetMasterAuth() (*AuthModel, error) {
	var masterAuth AuthModel

	result := r.db.First(&masterAuth)

	if result.Error != nil {
		return nil, result.Error
	}

	return &masterAuth, nil
}

func (r *AuthRepo) CreateMasterAuth(masterAuth *AuthModel) error {

	result := r.db.Create(masterAuth)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

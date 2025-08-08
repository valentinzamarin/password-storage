package services

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"password-storage/internal/app/encrypt"
	"password-storage/internal/infrastructure/sqlite/auth"

	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
)

type AuthService struct {
	authRepo   *auth.AuthRepo
	encryptSvc *encrypt.PasswordEncrypt
}

func NewAuthService(authRepo *auth.AuthRepo, encryptSvc *encrypt.PasswordEncrypt) *AuthService {
	return &AuthService{
		authRepo:   authRepo,
		encryptSvc: encryptSvc,
	}
}

func (s *AuthService) IsMasterPasswordSet() (bool, error) {
	auth, err := s.authRepo.GetMasterAuth()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return auth != nil, nil
}

func (s *AuthService) CreateMasterPassword(password string) error {
	salt, err := s.encryptSvc.GenerateSalt()
	if err != nil {
		return fmt.Errorf("could not generate salt: %w", err)
	}

	verificationHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	authData := &auth.AuthModel{
		Salt:             salt,
		VerificationHash: verificationHash,
	}
	if err := s.authRepo.CreateMasterAuth(authData); err != nil {
		return fmt.Errorf("could not save master auth data: %w", err)
	}

	s.encryptSvc.DeriveKeyFromPassword(password, salt)
	return nil
}

func (s *AuthService) Authenticate(password string) error {

	authData, err := s.authRepo.GetMasterAuth()
	if err != nil {
		return fmt.Errorf("could not get master auth data: %w", err)
	}

	hashToCompare := argon2.IDKey([]byte(password), authData.Salt, 1, 64*1024, 4, 32)

	if subtle.ConstantTimeCompare(authData.VerificationHash, hashToCompare) != 1 {
		return errors.New("invalid master password")
	}

	s.encryptSvc.DeriveKeyFromPassword(password, authData.Salt)
	return nil
}

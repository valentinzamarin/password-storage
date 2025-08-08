package encrypt

import (
	"crypto/rand"
	"errors"
	"io"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/nacl/secretbox"
)

const (
	keySize  = 32
	saltSize = 16
)

type PasswordEncrypt struct {
	encryptionKey *[keySize]byte
}

func NewPasswordEncrypt() *PasswordEncrypt {
	return &PasswordEncrypt{}
}

func (s *PasswordEncrypt) Encrypt(data []byte) ([]byte, error) {
	if s.encryptionKey == nil {
		return nil, errors.New("encryption key is not set")
	}

	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return nil, err
	}

	encrypted := secretbox.Seal(nonce[:], data, &nonce, s.encryptionKey)
	return encrypted, nil
}

func (s *PasswordEncrypt) Decrypt(encryptedData []byte) ([]byte, error) {
	if s.encryptionKey == nil {
		return nil, errors.New("encryption key is not set")
	}

	var nonce [24]byte
	copy(nonce[:], encryptedData[:24])

	decrypted, ok := secretbox.Open(nil, encryptedData[24:], &nonce, s.encryptionKey)
	if !ok {
		return nil, errors.New("decryption failed: invalid master password or corrupted data")
	}
	return decrypted, nil
}

func (s *PasswordEncrypt) DeriveKeyFromPassword(masterPassword string, salt []byte) {
	key := argon2.IDKey([]byte(masterPassword), salt, 1, 64*1024, 4, keySize)
	s.encryptionKey = &[keySize]byte{}
	copy(s.encryptionKey[:], key)
}

func (s *PasswordEncrypt) IsKeySet() bool {
	return s.encryptionKey != nil
}

func (s *PasswordEncrypt) GenerateSalt() ([]byte, error) {
	salt := make([]byte, saltSize)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}
	return salt, nil
}

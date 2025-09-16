package encrypt

import (
	"testing"
)

func TestPasswordEncrypt_EncryptDecrypt(t *testing.T) {
	encryptor := NewPasswordEncrypt()

	salt, err := encryptor.GenerateSalt()
	if err != nil {
		t.Fatalf("GenerateSalt failed: %v", err)
	}

	encryptor.DeriveKeyFromPassword("myMasterPassword", salt)

	if !encryptor.IsKeySet() {
		t.Fatal("expected key to be set after DeriveKeyFromPassword")
	}

	plainText := []byte("secret data to encrypt")

	encrypted, err := encryptor.Encrypt(plainText)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	decrypted, err := encryptor.Decrypt(encrypted)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	if string(decrypted) != string(plainText) {
		t.Errorf("decrypted data does not match original: got %s, want %s", decrypted, plainText)
	}
}

func TestPasswordEncrypt_EncryptWithoutKey(t *testing.T) {
	encryptor := NewPasswordEncrypt()

	_, err := encryptor.Encrypt([]byte("test"))
	if err == nil {
		t.Fatal("expected error when encryption key is not set")
	}
}

func TestPasswordEncrypt_DecryptWithoutKey(t *testing.T) {
	encryptor := NewPasswordEncrypt()

	_, err := encryptor.Decrypt([]byte("ciphertext"))
	if err == nil {
		t.Fatal("expected error when encryption key is not set")
	}
}

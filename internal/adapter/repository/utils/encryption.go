package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/madsilver/task-manager/internal/infrastructure/env"
	"io"
)

// Encrypt encrypts the given text using AES encryption with the provided key.
// text: the plaintext to encrypt
// Returns the encrypted text as a base64 encoded string, or an error if encryption fails.
func Encrypt(text string) (string, error) {
	enc, _ := base64.StdEncoding.DecodeString(env.App.TaskEncryptionKey)
	block, err := aes.NewCipher(enc)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(text), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the given base64 encoded encrypted text using AES decryption with the provided key.
// encrypted: the base64 encoded encrypted text to decrypt
// Returns the decrypted plaintext, or an error if decryption fails.
func Decrypt(encrypted string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	enc, _ := base64.StdEncoding.DecodeString(env.App.TaskEncryptionKey)
	block, err := aes.NewCipher(enc)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(data) < aesGCM.NonceSize() {
		return "", errors.New("invalid encrypted data")
	}

	nonce := data[:aesGCM.NonceSize()]
	ciphertext := data[aesGCM.NonceSize():]

	plain, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plain), nil
}

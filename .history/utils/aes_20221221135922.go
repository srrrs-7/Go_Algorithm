package utils

import (
	"crypto/aes"
	"crypto/rand"
)

func Aes() {}

func encryptAes(key, text []byte) ([]byte, error) {}

func DecryptAes(key, data []byte) ([]byte, error) {
	cb, err := aes.NewCipher(key)
	if err != nil {
	}
}

func generateKey() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

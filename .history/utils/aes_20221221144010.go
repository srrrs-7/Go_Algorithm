package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func Aes() {}

func encryptAes(key, text []byte) ([]byte, error) {
	cb, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(cb)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}
	cipher := gcm.Seal(nonce, nonce, text, nil)
	return cipher, nil
}

func DecryptAes(key, data []byte) ([]byte, error) {
	cb, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(cb)
	if err != nil {
		return nil, err
	}
	nonce, cipher := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	text, err := gcm.Open(nil, nonce, cipher, nil)
}

func generateKey() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

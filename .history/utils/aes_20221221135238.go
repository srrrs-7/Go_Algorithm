package utils

import "crypto/rand"

func Aes() {}

func aesEncrypt(key, text []byte) ([]byte, error) {}

func aesDecrypt(key, data []byte) ([]byte, error) {}

func generateKey() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
}

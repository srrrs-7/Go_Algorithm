package utils

import (
	"crypto/sha256"
	"fmt"
)

func Cipher() {
	res := HashPassword("srrrs4510")
	fmt.Println(res)
}

func HashPassword(password string) string {
	sha := sha256.New()
	sha.Write([]byte(password))
	hash := sha.Sum(nil)

	fmt.Println(string(hash))

	return
}
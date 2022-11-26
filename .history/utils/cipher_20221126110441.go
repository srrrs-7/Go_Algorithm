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
	hash := sha.Sum([]byte(password)

	var s string
	for _, x := range hash {
		s += string(x)
	}
	return s
}
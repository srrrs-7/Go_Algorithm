package utils

import (
	"crypto/sha256"
	"fmt"
)

func Cipher() {}

func HashPassword(password string) string {
	sha := sha256.New()
	hash := sha.Sum([]byte(password))
	fmt.Println(hash)

	s := ""
	for _, x := range hash {
		s += string(x)
	}
	fmt.Println(s)
	return s
}
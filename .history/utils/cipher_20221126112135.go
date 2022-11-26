package utils

import (
	"crypto/sha256"
	"fmt"
)

func Cipher() {
	res1 := HashPassword("srrrs4510")
	fmt.Println(res1)

	res2 := HashPassword("srrrs")
	fmt.Println(res2)

	err := VerifyHash(res1, res2)
	fmt.Println(err)
}

// SHA256
func HashPassword(password string) string {
	sha := sha256.New()
	sha.Write([]byte(password))
	hash := sha.Sum(nil)

	s := fmt.Sprintf("%x \n", hash)

	return s
}

// verify
func VerifyHash(password, hashPassword string) error {
	if password != hashPassword {
		return fmt.Errorf("verify password and hashPassword error")
	}
	return nil
}
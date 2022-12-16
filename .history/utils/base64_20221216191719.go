package utils

import (
	"encoding/base64"
	"fmt"
)

func ConvertBase64() {
	s := []byte("I am engineer")
}

func convertBase64(s []byte) string {
	b := base64.URLEncoding.EncodeToString(s)
	fmt.Println(b)
	return b
}

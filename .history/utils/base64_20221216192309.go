package utils

import (
	"encoding/base64"
	"fmt"
)

func ConvertBase64() {
	s := []byte("I am engineer")

	res := convertBase64(s)
	base64.URLEncoding.DecodeString(res)
}

func convertBase64(s []byte) string {
	b := base64.URLEncoding.EncodeToString(s)
	fmt.Println(b)
	return b
}

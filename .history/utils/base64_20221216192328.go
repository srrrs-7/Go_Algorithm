package utils

import (
	"encoding/base64"
	"fmt"
)

func ConvertBase64() {
	s := []byte("I am engineer")

	res := convertBase64(s)
	base64.URLEncoding.DecodeString(res)
	fmt.Println(res)
}

func convertBase64(s []byte) string {
	b := base64.URLEncoding.EncodeToString(s)
	return b
}

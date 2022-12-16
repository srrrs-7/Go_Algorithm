package utils

import (
	"encoding/base64"
	"fmt"
)

func ConvertBase64() {
	s := []byte("I am engineer")

	res := convertBase64(s)
	fmt.Println(res)
	base64.URLEncoding.DecodeString(res)

}

func convertBase64(s []byte) string {
	b := base64.URLEncoding.EncodeToString(s)
	return b
}

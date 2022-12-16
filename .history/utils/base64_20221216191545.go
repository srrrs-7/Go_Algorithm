package utils

import (
	"encoding/base64"
	"fmt"
)

func ConvertBase64() {
	s := []byte("I am engineer")
}

func convertBase64(s string) string {
	b := base64.URLEncoding.EncodeToString(s)
	fmt.Println(b)
}

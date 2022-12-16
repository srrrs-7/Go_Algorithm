package utils

import (
	"encoding/base64"
	"fmt"
	"log"
)

func ConvertBase64() {
	s := []byte("I am engineer")

	res := convertBase64(s)
	fmt.Println("base64: ", res)
	b, err := base64.URLEncoding.DecodeString(res)
	if err != nil {
		log.Println("base64 decode error: ", err)
	}
	fmt.Println(b)

}

func convertBase64(s []byte) string {
	b := base64.URLEncoding.EncodeToString(s)
	return b
}

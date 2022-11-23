package utils

import (
	"fmt"
	"os"
)

func Io() {
	w := os.Stdout
	fmt.Println(w)
}
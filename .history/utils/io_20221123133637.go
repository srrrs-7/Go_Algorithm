package utils

import (
	"fmt"
	"io"
	"os"
)

var w io.Writer

func Io() {
	w = os.Stdout
	fmt.Println(w)
}
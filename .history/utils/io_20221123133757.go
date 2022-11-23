package utils

import (
	"fmt"
	"os"
)

var w io.Writer

func Io() {
	w = os.Stdout
	fmt.Println(w)
}

func io() {}

func flag() {}
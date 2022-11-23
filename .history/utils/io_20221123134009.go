package utils

import (
	"fmt"
	"os"
)

func Io() {
	io()
	flag()
}

func io() {
	w = os.Stdout
	fmt.Println(w)
}

func flag() {
	var x, y int

	flags := flag("")
}
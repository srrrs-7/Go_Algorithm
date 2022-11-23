package utils

import (
	"fmt"
	"os"
)

var w io.Writer

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

	flags := flag.NewFlagSet("")
}
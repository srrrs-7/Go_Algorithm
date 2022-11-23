package utils

import (
	"flag"
	"fmt"
	"os"
)

func Io() {
	io()
	flag()
}

func setIo() {
	w := os.Stdout
	fmt.Println(w)
}

func srtFlag() {
	var x, y int

	flags := flag.NewFlagSet("args", flag.ContinueOnError)
}
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
	flags.IntVar(&x, "x", 0, "input x number")
	flags.IntVar(&y, "y", 0, "input y number")
}
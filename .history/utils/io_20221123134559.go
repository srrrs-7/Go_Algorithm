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
	args := os.Args[1:]

	var x, y int

	flags := flag.NewFlagSet("args", flag.ContinueOnError)
	flags.IntVar(&x, "x", 0, "input x number")
	flags.IntVar(&y, "y", 0, "input y number")
	flags.Parse(args)

	fmt.Printf("x: %d, y: %d \n", x, y)
}
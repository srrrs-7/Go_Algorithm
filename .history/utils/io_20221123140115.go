package utils

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func Io() {
	setIo()
	setFlag()
	setScan()
}

func setIo() {
	w := os.Stdout
	fmt.Println(w)
}


const args = 2
func setScan() {
	fmt.Printf("input string %d: ", args)
	var s [args]string
	for i := 0; i < args; i++ {
		fmt.Scan(&s[i])
	}
	if len(s) > args {
		log.Fatalf("input value limit %d", args)
	}
	fmt.Printf("output %s: ", s)
}

// go run main.go -x 2 -y 3
func setFlag() {
	args := os.Args[1:]

	var x, y int

	flags := flag.NewFlagSet("args", flag.ContinueOnError)
	flags.IntVar(&x, "x", 0, "input x number")
	flags.IntVar(&y, "y", 0, "input y number")
	flags.Parse(args)

	fmt.Printf("x: %d, y: %d \n", x, y)
}
package utils

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func Io() {
	setIo()
	setFlag()
	setScan()
}

func setBufScan() {
	file, err := os.Open("/text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		fmr.Println(scanner.Err())
	}
}


const args = 2
func setScan() {
	fmt.Printf("input string limit %d value: ", args)

	var s [args]string
	for i := 0; i < args; i++ {
		fmt.Scan(&s[i])
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
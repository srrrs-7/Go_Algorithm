package utils

import (
	"fmt"
	"os"
)

var w io.Writer

func Io() {

}

func io() {
	w = os.Stdout
	fmt.Println(w)
}

func flag() {

}
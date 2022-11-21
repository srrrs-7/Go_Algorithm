package utils

import (
	"fmt"
	"strings"
)

func StringContains() {
	s1 := "abc"
	s2 := "a"

	res := stringContains(s1, s2)
	fmt.Println("contain?: ", res)

	res2 := searchIndex(s1, s2)
	fmt.Println("index", res2)
}

func stringContains(s1, s2 string) bool {
	b := strings.Contains(s1, s2)
	return b
}

func searchIndex(s1, s2 string) bool {

}
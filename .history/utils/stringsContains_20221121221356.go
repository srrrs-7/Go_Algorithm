package utils

import (
	"fmt"
	"strings"
)

func StringContains() {
	s1 := "abc"
	s2 := "a"

	res := stringContains(s1, s2)
	fmt.Println(res)
}

func stringContains(s1, s2 string) bool {
	b := strings.Contains(s1, s2)
	return b
}
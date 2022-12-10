package leetcode

import (
	"fmt"
	"regexp"
)

func MatchStr() {
	s := "abba"
	p := ".b*a"

	res := isMatch(s, p)
	fmt.Println("match string", res)
}

func isMatch(s string, p string) bool {
	r := regexp.MustCompile(p).MatchString(s)

	return r
}

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
	r3 := ""

	r := regexp.MustCompile(p)
	r2 := r.FindAllString(s, 1)

	for _, v := range r2 {
		r3 += v
	}

	return len(s) == len(r3)
}

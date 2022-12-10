package leetcode

import (
	"fmt"
	"regexp"
)

func MatchStr() {
	s := "absfsadasda"
	p := ".b*a"

	res := isMatch(s, p)
	fmt.Println("match string", res)
}

func isMatch(s string, p string) bool {
	re := regexp.MustCompile(fmt.Sprintf("^%s$", p))
	return re.Match([]byte(s))
}

package leetcode

import "fmt"

func MatchStr() {
	s := "abc"
	p := ".b*"

	res := isMatch(s, p)
	fmt.Println(res)
}

func isMatch(s string, p string) bool {
	for i := range s {
		s[i] == p[i]
	}
	return true
}

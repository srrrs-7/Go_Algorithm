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
		if s[i] != p[i] || p[i] != '.' {
			return false
		}
		if p[i] == * {

		}
	}
	return true
}

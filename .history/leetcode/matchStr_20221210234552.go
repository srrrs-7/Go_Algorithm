package leetcode

import "fmt"

func MatchStr() {
	s := "abcb"
	p := ".b*"

	res := isMatch(s, p)
	fmt.Println("match string", res)
}

func isMatch(s string, p string) bool {
	if len(s) > len(p) && p[len(p)-1] != '*' {
		return false
	}
	for i := range p {
		if s[i] == p[i] || p[i] == '.' || p[i] == '*' {
			continue
		} else {
			return false
		}
	}
	return true
}

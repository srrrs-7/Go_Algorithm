package leetcode

import "fmt"

func MatchStr() {
	s := "absfsadasda"
	p := ".b*a"

	res := isMatch(s, p)
	fmt.Println("match string", res)
}

func isMatch(s string, p string) bool {
	if len(s) > len(p) {
		if p[len(p)-1] != '*' || s[len(s)-1] != p[len(p)-1] {
			return false
		}
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

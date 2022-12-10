package leetcode

import "fmt"

func MatchStr() {
	s := "abc"
	p := ".b*"

	res := isMatch(s, p)
	fmt.Println(res)
}

func isMatch(s string, p string) bool {
	for _, a := range s {
		for _, b := range p {
			switch {
			case b == a:
				continue
			case b == ''.':
				continue
			}
		}
	}
}

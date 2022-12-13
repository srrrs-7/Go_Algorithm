package leetcode

import "fmt"

func LongestPrefix() {
	s := "level"

	res := longestPrefix(s)
	fmt.Println("longest prefix and suffix", res)
}

func longestPrefix(s string) string {
	res := ""
	for i, j := 0, len(s); i < j; i, j = i+1, j-1 {
		pre := s[i]
		suf := s[j]
		if pre == suf {

		}
	}
}

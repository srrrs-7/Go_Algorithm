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
		pre := string(s[i])
		suf := string(s[j])
		if pre == suf {
			res = pre
		}
		pre += string(s[i+1])
		suf += string(s[i-1])
	}
}

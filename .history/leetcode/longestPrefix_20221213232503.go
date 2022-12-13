package leetcode

import "fmt"

func LongestPrefix() {
	s := "level"

	res := longestPrefix(s)
	fmt.Println("longest prefix and suffix: ", res)
}

func longestPrefix(s string) string {
	var res string
	for i, j := 1, len(s); i < j; i, j = i+1, j-1 {
		pre := string(s[i-1])
		suf := string(s[j-1])
		if pre == suf {
			res = pre
		}
		pre += string(s[i+1])
		suf += string(s[i-1])
	}
	return res
}

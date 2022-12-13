package leetcode

import (
	"fmt"
	"strings"
)

func LongestPrefix() {
	s := "level"

	res := longestPrefix(s)
	fmt.Println("longest prefix and suffix: ", res)
}

func longestPrefix(s string) string {
	var res string
	var pre string
	var suf string

	for i, j := 0, len(s)-1; i < len(s)-2; i, j = i+1, j-1 {
		pre += string(s[i])
		suf += string(s[j])

		strings.Split(suf, "")
		if pre == suf {
			res = pre
		}
	}
	return res
}

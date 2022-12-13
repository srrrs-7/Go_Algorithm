package leetcode

import (
	"fmt"
	"strings"
)

func LongestPrefix() {
	s := "ababab"

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
		arr := strings.Split(suf, "")
		for i2, j2 := 0, len(arr)-1; i2 < j2; i2, j2 = i2+1, j2-1 {
			arr[i2], arr[j2] = arr[j2], arr[i2]
		}
		suff = strings.Join(arr, "")

		if pre == suf {
			res = pre
		}

		suf =
	}
	return res
}

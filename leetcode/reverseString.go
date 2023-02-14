package leetcode

import (
	"fmt"
	"strings"
)

func ReverseString() {
	s := "a good   example"
	res := reverseString(s)
	fmt.Println(res)
}

func reverseString(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Join(strings.Fields(s), " ")
	arr := strings.Split(s, " ")

	res := make([]string, len(arr))
	for l, r := 0, len(arr)-1; l <= len(arr)/2; l, r = l+1, r-1 {
		res[l] = arr[r]
		res[r] = arr[l]
	}

	return strings.Join(res, " ")
}

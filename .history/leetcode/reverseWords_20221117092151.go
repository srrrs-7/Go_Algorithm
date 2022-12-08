package leetcode

import (
	"strings"
)

func ReverseWords() {
	s := "the sky is blue"
	_ = reverseWord(s)

	// fmt.Println(res)
}

func reverseWord(s string) string {
	arr := strings.Split(s, " ")
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	str := strings.Join(arr, " ")
	return str
}
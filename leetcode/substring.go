package leetcode

import (
	"fmt"
	"strings"
)

func SubString() {
	s := "bbbbb"
	n := subString(s)
	fmt.Println(n)

	st := distinctString(s)
	fmt.Println(st)
}

func subString(s string) int {
	result := -1
	dict := make(map[byte]int)
	d := -1

	for i := 0; i < len(s); i++ {
		if v, ok := dict[s[i]]; !ok {
			dict[s[i]] = i
			d++
		} else {
			d = minString(i-v-1, d+1)
			dict[s[i]] = i
		}
		result = maxString(result, d)
	}
	return result + 1
}

func maxString(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func minString(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func distinctString(s string) string {
	var sb strings.Builder
	dict := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, ok := dict[s[i]]; !ok {
			dict[s[i]] = i
			sb.WriteString(string(s[i]))
		}
	}
	return sb.String()
}

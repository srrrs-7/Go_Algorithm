package leetcode

import (
	"fmt"
	"strings"
)

func GroupOfStrings() {
	words := []string{"a","b","ab","cde"}
	res := groupOfStrings(words)
	fmt.Println(res)
}

func groupOfStrings(words []string) []string {
	res := [2]int{0, 0} // [divide number, max length group]
	div := 0
	maxLen := 0

	for _, w1 := range words {
		for _, w2 := range words {
			if strings.Contains(w1, w2) { maxLen += 1 }
		}
		if res[1] < maxLen { res[1] = maxLen  }
	}

	return res
}
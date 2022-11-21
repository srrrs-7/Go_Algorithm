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
	res := [2]int{0, 0}

	for _, w1 := range words {
		for _, w2 := range words {
			if strings.Contains(w1, w2) {

			}
		}
	}

	return res
}
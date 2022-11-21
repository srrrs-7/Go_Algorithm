package leetcode

import "fmt"

func GroupOfStrings() {
	words := []string{"a","b","ab","cde"}
	res := groupOfStrings(words)
	fmt.Println(res)
}

func groupOfStrings(words []string) []string {
	res := [2]int{0, 0}

	for _, w1 := range words {
		for _, w2 := range words {}
	}

	return res
}
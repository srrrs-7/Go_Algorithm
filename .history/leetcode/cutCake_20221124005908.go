package leetcode

import "fmt"

func CutCake() {
	var L = 34
	var div = []int{8, 13, 26}

	res := cutCake(L, div)
	fmt.Println(res)
}

func cutCake(L int, div []int) int {
	var cut = len(div)
	var point = div[1]

	if div[cut + 1]
}
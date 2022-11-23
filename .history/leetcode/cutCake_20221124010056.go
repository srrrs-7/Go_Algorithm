package leetcode

import "fmt"

func CutCake() {
	var L = 34
	var div = []int{8, 13, 26}

	res := cutCake(L, div)
	fmt.Println(res)
}

func cutCake(L int, div []int) int {
	var cut = 1
	var point = div[k + 1]

	for _, d := range div {
		if point > div[d + 1]
	}
}
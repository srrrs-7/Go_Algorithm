package leetcode

import "fmt"

func CutCake() {
	var L = 34
	var div = []int{8, 13, 26}

	res := cutCake(L, div)
	fmt.Println(res)
}

func cutCake(L int, div []int) int {
	var point = 0

	for i := range div {
		if point > div[i + 1] {
			point = div[i + 1]
		}
	}
}
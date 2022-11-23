package leetcode

import "fmt"

func CutCake() {
	var L = 34
	var div = []int{8, 13, 26}

	res := cutCake(L, div)
	fmt.Println("cut cake: ", res)
}

func cutCake(L int, div []int) int {
	var point =

	for i := range div {
		if i >= len(div)-1 { break }
		if point > div[i + 1] {
			point = div[i + 1]
		}
	}

	return point
}
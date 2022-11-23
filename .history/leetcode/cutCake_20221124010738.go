package leetcode

import "fmt"

func CutCake() {
	var div = []int{0, 8, 13, 26, 34}

	res := cutCake(L, div)
	fmt.Println("cut cake: ", res)
}

func cutCake(L int, div []int) int {
	var point = 0

	for i := range div {
		if i >= len(div)-1 { break }
		if point > div[i + 1] {
			point = div[i + 1]
		}
	}

	return point
}
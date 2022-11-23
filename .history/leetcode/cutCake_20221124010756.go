package leetcode

import "fmt"

func CutCake() {
	var cake = []int{0, 8, 13, 26, 34}

	res := cutCake(cake)
	fmt.Println("cut cake: ", res)
}

func cutCake(cake []int) int {
	var point = 0

	for i := range div {
		if i >= len(div)-1 { break }
		if point > div[i + 1] {
			point = div[i + 1]
		}
	}

	return point
}
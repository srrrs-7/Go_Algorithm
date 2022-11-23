package leetcode

import "fmt"

func CutCake() {
	var cake = []int{0, 8, 13, 26, 34}

	res := cutCake(cake)
	fmt.Println("cut cake: ", res)
}

func cutCake(cake []int) int {
	var point = 0

	for i := range cake {
		if i >= len(cake)-1 { break }
		if cake[i] > cake[i + 2] {
			point = cake[i + 1]
		}
	}

	return point
}
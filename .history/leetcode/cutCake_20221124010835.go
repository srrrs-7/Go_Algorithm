package leetcode

import "fmt"

func CutCake() {
	var cake = []int{0, 8, 13, 26, 34}

	res := cutCake(cake)
	fmt.Println("cut cake: ", res)
}

func cutCake(cake []int) int {
	var point = cake[2]

	for i := range cake {
		if i >= len(cake)-1 { break }
		if point > cake[i + 1] {
			point = cake[i + 1]
		}
	}

	return point
}
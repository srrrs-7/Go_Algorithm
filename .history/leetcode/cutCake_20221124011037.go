package leetcode

import "fmt"

func CutCake() {
	var cake = []int{0, 8, 13, 26, 34}

	res := cutCake(cake)
	fmt.Println("cut cake: ", res)
}

func cutCake(cake []int) int {
	var point = cake[i + 2] - cake[i]

	for i := range cake {
		if i >= len(cake)-2 { break }
		if cake[i + 2] - cake[i] < point {
			point = cake[i + 2] - cake[i]
		}
	}

	return point
}
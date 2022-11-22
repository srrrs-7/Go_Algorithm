package leetcode

import "fmt"

func CountIdealArray() {
	num := 2
	maxValue := 5

	res := countIdealArray(num, maxValue)
	fmt.Printf("ideal array: %d \r", res)
}

func countIdealArray(num, maxValue int) int {
	res := 0

	for i := 1; i <= maxValue; i++ {
		if maxValue / i < 2  {
			res += 1
		} else {
			res += maxValue / i
		}
	}

	return res
}

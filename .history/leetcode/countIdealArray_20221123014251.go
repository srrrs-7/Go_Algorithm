package leetcode

import "fmt"

func CountIdealArray() {
	num := 5
	maxValue := 2

	res := countIdealArray(num, maxValue)
	fmt.Printf("ideal array: %d \n", res)
}

func countIdealArray(num, maxValue int) int {
	res := maxValue

	for i := 1; i <= maxValue; i++ {
		if i+i < maxValue {
			res += num - 1 / maxValue
		} else {
			res += 1
		}
	}

	return res
}

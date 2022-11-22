package leetcode

import "fmt"

func CountIdealArray() {
	num := 5
	maxValue := 3

	res := countIdealArray(num, maxValue)
	fmt.Printf("ideal array: %d \n", res)
}

func countIdealArray(num, maxValue int) int {
	res := 0

	for i := 1; i <= maxValue; i++ {
		res += 1
		if i+i < maxValue {
			res += maxValue / i - i
		}
	}

	return res
}

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
		if maxValue / i+i <= 1 {
			res += 1
		} else {

		}
	}

	if num >= maxValue { res += 1 }

	return res
}

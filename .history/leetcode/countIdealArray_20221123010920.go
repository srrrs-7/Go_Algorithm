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
		res += 1 // constant res++
		if i + i < maxValue {
			res += (maxValue-1) * (num-1)
		}
	}

	return res
}

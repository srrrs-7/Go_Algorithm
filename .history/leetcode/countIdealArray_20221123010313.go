package leetcode

import "fmt"

func CountIdealArray() {
	num := 2
	maxValue := 5

	res := countIdealArray(num, maxValue)
	fmt.Printf("ideal array: %d \n", res)
}

func countIdealArray(num, maxValue int) int {
	res := 0

	for i := 1; i <= maxValue; i++ {
		res += 1 // constant res++
		if i + i > maxValue {
		} else if i + i < maxValue {
			res += (maxValue-1) * (num - 1)
			res -= 1
		}
	}

	return res
}

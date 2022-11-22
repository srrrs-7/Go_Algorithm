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
		if i+i > maxValue {
			res += 1
		} else if i+i < maxValue {
			res += (maxValue-i) * (num-1)
		}
	}

	if num > maxValue { res += 1 }

	return res
}

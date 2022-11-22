package leetcode

import "fmt"

func CountIdealArray() {
	num := 2
	maxValue := 5

	res := countIdealArray(num, maxValue)
	fmt.Printf("ideal array: %d \n", res)
}

func countIdealArray(num, maxValue int) int {
	res := 1

	for i := 1; i <= maxValue; i++ {
		if i+i < maxValue {
			res += (num - 1) * (maxValue-1) / i
		} else {
			res += 1
		}
		fmt.Println(res)
	}

	return res
}

package leetcode

import "fmt"

func CountIdealArray() {
	num := 2
	maxValue := 5

	res := countIdealArray(num, maxValue)
	fmt.Printf("%d", res)
}

func countIdealArray(num, maxValue int) int {
	res := 0

	for i := 0; i < maxValue; i++ {
		if maxValue / i < 2  { res += 1 }
		else if
	}

	return res
}

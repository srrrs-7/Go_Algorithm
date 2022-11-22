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
		if i + i > maxValue { res += 1 }
		arr := sameValueArray(i)
	}

	return arr[0]
}

func sameValueArray(num int) []int {
    a := make([]int, num)
    for i := range a {
        a[i] = 1
    }
    return a
}
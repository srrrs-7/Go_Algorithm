package leetcode

import "fmt"

func CountIdealArray() {
	num := 2
	maxValue := 5

	res := countIdealArray(num, maxValue)
	fmt.Printf("%d", res)
}

func countIdealArray(num, maxValue int) int {
	for i := 0; i < maxValue; i++ {
		arr := sequenceArray(i)
	}

	return arr[0]
}

func sequenceArray(num int) []int {
    a := make([]int, max)
    for i := range a {
        a[i] = 1
    }
    return a
}
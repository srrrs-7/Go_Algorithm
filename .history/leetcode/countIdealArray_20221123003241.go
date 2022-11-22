package leetcode

import "fmt"

func CountIdealArray() {
	num := 2
	maxValue := 5

	res := countIdealArray(num, maxValue)
	fmt.Printf("%d", res)
}

func countIdealArray(num, maxValue int) int {


	for i := range maxValue {
		arr := sequenceArray(maxValue)
	}

	return arr[0]
}

func sequenceArray(max int) []int {
    a := make([]int, max)
    for i := range a {
        a[i] = 1
    }
    return a
}
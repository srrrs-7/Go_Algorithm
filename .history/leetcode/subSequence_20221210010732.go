package leetcode

import (
	"fmt"
	"math"
	"sort"
)

func SubSequence() {
	nums := []int{5, -7, 3, 5}
	goal := 6

	res := subSequence(nums, goal)
	fmt.Println("sub sequence sum: ", res)
}

func subSequence(nums []int, goal int) int {
	arr := getAllSum(nums)
	sort.Ints(arr)

	res := getMinVal(arr, goal)

	return res
}

func getAllSum(nums []int) []int {
	n := len(nums)

	var res []int

	var iter func(idx, sumSoFar int)
	iter = func(idx, sumSoFar int) {
		if idx == n {
			res = append(res, sumSoFar)
			return
		}

		iter(idx+1, sumSoFar)
		iter(idx+1, sumSoFar+nums[idx])
	}

	iter(0, 0)

	return res
}

func getMinVal(arr []int, goal int) int {
	minSoFar := math.MaxInt

	i := 0
	j := len(arr) - 1

	for i < len(arr) && j >= 0 {
		if arr[i] < goal {
			i++
		} else if sum > goal {
			j--
		} else {
			return 0
		}
	}

	return minSoFar
}

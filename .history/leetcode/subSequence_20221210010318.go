package leetcode

import (
	"fmt"
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

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {

	}

	return 111111
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

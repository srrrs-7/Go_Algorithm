package leetcode

import "fmt"

func SingleAndTwice() {
	arr := []int{1, 1, 2, 2, 3, 3, 5, 5, 8, 8, 9}
	res := singleAndTwice(arr)
	fmt.Println(res)
}

func singleAndTwice(nums []int) int {
	// left and right length
	l, r := 0, len(nums)-1
	// binary search
	for l < r {
		// middle index -> if mid == even -> single int is right side exist
		mid := l + (r-l)/2
		// odd -> mid - 1
		if mid%2 == 1 {
			mid--
		}
		// check duplicate
		if nums[mid] != nums[mid+1] {
			r = mid
		} else {
			l = mid + 2
		}
	}
	return nums[l]
}

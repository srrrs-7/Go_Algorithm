package algorithm

func AllSum(nums []int) []int {
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

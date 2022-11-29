package atCorder

import "fmt"

func Dp() {
	l := []int{1, 2, 3}
	t := 7

	res := dp(l, t)
	fmt.Println("multi number count: ", res)
}

// digit DP
func dp(l []int, t int) int {
	digit := len(l)
	var dp[0][0][0] [][][]int
}
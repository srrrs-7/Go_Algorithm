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
	var dp[100][2][2] [][][]int
	digit := len(l)

	dp[0][0][0] = 1 // dp[d][s][t]
	for i := 0; i < l; i++ {}
}
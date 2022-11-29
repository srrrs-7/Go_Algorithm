package atCorder

import "fmt"

func Dp() {
	n := []int{1, 2, 3}
	t := 7

	res := dp(n, t)
	fmt.Println("multi number count: ", res)
}

// digit DP
func dp(n []int, t int) int {
	var dp[100][2][2] int
	digit := len(n)

	dp[0][0][0] = 1 // dp[d][s][t]
	for i := 0; i < digit; i++ {
		for smaller := 0; smaller < 2; smaller++ {
			for j := 0; j < 2; j++ {
				for x := 0; x <= smaller; x++ {
					if smaller == 1 {
						smaller = 9
					} else {
						smaller = n[i]
					}
					dp[i + 1][smaller || x < n[i]][j || x == 3]
				}
			}
		}
		}

}
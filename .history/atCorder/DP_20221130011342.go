package atCorder

import "fmt"

func Dp() {
	n := []int{1, 2, 3}
	req := 7

	res := dp(n, req)
	fmt.Println("digit DP: ", res)
}

// digit DP
func dp(n []int, req int) int {
	var dp[100][2][2] int
	l := len(n)

	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}

	dp[0][0][0] = 1 // dp[d][s][t]
	for i := 0; i < l; i++ {
		for smaller := 0; smaller < 2; smaller++ {
			for j := 0; j < 2; j++ {
				if smaller == 1 {
					smaller = 9
				} else {
					smaller = n[i]
				}
				for x := 0; x <= small; x++ {
					var s = 0
					if smaller == 1 || x < n[i] { s = 1 }
					var t = 0
					if j == 1 || x == 3 { t = 1 }

					dp[i + 1][s][t] += dp[i][smaller][j]
				}
			}
		}
	}

	return dp[l][0][1] + dp[l][1][1]
}
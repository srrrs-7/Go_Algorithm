package atcorder

import "fmt"

func Dp() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	req := 7

	res := dp(n, req)
	fmt.Println("digit DP: ", res)
}

// digit DP
func dp(n []int, req int) int {
	var dp [100][2][2]int
	l := len(n)

	dp[0][0][0] = 1          // dp[d][s][t]
	for i := 0; i < l; i++ { // digit loop
		for smaller := 0; smaller < 2; smaller++ { // smaller loop
			for j := 0; j < 2; j++ { //target loop
				var iter = n[i]
				if smaller == 1 {
					iter = 9
				} // more than 1?
				for x := 0; x <= iter; x++ { // 9 or digit 1 value loop
					var s = 0
					if smaller == 1 || x < n[i] {
						s = 1
					} // smaller ?
					var t = 0
					if j == 1 || x == req {
						t = 1
					} // req value compare
					dp[i+1][s][t] += dp[i][smaller][j] // each digit result sum
				}
			}
		}
	}
	return dp[l][0][1] + dp[l][1][1]
}

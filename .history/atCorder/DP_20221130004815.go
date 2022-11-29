package atCorder

import "fmt"

func Dp() {
	n := []int{1, 2, 3}
	req := 7

	res := dp(n, req)
	fmt.Println("multi number count: ", res)
}

// digit DP
func dp(n []int, req int) int {
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
					var s bool
					var t bool
					if smaller == 1 || x < n[i] {
						s = true
					}
					dp[i + 1][s][t]
				}
			}
		}
		}

}
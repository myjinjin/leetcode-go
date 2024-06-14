package leetcode

import "math"

func numTilings(n int) int {
	mod := math.Pow10(9) + 7
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5

	for i := 4; i < n+1; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % int(mod)
	}

	return dp[n]
}

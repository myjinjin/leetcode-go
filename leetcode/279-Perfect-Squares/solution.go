package leetcode

func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

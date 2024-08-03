package leetcode

func maxPoints(points [][]int) int64 {
	m, n := len(points), len(points[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	copy(dp[0], points[0])

	for i := 1; i < m; i++ {
		for j := n - 2; j >= 0; j-- {
			dp[i-1][j] = max(dp[i-1][j], dp[i-1][j+1]-1)
		}
		for j := 1; j < n; j++ {
			dp[i-1][j] = max(dp[i-1][j], dp[i-1][j-1]-1)
		}
		for j := 0; j < n; j++ {
			dp[i][j] = dp[i-1][j] + points[i][j]
		}
	}

	result := dp[m-1][0]
	for j := 1; j < n; j++ {
		result = max(result, dp[m-1][j])
	}

	return int64(result)
}

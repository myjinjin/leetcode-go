package leetcode

func minimumDeletions(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	bCount := 0
	for i := 0; i < n; i++ {
		if s[i] == 'a' {
			dp[i+1] = min(dp[i]+1, bCount)
		} else if s[i] == 'b' {
			bCount++
			dp[i+1] = dp[i]
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

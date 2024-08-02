package leetcode

func maximizeWin(prizePositions []int, k int) int {
	n := len(prizePositions)
	dp := make([]int, n+1)
	left := 0
	result := 0

	for right := 0; right < n; right++ {
		for prizePositions[left] < prizePositions[right]-k {
			left++
		}

		dp[right+1] = max(dp[right], right-left+1)
		result = max(result, right-left+1+dp[left])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

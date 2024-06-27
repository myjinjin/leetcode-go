package leetcode

import "math"

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i < n-1; i++ {
		for jump := 1; jump <= nums[i] && i+jump < n; jump++ {
			dp[i+jump] = min(dp[i+jump], dp[i]+1)
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package leetcode

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	longest := 1
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1]+1 {
			dp[i] = dp[i-1] + 1
			longest = max(longest, dp[i])
		} else if nums[i] == nums[i-1] {
			dp[i] = dp[i-1]
		}
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package leetcode

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	result := nums[0]

	for i := 1; i < n; i++ {
		dp[i] = nums[i]
		if dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}
		result = max(result, dp[i])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

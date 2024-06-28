package leetcode

func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[0] = true
	for i := 0; i < n; i++ {
		for jump := 1; jump <= nums[i] && i+jump < n; jump++ {
			dp[i+jump] = dp[i]
		}
	}
	return dp[n-1]
}

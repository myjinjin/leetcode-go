package leetcode

func canPartition(nums []int) bool {
	target := 0
	for _, num := range nums {
		target += num
	}
	if target%2 != 0 {
		return false
	}
	target /= 2

	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for i := target; i > 0; i-- {
			if i >= num {
				dp[i] = dp[i] || dp[i-num]
			}
		}
	}
	return dp[target]
}

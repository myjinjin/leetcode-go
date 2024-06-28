package leetcode

func subarraySum(nums []int, k int) int {
	count := 0
	n := len(nums)
	prefixSum := make([]int, n+1)
	prefixSum[0] = 0

	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	for left := 0; left < n; left++ {
		for right := left + 1; right < n+1; right++ {
			if prefixSum[right]-prefixSum[left] == k {
				count++
			}
		}
	}
	return count
}

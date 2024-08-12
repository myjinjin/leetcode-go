package leetcode

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)

	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	maxSum := 0
	for i := firstLen; i <= n-secondLen; i++ {
		firstSum := prefixSum[i] - prefixSum[i-firstLen]
		secondSum := 0
		for j := i + secondLen; j <= n; j++ {
			secondSum = max(secondSum, prefixSum[j]-prefixSum[j-secondLen])
		}
		maxSum = max(maxSum, firstSum+secondSum)
	}

	for i := secondLen; i <= n-firstLen; i++ {
		secondSum := prefixSum[i] - prefixSum[i-secondLen]
		firstSum := 0
		for j := i + firstLen; j <= n; j++ {
			firstSum = max(firstSum, prefixSum[j]-prefixSum[j-firstLen])
		}
		maxSum = max(maxSum, firstSum+secondSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

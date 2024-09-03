package leetcode

func longestNiceSubarray(nums []int) int {
	maxLen := 1
	start := 0
	used := 0

	for end := 0; end < len(nums); end++ {
		for (used & nums[end]) != 0 {
			used ^= nums[start]
			start++
		}
		used |= nums[end]
		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}

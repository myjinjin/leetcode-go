package leetcode

import (
	"math"
	"sort"
)

func minDifference(nums []int) int {
	n := len(nums)
	if n < 5 {
		return 0
	}

	sort.Ints(nums)
	minDiff := math.MaxInt32
	for i := 0; i < 4; i++ {
		if i < n {
			left := i
			right := n - 4 + i
			currDiff := nums[right] - nums[left]
			if currDiff < minDiff {
				minDiff = currDiff
			}
		}
	}

	if minDiff < math.MaxInt32 {
		return minDiff
	}

	return 0
}

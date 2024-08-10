package leetcode

import (
	"sort"
)

func maximumGap(nums []int) int {
	result := 0

	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		result = max(result, nums[i]-nums[i-1])
	}

	return result
}

package leetcode

import "sort"

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	avg := make(map[float64]bool)
	for i := 0; i < len(nums)/2; i++ {
		avg[float64(nums[i]+nums[len(nums)-1-i])/2.0] = true
	}
	return len(avg)
}

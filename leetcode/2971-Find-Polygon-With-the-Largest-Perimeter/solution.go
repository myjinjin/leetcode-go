package leetcode

import "sort"

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	n := len(nums)
	sum := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	for i := n - 1; i >= 2; i-- {
		if nums[i] < sum[i] {
			return int64(sum[i+1])
		}
	}
	return -1
}

package leetcode

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	count := 0

	for i := n - 1; i >= 2; i-- {
		left := 0
		right := i - 1

		for left < right {
			if nums[left]+nums[right] > nums[i] {
				count += right - left
				right--
			} else {
				left++
			}
		}
	}

	return count
}

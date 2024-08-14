package leetcode

import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	left, right := 0, nums[n-1]-nums[0]

	for left < right {
		mid := left + (right-left)/2
		count := countPairs(nums, mid)

		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func countPairs(nums []int, distance int) int {
	count, left := 0, 0
	for right := 0; right < len(nums); right++ {
		for nums[right]-nums[left] > distance {
			left++
		}
		count += right - left
	}
	return count
}

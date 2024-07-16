package leetcode

import "sort"

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)

	left, right := 0, len(nums)-1
	result := 0
	mod := 1000000007

	pow2 := make([]int, len(nums))
	pow2[0] = 1
	for i := 1; i < len(nums); i++ {
		pow2[i] = (pow2[i-1] * 2) % mod
	}

	for left <= right {
		if nums[left]+nums[right] > target {
			right--
		} else {
			result = (result + pow2[right-left]) % mod
			left++
		}
	}

	return result
}

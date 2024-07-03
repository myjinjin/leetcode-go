package leetcode

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 {
		return
	}

	left := nums[n-k:]
	right := nums[:n-k]
	result := append(left, right...)
	copy(nums, result)
}

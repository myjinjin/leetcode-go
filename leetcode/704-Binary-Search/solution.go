package leetcode

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right && left < len(nums) && right >= 0 {
		mid := (left + right) / 2
		curr := nums[mid]
		if curr == target {
			return mid
		} else if curr > target {
			right = mid - 1
		} else if curr < target {
			left = mid + 1
		}
	}
	return -1
}

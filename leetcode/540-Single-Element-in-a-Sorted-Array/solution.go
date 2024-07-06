package leetcode

func singleNonDuplicate(nums []int) int {
	n := len(nums)
	left := 0
	right := n - 1

	for left < right {
		mid := (left + right) / 2
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] != nums[mid+1] {
			right = mid
		} else {
			left = mid + 2
		}
	}

	return nums[left]
}

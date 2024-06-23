package leetcode

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	result := []int{-1, -1}

	if len(nums) == 0 {
		return result
	}

	for left < right {
		mid := (left + right) / 2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] != target {
		return result
	}

	result[0] = left
	right = len(nums) - 1
	for left < right {
		mid := (left+right)/2 + 1
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	result[1] = right
	return result
}

package leetcode

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	left, right := 0, len(nums)-1
	for {
		mid := left + (right-left)/2
		if mid == 0 {
			if nums[mid] > nums[mid+1] {
				return mid
			}
			left = mid + 1
			continue
		}

		if mid == len(nums)-1 {
			if nums[mid] > nums[mid-1] {
				return mid
			}
			right = mid - 1
			continue
		}

		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}
		if nums[mid] < nums[mid-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}

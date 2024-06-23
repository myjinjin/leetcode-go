package leetcode

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	pivot := left

	leftNums := nums[:pivot]
	rightNums := nums[pivot:]
	if len(leftNums) > 0 && leftNums[0] <= target && target <= leftNums[len(leftNums)-1] {
		l, r := 0, len(leftNums)-1
		for l <= r {
			mid := (l + r) / 2
			if leftNums[mid] == target {
				return mid
			}
			if leftNums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	if len(rightNums) > 0 && rightNums[0] <= target && target <= rightNums[len(rightNums)-1] {
		l, r := 0, len(rightNums)-1
		for l <= r {
			mid := (l + r) / 2
			if rightNums[mid] == target {
				return pivot + mid
			}
			if rightNums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

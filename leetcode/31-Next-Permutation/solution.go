package leetcode

func nextPermutation(nums []int) {
	descendingStartIndex := len(nums) - 1
	swapIndex := len(nums) - 1

	for descendingStartIndex > 0 && nums[descendingStartIndex-1] >= nums[descendingStartIndex] {
		descendingStartIndex--
	}

	if descendingStartIndex == 0 {
		reverse(nums, 0, len(nums)-1)
		return
	}

	pivotIndex := descendingStartIndex - 1

	for nums[swapIndex] <= nums[pivotIndex] {
		swapIndex--
	}

	nums[pivotIndex], nums[swapIndex] = nums[swapIndex], nums[pivotIndex]
	reverse(nums, descendingStartIndex, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

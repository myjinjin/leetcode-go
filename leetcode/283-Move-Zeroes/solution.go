package leetcode

func moveZeroes(nums []int) {
	left := 0

	for right := range nums {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
	}

	for i := left; i < len(nums); i++ {
		nums[i] = 0
	}
}

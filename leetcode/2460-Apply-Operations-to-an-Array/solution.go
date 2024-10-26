package leetcode

func applyOperations(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = 2 * nums[i]
			nums[i+1] = 0
		}
	}

	i := 0
	for _, n := range nums {
		if n != 0 {
			result[i] = n
			i++
		}
	}

	return result
}

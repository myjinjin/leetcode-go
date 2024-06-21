package leetcode

func permute(nums []int) [][]int {
	result := [][]int{}
	var backtrack func(result *[][]int, curr []int)
	backtrack = func(result *[][]int, curr []int) {
		if len(curr) == len(nums) {
			*result = append(*result, append([]int{}, curr...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if contains(curr, nums[i]) {
				continue
			}
			curr = append(curr, nums[i])
			backtrack(result, curr)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(&result, []int{})
	return result
}

func contains(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}

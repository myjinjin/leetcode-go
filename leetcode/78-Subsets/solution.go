package leetcode

func subsets(nums []int) [][]int {
	result := [][]int{}

	var backtrack func(result *[][]int, curr []int, start int)
	backtrack = func(result *[][]int, curr []int, start int) {
		*result = append(*result, append([]int{}, curr...))
		if len(curr) == len(nums) {
			return
		}

		for i := start; i < len(nums); i++ {
			curr = append(curr, nums[i])
			backtrack(result, curr, i+1)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(&result, []int{}, 0)
	return result
}

package leetcode

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	var backtrack func(curr []int, start int, target int, result *[][]int)
	backtrack = func(curr []int, start int, target int, result *[][]int) {
		if target == 0 {
			*result = append(*result, append([]int{}, curr...))
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				continue
			}

			curr = append(curr, candidates[i])
			backtrack(curr, i, target-candidates[i], result)
			curr = curr[:len(curr)-1]
		}
	}
	backtrack([]int{}, 0, target, &result)
	return result
}

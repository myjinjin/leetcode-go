package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)

	var backtrack func(curr []int, currSum int, start int)
	backtrack = func(curr []int, currSum int, start int) {
		if currSum == target {
			result = append(result, append([]int{}, curr...))
			return
		}

		if currSum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			backtrack(append(curr, candidates[i]), currSum+candidates[i], i+1)
		}
	}

	backtrack([]int{}, 0, 0)

	return result
}

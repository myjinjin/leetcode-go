package leetcode

func combinationSum3(k int, n int) [][]int {
	maxNum := min(9, n)
	candidates := []int{}
	for i := 1; i <= maxNum; i++ {
		candidates = append(candidates, i)
	}

	combinations := [][]int{}
	var backtrack func(curr []int, currSum int, index int, combinations *[][]int)
	backtrack = func(curr []int, currSum int, index int, combinations *[][]int) {
		if len(curr) == k {
			if currSum == n {
				copied := make([]int, len(curr))
				copy(copied, curr)
				*combinations = append(*combinations, copied)
			}
			return
		}
		if index == len(candidates) {
			return
		}
		if currSum > n {
			return
		}
		backtrack(curr, currSum, index+1, combinations)
		backtrack(append(curr, candidates[index]), currSum+candidates[index], index+1, combinations)
	}

	backtrack([]int{}, 0, 0, &combinations)
	return combinations
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

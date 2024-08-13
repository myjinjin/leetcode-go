# [40. Combination Sum II](https://leetcode.com/problems/combination-sum-ii/)

## Problem

Given a collection of candidate numbers (`candidates`) and a target number (`target`), find all unique combinations in `candidates` where the candidate numbers sum to `target`.

Each number in `candidates` may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.


Example 1:

```
Input: candidates = [10,1,2,7,6,1,5], target = 8
Output: 
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
```

Example 2:

```
Input: candidates = [2,5,2,1,2], target = 5
Output: 
[
[1,2,2],
[5]
]
``` 

Constraints:

- `1 <= candidates.length <= 100`
- `1 <= candidates[i] <= 50`
- `1 <= target <= 30`

## Solution

```go
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
```
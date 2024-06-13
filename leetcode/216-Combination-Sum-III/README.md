# [216. Combination Sum III](https://leetcode.com/problems/combination-sum-iii/)

## Problem

Find all valid combinations of `k` numbers that sum up to `n` such that the following conditions are true:

- Only numbers `1` through `9` are used.
- Each number is used **at most once**.

Return a list of all possible valid combinations. The list must not contain the same combination twice, and the combinations may be returned in any order.

 

Example 1:

```
Input: k = 3, n = 7
Output: [[1,2,4]]
Explanation:
1 + 2 + 4 = 7
There are no other valid combinations.
```

Example 2:

```
Input: k = 3, n = 9
Output: [[1,2,6],[1,3,5],[2,3,4]]
Explanation:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
There are no other valid combinations.
```

Example 3:

```
Input: k = 4, n = 1
Output: []
Explanation: There are no valid combinations.
Using 4 different numbers in the range [1,9], the smallest sum we can get is 1+2+3+4 = 10 and since 10 > 1, there are no valid combination.
```
 

Constraints:

- `2 <= k <= 9`
- `1 <= n <= 60`


## Solution

```go
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
```
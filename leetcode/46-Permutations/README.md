# [46. Permutations](https://leetcode.com/problems/permutations/)

## Problem

Given an array `nums` of distinct integers, return all the possible permutations. You can return the answer in any order.

 

Example 1:

```
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

Example 2:

```
Input: nums = [0,1]
Output: [[0,1],[1,0]]
```

Example 3:

```
Input: nums = [1]
Output: [[1]]
``` 

Constraints:

- `1 <= nums.length <= 6`
- `-10 <= nums[i] <= 10`
- All the integers of `nums` are unique.

## Solution

```go
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
```
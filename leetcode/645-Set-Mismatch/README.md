# [645. Set Mismatch](https://leetcode.com/problems/set-mismatch/)

## Problem

You have a set of integers `s`, which originally contains all the numbers from `1` to `n`. Unfortunately, due to some error, one of the numbers in `s` got duplicated to another number in the set, which results in repetition of **one** number and loss of another number.

You are given an integer array `nums` representing the data status of this set after the error.

Find the number that occurs twice and the number that is missing and return them in the form of an array.



Example 1:

```
Input: nums = [1,2,2,4]
Output: [2,3]
```

Example 2:

```
Input: nums = [1,1]
Output: [1,2]
```

Constraints:

- `2 <= nums.length <= 10^4`
- `1 <= nums[i] <= 10^4`

## Solution

```go
func findErrorNums(nums []int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}
	duplicated, missing := -1, -1
	for i := 1; i <= len(nums); i++ {
		if val, exist := count[i]; !exist {
			missing = i
		} else if val == 2 {
			duplicated = i
		}
	}
	return []int{duplicated, missing}
}
```
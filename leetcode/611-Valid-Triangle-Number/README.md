# [611. Valid Triangle Number](https://leetcode.com/problems/valid-triangle-number/)

## Problem

Given an integer array `nums`, return the number of triplets chosen from the array that can make triangles if we take them as side lengths of a triangle.


Example 1:

```
Input: nums = [2,2,3,4]
Output: 3
Explanation: Valid combinations are: 
2,3,4 (using the first 2)
2,3,4 (using the second 2)
2,2,3
```

Example 2:

```
Input: nums = [4,2,3,4]
Output: 4
``` 

Constraints:

- `1 <= nums.length <= 1000`
- `0 <= nums[i] <= 1000`

## Solution

```go
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	count := 0

	for i := n - 1; i >= 2; i-- {
		left := 0
		right := i - 1

		for left < right {
			if nums[left]+nums[right] > nums[i] {
				count += right - left
				right--
			} else {
				left++
			}
		}
	}

	return count
}
```
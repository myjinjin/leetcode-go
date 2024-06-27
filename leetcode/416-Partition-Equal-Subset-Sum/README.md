# [416. Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/)

## Problem

Given an integer array `nums`, return `true` if you can partition the array into two subsets such that the sum of the elements in both subsets is equal or `false` otherwise.

 
Example 1:

```
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].
```

Example 2:

```
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.
``` 

Constraints:

- `1 <= nums.length <= 200`
- `1 <= nums[i] <= 100`

## Solution

```go
func canPartition(nums []int) bool {
	target := 0
	for _, num := range nums {
		target += num
	}
	if target%2 != 0 {
		return false
	}
	target /= 2

	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for i := target; i > 0; i-- {
			if i >= num {
				dp[i] = dp[i] || dp[i-num]
			}
		}
	}
	return dp[target]
}
```
# [53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)

## Problem

Given an integer array `nums`, find the 
subarray with the largest sum, and return its sum.

Example 1:

```
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
```

Example 2:

```
Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
```

Example 3:

```
Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
```

Constraints:

- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`

## Solution

```go
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)    // dp[i]: nums[i]로 끝나는 maxSubArray
	dp[0] = nums[0]
	result := nums[0]

	for i := 1; i < n; i++ {
		dp[i] = nums[i]
		if dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}
		result = max(result, dp[i])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
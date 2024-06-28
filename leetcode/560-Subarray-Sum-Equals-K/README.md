# [560. Subarray Sum Equals K](https://leetcode.com/problems/subarray-sum-equals-k/)

## Problem

Given an array of integers `nums` and an integer `k`, return the total number of subarrays whose sum equals to `k`.

A subarray is a contiguous non-empty sequence of elements within an array.

 

Example 1:

```
Input: nums = [1,1,1], k = 2
Output: 2
```

Example 2:

```
Input: nums = [1,2,3], k = 3
Output: 2
``` 

Constraints:

- `1 <= nums.length <= 2 * 10^4`
- `-1000 <= nums[i] <= 1000`
- `-10^7 <= k <= 10^7`

## Solution

```go
func subarraySum(nums []int, k int) int {
	count := 0
	n := len(nums)
	prefixSum := make([]int, n+1)
	prefixSum[0] = 0

	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	for left := 0; left < n; left++ {
		for right := left + 1; right < n+1; right++ {
			if prefixSum[right]-prefixSum[left] == k {
				count++
			}
		}
	}
	return count
}
```
# [1031. Maximum Sum of Two Non-Overlapping Subarrays](https://leetcode.com/problems/maximum-sum-of-two-non-overlapping-subarrays/)

## Problem

Given an integer array `nums` and two integers `firstLen` and `secondLen`, return the maximum sum of elements in two non-overlapping subarrays with lengths `firstLen` and `secondLen`.

The array with length `firstLen` could occur before or after the array with length `secondLen`, but they have to be non-overlapping.

A **subarray** is a **contiguous** part of an array.


Example 1:

```
Input: nums = [0,6,5,2,2,5,1,9,4], firstLen = 1, secondLen = 2
Output: 20
Explanation: One choice of subarrays is [9] with length 1, and [6,5] with length 2.
```

Example 2:

```
Input: nums = [3,8,1,3,2,1,8,9,0], firstLen = 3, secondLen = 2
Output: 29
Explanation: One choice of subarrays is [3,8,1] with length 3, and [8,9] with length 2.
```

Example 3:

```
Input: nums = [2,1,5,6,0,9,5,0,3,8], firstLen = 4, secondLen = 3
Output: 31
Explanation: One choice of subarrays is [5,6,0,9] with length 4, and [0,3,8] with length 3.
``` 

Constraints:

- `1 <= firstLen, secondLen <= 1000`
- `2 <= firstLen + secondLen <= 1000`
- `firstLen + secondLen <= nums.length <= 1000`
-`0 <= nums[i] <= 1000`

## Solution

```go
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)

	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	maxSum := 0
	for i := firstLen; i <= n-secondLen; i++ {
		firstSum := prefixSum[i] - prefixSum[i-firstLen]
		secondSum := 0
		for j := i + secondLen; j <= n; j++ {
			secondSum = max(secondSum, prefixSum[j]-prefixSum[j-secondLen])
		}
		maxSum = max(maxSum, firstSum+secondSum)
	}

	for i := secondLen; i <= n-firstLen; i++ {
		secondSum := prefixSum[i] - prefixSum[i-secondLen]
		firstSum := 0
		for j := i + firstLen; j <= n; j++ {
			firstSum = max(firstSum, prefixSum[j]-prefixSum[j-firstLen])
		}
		maxSum = max(maxSum, firstSum+secondSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
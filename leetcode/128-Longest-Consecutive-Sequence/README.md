# [128. Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/)

## Problem

Given an unsorted array of integers `nums`, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in `O(n)` time.

 

Example 1:

```
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
```

Example 2:

```
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
``` 

Constraints:

- `0 <= nums.length <= 10^5`
- `-10^9 <= nums[i] <= 10^9`

## Solution

```go
func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	longest := 1
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1]+1 {
			dp[i] = dp[i-1] + 1
			longest = max(longest, dp[i])
		} else if nums[i] == nums[i-1] {
			dp[i] = dp[i-1]
		}
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
# [440. K-th Smallest in Lexicographical Order](https://leetcode.com/problems/k-th-smallest-in-lexicographical-order/)

## Problem

Given two integers `n` and `k`, return the `kth` lexicographically smallest integer in the range `[1, n]`.


Example 1:

```
Input: n = 13, k = 2
Output: 10
Explanation: The lexicographical order is [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9], so the second smallest number is 10.
```

Example 2:

```
Input: n = 1, k = 1
Output: 1
``` 

Constraints:

- `1 <= k <= n <= 10^9`

## Solution

```go
func findKthNumber(n int, k int) int {
	curr := 1
	k--

	for k > 0 {
		count := countNumbersInRange(n, curr, curr+1)
		if count <= k {
			curr++
			k -= count
		} else {
			curr *= 10
			k--
		}
	}

	return curr
}

func countNumbersInRange(maxNum, start, nextStart int) int {
	count := 0
	for start <= maxNum {
		count += min(maxNum+1, nextStart) - start
		start *= 10
		nextStart *= 10
	}
	return count
}
```
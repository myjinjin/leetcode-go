# [658. Find K Closest Elements](https://leetcode.com/problems/find-k-closest-elements/)

## Problem

Given a sorted integer array `arr`, two integers `k` and `x`, return the `k` closest integers to `x` in the array. The result should also be sorted in ascending order.

An integer `a` is closer to `x` than an integer `b` if:

- `|a - x| < |b - x|`, or
- `|a - x| == |b - x|` and `a < b`
 

Example 1:

```
Input: arr = [1,2,3,4,5], k = 4, x = 3
Output: [1,2,3,4]
```

Example 2:

```
Input: arr = [1,2,3,4,5], k = 4, x = -1
Output: [1,2,3,4]
``` 

Constraints:

- `1 <= k <= arr.length`
- `1 <= arr.length <= 10^4`
- `arr` is sorted in ascending order.
- `-10^4 <= arr[i], x <= 10^4`

## Solution

```go
func findClosestElements(arr []int, k int, x int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if abs(x, arr[i]) == abs(x, arr[j]) {
			return arr[i] < arr[j]
		}
		return abs(x, arr[i]) < abs(x, arr[j])
	})

	result := arr[:k]
	sort.Ints(result)
	return result
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}
```
# [978. Longest Turbulent Subarray](https://leetcode.com/problems/longest-turbulent-subarray/)

## Problem

Given an integer array `arr`, return the length of a maximum size turbulent subarray of `arr`.

A subarray is **turbulent** if the comparison sign flips between each adjacent pair of elements in the subarray.

More formally, a subarray `[arr[i], arr[i + 1], ..., arr[j]]` of `arr` is said to be turbulent if and only if:

- For `i <= k < j`:
    - `arr[k] > arr[k + 1]` when `k` is odd, and
    - `arr[k] < arr[k + 1]` when `k` is even.
- Or, for `i <= k < j`:
    - `arr[k] > arr[k + 1]` when `k` is even, and
    - `arr[k] < arr[k + 1]` when `k` is odd.


Example 1:

```
Input: arr = [9,4,2,10,7,8,8,1,9]
Output: 5
Explanation: arr[1] > arr[2] < arr[3] > arr[4] < arr[5]
```

Example 2:

```
Input: arr = [4,8,12,16]
Output: 2
```

Example 3:

```
Input: arr = [100]
Output: 1
```

Constraints:

- `1 <= arr.length <= 4 * 10^4`
- `0 <= arr[i] <= 10^9`

## Solution

```go
func maxTurbulenceSize(arr []int) int {
	if len(arr) == 1 {
		return 1
	}

	maxSize := 1
	up := 1
	down := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			up = down + 1
			down = 1
		} else if arr[i] < arr[i-1] {
			down = up + 1
			up = 1
		} else {
			up = 1
			down = 1
		}
		maxSize = max(maxSize, max(up, down))
	}

	return maxSize
}
```
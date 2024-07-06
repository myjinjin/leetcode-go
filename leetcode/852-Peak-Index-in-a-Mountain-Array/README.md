# [852. Peak Index in a Mountain Array](https://leetcode.com/problems/peak-index-in-a-mountain-array/)

## Problem

You are given an integer mountain array `arr` of length `n` where the values increase to a peak element and then decrease.

Return the index of the peak element.

Your task is to solve it in `O(log(n))` time complexity.


Example 1:

```
Input: arr = [0,1,0]
Output: 1
```

Example 2:

```
Input: arr = [0,2,1,0]
Output: 1
```

Example 3:

```
Input: arr = [0,10,5,2]
Output: 1
```
 

Constraints:

- `3 <= arr.length <= 10^5`
- `0 <= arr[i] <= 10^6`
- `arr` is guaranteed to be a mountain array.


## Solution

```go
func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	left := 0
	right := n - 1

	for left < right {
		mid := (left + right) / 2
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
```
# [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)

## Problem

Given an integer array `nums` and an integer `k`, return the `kth` largest element in the array.

Note that it is the `kth` largest element in the sorted order, not the `kth` distinct element.

Can you solve it without sorting?

Example 1:

```
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
```

Example 2:

```
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4
```

Constraints:

- `1 <= k <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`

## Solution

```go
func findKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)-1
	k = k - 1
	for l < r {
		i, j := quickSelect(l, r, nums)
		if k <= j {
			r = j
		} else if k >= i {
			l = i
		} else {
			break
		}
	}
	return nums[k]
}

func quickSelect(l, r int, nums []int) (int, int) {
	pivot := nums[(l+r)/2]
	for l <= r {
		for nums[r] < pivot {
			r--
		}
		for nums[l] > pivot {
			l++
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	return l, r
}
```

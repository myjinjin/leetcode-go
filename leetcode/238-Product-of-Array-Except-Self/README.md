# [238. Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/)

## Problem

Given an integer array `nums`, return an array `answer` such that `answer[i]` is equal to the product of all the elements of `nums` except `nums[i]`.

The product of any prefix or suffix of `nums` is guaranteed to fit in a **32-bit integer**.

You must write an algorithm that runs in `O(n)` time and without using the division operation.


Example 1:

```
Input: nums = [1,2,3,4]
Output: [24,12,8,6]
```

Example 2:

```
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
``` 

Constraints:

- `2 <= nums.length <= 10^5`
- `-30 <= nums[i] <= 30`
- The product of any prefix or suffix of `nums` is guaranteed to fit in a 32-bit integer.


## Solution

```go
func productExceptSelf(nums []int) []int {
	n := len(nums)

	product := make([]int, n)
	product[0] = 1
	for i := 1; i < n; i++ {
		product[i] = product[i-1] * nums[i-1]
	}

	right := 1
	for i := n - 2; i >= 0; i-- {
		right *= nums[i+1]
		product[i] = product[i] * right
	}

	return product
}
```
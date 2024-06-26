# [152. Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/)

## Problem

Given an integer array `nums`, find a subarray that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.


Example 1:

```
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
```

Example 2:

```
Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
``` 

Constraints:

- `1 <= nums.length <= 2 * 10^4`
- `-10 <= nums[i] <= 10`
- The product of any prefix or suffix of `nums` is guaranteed to fit in a 32-bit integer.


## Solution

```go
func maxProduct(nums []int) int {
	n := len(nums)

	maxProduct := nums[0]
	minProduct := nums[0]
	result := maxProduct

	for i := 1; i < n; i++ {
		curr := nums[i]
		if curr < 0 {
			minProduct, maxProduct = maxProduct, minProduct
		}
		minProduct = min(curr, curr*minProduct)
		maxProduct = max(curr, curr*maxProduct)

		result = max(result, maxProduct)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
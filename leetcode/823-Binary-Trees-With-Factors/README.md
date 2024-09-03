# [823. Binary Trees With Factors](https://leetcode.com/problems/binary-trees-with-factors/)

## Problem

Given an array of unique integers, `arr`, where each integer `arr[i]` is strictly greater than `1`.

We make a binary tree using these integers, and each number may be used for any number of times. Each non-leaf node's value should be equal to the product of the values of its children.

Return the number of binary trees we can make. The answer may be too large so return the answer modulo `10^9 + 7`.

Example 1:

```
Input: arr = [2,4]
Output: 3
Explanation: We can make these trees: [2], [4], [4, 2, 2]
```

Example 2:

```
Input: arr = [2,4,5,10]
Output: 7
Explanation: We can make these trees: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].
``` 

Constraints:

- `1 <= arr.length <= 1000`
- `2 <= arr[i] <= 10^9`
- All the values of `arr` are **unique**.

## Solution

```go
func numFactoredBinaryTrees(arr []int) int {
	dp := make(map[int]int64)
	mod := int64(1000000007)
	sort.Ints(arr)

	for i, n := range arr {
		dp[n] = 1
		for j := 0; j < i; j++ {
			factor := arr[j]
			if n%factor == 0 {
				quotient := n / factor
				if _, ok := dp[quotient]; ok {
					dp[n] = (dp[n] + int64(dp[factor])*int64(dp[quotient])) % mod
				}
			}
		}
	}

	var count int64
	for _, c := range dp {
		count = (count + c) % mod
	}
	return int(count)
}
```
# [69. Sqrt(x)](https://leetcode.com/problems/sqrtx/)

## Problem

Given a non-negative integer `x`, return the square root of `x` rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

- For example, do not use `pow(x, 0.5)` in c++ or `x ** 0.5` in python.
 

Example 1:

```
Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.
```

Example 2:

```
Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.
``` 

Constraints:

- `0 <= x <= 2^31 - 1`

## Solution

```go
func mySqrt(x int) int {
	left := 0
	right := x/2 + 1

	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left*left > x {
		return left - 1
	}
	return left
}
```
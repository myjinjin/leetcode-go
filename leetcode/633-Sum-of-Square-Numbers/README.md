# [633. Sum of Square Numbers](https://leetcode.com/problems/sum-of-square-numbers/)

## Problem

Given a non-negative integer `c`, decide whether there're two integers `a` and `b` such that `a^2 + b^2 = c`.

 

Example 1:

```
Input: c = 5
Output: true
Explanation: 1 * 1 + 2 * 2 = 5
```

Example 2:

```
Input: c = 3
Output: false
``` 

Constraints:

- `0 <= c <= 2^31 - 1`

## Solution

```go
func judgeSquareSum(c int) bool {
	max := int(math.Sqrt(float64(c)))
	for i := 0; i <= max; i++ {
		a := i
		b := int(math.Sqrt(float64(c - a*a)))
		if a*a+b*b == c {
			return true
		}
	}
	return false
}
```
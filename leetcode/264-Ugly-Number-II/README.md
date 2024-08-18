# [264. Ugly Number II](https://leetcode.com/problems/ugly-number-ii/)

## Problem

An **ugly number** is a positive integer whose prime factors are limited to `2`, `3`, and `5`.

Given an integer `n`, return the `nth` **ugly number**.


Example 1:

```
Input: n = 10
Output: 12
Explanation: [1, 2, 3, 4, 5, 6, 8, 9, 10, 12] is the sequence of the first 10 ugly numbers.
```

Example 2:

```
Input: n = 1
Output: 1
Explanation: 1 has no prime factors, therefore all of its prime factors are limited to 2, 3, and 5.
```

Constraints:

- `1 <= n <= 1690`

## Solution

```go
func nthUglyNumber(n int) int {
	uglyNumber := make([]int, n)
	uglyNumber[0] = 1

	twoIndex, threeIndex, fiveIndex := 0, 0, 0
	two, three, five := 2, 3, 5

	for i := 1; i < n; i++ {
		uglyNumber[i] = min(two, min(three, five))

		if uglyNumber[i] == two {
			twoIndex++
			two = uglyNumber[twoIndex] * 2
		}
		if uglyNumber[i] == three {
			threeIndex++
			three = uglyNumber[threeIndex] * 3
		}
		if uglyNumber[i] == five {
			fiveIndex++
			five = uglyNumber[fiveIndex] * 5
		}
	}

	return uglyNumber[n-1]
}
```
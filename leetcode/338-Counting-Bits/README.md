# [338. Counting Bits](https://leetcode.com/problems/counting-bits/)

## Problem

Given an integer `n`, return an array `ans` of length `n + 1` such that for each `i` (`0 <= i <= n`), `ans[i]` is the number of `1`'s in the binary representation of `i`.

 

Example 1:

```
Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10
```

Example 2:

```
Input: n = 5
Output: [0,1,1,2,1,2]
Explanation:
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101
``` 

Constraints:

- `0 <= n <= 10^5`
 
## Solution

```go
func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		num := i
		for num != 0 {
			result[i] += num & 1
			num >>= 1
		}
	}
	return result
}
```
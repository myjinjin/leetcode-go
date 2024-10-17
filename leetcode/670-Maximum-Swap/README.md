# [670. Maximum Swap](https://leetcode.com/problems/maximum-swap/)

## Problem

You are given an integer `num`. You can swap two digits at most once to get the maximum valued number.

Return the maximum valued number you can get.


Example 1:

```
Input: num = 2736
Output: 7236
Explanation: Swap the number 2 and the number 7.
```

Example 2:

```
Input: num = 9973
Output: 9973
Explanation: No swap.
```

Constraints:

- `0 <= num <= 10^8`

## Solution

```go
func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	n := len(s)

	maxDigits := make([]int, n)
	maxDigits[n-1] = n - 1

	for i := n - 2; i >= 0; i-- {
		if s[i] > s[maxDigits[i+1]] {
			maxDigits[i] = i
		} else {
			maxDigits[i] = maxDigits[i+1]
		}
	}

	for i := 0; i < n; i++ {
		if s[i] < s[maxDigits[i]] {
			s[i], s[maxDigits[i]] = s[maxDigits[i]], s[i]
			break
		}
	}

	result, _ := strconv.Atoi(string(s))
	return result
}
```
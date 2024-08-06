# [1461. Check If a String Contains All Binary Codes of Size K](https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/)

## Problem

Given a binary string `s` and an integer `k`, return `true` if every binary code of length `k` is a substring of `s`. Otherwise, return `false`.

 

Example 1:

```
Input: s = "00110110", k = 2
Output: true
Explanation: The binary codes of length 2 are "00", "01", "10" and "11". They can be all found as substrings at indices 0, 1, 3 and 2 respectively.
```

Example 2:

```
Input: s = "0110", k = 1
Output: true
Explanation: The binary codes of length 1 are "0" and "1", it is clear that both exist as a substring. 
```

Example 3:

```
Input: s = "0110", k = 2
Output: false
Explanation: The binary code "00" is of length 2 and does not exist in the array.
``` 

Constraints:

- `1 <= s.length <= 5 * 10^5`
- `s[i]` is either `'0'` or `'1'`.
- `1 <= k <= 20`

## Solution

```go
func hasAllCodes(s string, k int) bool {
	if len(s) < k {
		return false
	}

	substrs := make(map[string]bool)

	for start := 0; start < len(s)-k+1; start++ {
		curr := s[start : start+k]

		if _, exist := substrs[curr]; !exist {
			substrs[curr] = true
		}
	}

	return len(substrs) == int(math.Pow(2, float64(k)))
}
```
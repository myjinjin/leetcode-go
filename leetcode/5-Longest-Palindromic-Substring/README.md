# [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

## Problem

Given a string `s`, return the longest palindromic substring in `s`.


Example 1:

```
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
```

Example 2:

```
Input: s = "cbbd"
Output: "bb"
``` 

Constraints:

- `1 <= s.length <= 1000`
- `s` consist of only digits and English letters.

## Solution

```go
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLength := 0, 1
	for i := 0; i < len(s); i++ {
		left, right := i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLength {
				start = left
				maxLength = right - left + 1
			}
			left--
			right++
		}

		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLength {
				start = left
				maxLength = right - left + 1
			}
			left--
			right++
		}
	}
	return s[start : start+maxLength]
}
```
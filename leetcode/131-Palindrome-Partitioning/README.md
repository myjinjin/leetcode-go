# [131. Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/)

## Problem

Given a string `s`, partition `s` such that every 
substring of the partition is a palindrome. Return all possible palindrome partitioning of `s`.

 

Example 1:

```
Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]
```

Example 2:

```
Input: s = "a"
Output: [["a"]]
``` 

Constraints:

- `1 <= s.length <= 16`
- `s` contains only lowercase English letters.

## Solution

```go
func partition(s string) [][]string {
	result := [][]string{}
	var backtrack func(start int, curr []string)
	backtrack = func(start int, curr []string) {
		if start == len(s) {
			result = append(result, append([]string{}, curr...))
			return
		}

		for end := start; end < len(s); end++ {
			if isPalindrome(s[start : end+1]) {
				curr = append(curr, s[start:end+1])
				backtrack(end+1, curr)
				curr = curr[:len(curr)-1]
			}
		}
	}
	backtrack(0, []string{})
	return result
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
```
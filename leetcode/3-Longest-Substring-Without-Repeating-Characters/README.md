# [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## Problem

Given a string `s`, find the length of the **longest substring** without repeating characters.


Example 1:

```
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

Example 2:

```
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

Example 3:

```
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
```

Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:

- `0 <= s.length <= 5 * 10^4`
- `s` consists of English letters, digits, symbols and spaces.

## Solution

```go
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	chars := make(map[byte]int)
	result := 0
	start := 0

	for end := 0; end < len(s); end++ {
		if idx, found := chars[s[end]]; found && idx >= start {
			start = idx + 1
		}
		chars[s[end]] = end
		result = max(result, end-start+1)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
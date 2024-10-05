# [567. Permutation in String](https://leetcode.com/problems/permutation-in-string/)

## Problem

Given two strings `s1` and `s2`, return `true` if `s2` contains a permutation of `s1`, or `false` otherwise.

In other words, return `true` if one of `s1`'s permutations is the substring of `s2`.
 

Example 1:

```
Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
```

Example 2:

```
Input: s1 = "ab", s2 = "eidboaoo"
Output: false
```

Constraints:

- `1 <= s1.length, s2.length <= 10^4`
- `s1` and `s2` consist of lowercase English letters.

## Solution

```go
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Count := make(map[rune]int)
	for _, r := range s1 {
		s1Count[r]++
	}

	windowCount := make(map[rune]int)
	for i := 0; i < len(s1); i++ {
		windowCount[rune(s2[i])]++
	}

	if reflect.DeepEqual(s1Count, windowCount) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		windowCount[rune(s2[i-len(s1)])]--
		if windowCount[rune(s2[i-len(s1)])] == 0 {
			delete(windowCount, rune(s2[i-len(s1)]))
		}
		windowCount[rune(s2[i])]++

		if reflect.DeepEqual(s1Count, windowCount) {
			return true
		}
	}
	return false
}
```
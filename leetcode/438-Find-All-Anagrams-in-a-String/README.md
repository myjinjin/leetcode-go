# [438. Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string/)

## Problem

Given two strings `s` and `p`, return an array of all the start indices of `p`'s anagrams in `s`. You may return the answer in **any order**.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

 
Example 1:

```
Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
```

Example 2:

```
Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
```

Constraints:

- `1 <= s.length, p.length <= 3 * 10^4`
- `s` and `p` consist of lowercase English letters.

## Solution

```go
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	chars := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		chars[p[i]]++
	}
	result := []int{}
	windowSize := len(p)
	window := make(map[byte]int)
	for i := 0; i < windowSize; i++ {
		window[s[i]]++
	}
	if reflect.DeepEqual(window, chars) {
		result = append(result, 0)
	}

	for i := 1; i < len(s)-windowSize+1; i++ {
		if _, ok := window[s[i-1]]; ok {
			window[s[i-1]]--
			if window[s[i-1]] == 0 {
				delete(window, s[i-1])
			}
		}
		window[s[i+windowSize-1]]++
		if reflect.DeepEqual(window, chars) {
			result = append(result, i)
		}
	}

	return result
}
```
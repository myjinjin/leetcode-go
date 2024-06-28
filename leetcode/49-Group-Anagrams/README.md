# [49. Group Anagrams](https://leetcode.com/problems/group-anagrams/)

## Problem

Given an array of strings `strs`, group **the anagrams** together. You can return the answer in **any order**.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.


Example 1:

```
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

Example 2:

```
Input: strs = [""]
Output: [[""]]
```

Example 3:

```
Input: strs = ["a"]
Output: [["a"]]
```

Constraints:

- `1 <= strs.length <= 10^4`
- `0 <= strs[i].length <= 100`
- `strs[i]` consists of lowercase English letters.

## Solution

```go
func groupAnagrams(strs []string) [][]string {
	anagrams := [][]string{}
	anagramMap := make(map[string][]string)
	for _, str := range strs {
		sorted := sortString(str)
		anagramMap[sorted] = append(anagramMap[sorted], str)
	}

	for _, value := range anagramMap {
		anagrams = append(anagrams, value)
	}

	return anagrams
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	var b strings.Builder
	b.Grow(len(s))
	for _, ch := range r {
		b.WriteRune(ch)
	}
	return b.String()
}
```
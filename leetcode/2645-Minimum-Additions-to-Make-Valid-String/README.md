# [2645. Minimum Additions to Make Valid String](https://leetcode.com/problems/minimum-additions-to-make-valid-string/)

## Problem

Given a string `word` to which you can insert letters "a", "b" or "c" anywhere and any number of times, return the minimum number of letters that must be inserted so that `word` becomes valid.

A string is called valid if it can be formed by concatenating the string "abc" several times.

Example 1:

```
Input: word = "b"
Output: 2
Explanation: Insert the letter "a" right before "b", and the letter "c" right next to "b" to obtain the valid string "abc".
```

Example 2:

```
Input: word = "aaa"
Output: 6
Explanation: Insert letters "b" and "c" next to each "a" to obtain the valid string "abcabcabc".
```

Example 3:

```
Input: word = "abc"
Output: 0
Explanation: word is already valid. No modifications are needed. 
```

Constraints:

- `1 <= word.length <= 50`
- `word` consists of letters "a", "b" and "c" only. 

## Solution

```go
func addMinimum(word string) int {
	j := 0
	add := 0
	for i := 0; i < len(word); i++ {
		for j%3 != int(word[i]-'a') {
			add++
			j++
		}
		j++
	}

	if j%3 != 0 {
		add += (3 - j%3)
	}
	return add
}
```
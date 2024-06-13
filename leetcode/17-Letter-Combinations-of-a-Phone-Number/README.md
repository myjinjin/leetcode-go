# [17. Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/)

## Problem

Given a string containing digits from `2-9` inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.


 

Example 1:

```
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

Example 2:

```
Input: digits = ""
Output: []
```

Example 3:

```
Input: digits = "2"
Output: ["a","b","c"]
``` 

Constraints:

- `0 <= digits.length <= 4`
- `digits[i]` is a digit in the range `['2', '9']`.

## Solution

```go
func letterCombinations(digits string) []string {
	combinations := []string{}
	phoneMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	if len(digits) == 0 {
		return []string{}
	}

	var backtrack func(input string, index int, curr string, combinations *[]string)
	backtrack = func(input string, index int, curr string, combinations *[]string) {
		if len(input) == index {
			*combinations = append(*combinations, curr)
			return
		}

		for _, char := range phoneMap[input[index]] {
			backtrack(input, index+1, curr+string(char), combinations)
		}
	}

	backtrack(digits, 0, "", &combinations)

	return combinations
}
```
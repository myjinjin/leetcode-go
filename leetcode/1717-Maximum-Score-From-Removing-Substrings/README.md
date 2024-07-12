# [1717. Maximum Score From Removing Substrings](https://leetcode.com/problems/maximum-score-from-removing-substrings/)

## Problem

You are given a string `s` and two integers `x` and `y`. You can perform two types of operations any number of times.

- Remove substring `"ab"` and gain `x` points.
    - For example, when removing `"ab"` from `"cabxbae"` it becomes `"cxbae"`.
- Remove substring `"ba"` and gain `y` points.
    - For example, when removing `"ba"` from `"cabxbae"` it becomes `"cabxe"`.

Return the maximum points you can gain after applying the above operations on `s`.


Example 1:

```
Input: s = "cdbcbbaaabab", x = 4, y = 5
Output: 19
Explanation:
- Remove the "ba" underlined in "cdbcbbaaabab". Now, s = "cdbcbbaaab" and 5 points are added to the score.
- Remove the "ab" underlined in "cdbcbbaaab". Now, s = "cdbcbbaa" and 4 points are added to the score.
- Remove the "ba" underlined in "cdbcbbaa". Now, s = "cdbcba" and 5 points are added to the score.
- Remove the "ba" underlined in "cdbcba". Now, s = "cdbc" and 5 points are added to the score.
Total score = 5 + 4 + 5 + 5 = 19.
```

Example 2:

```
Input: s = "aabbaaxybbaabb", x = 5, y = 4
Output: 20
``` 

Constraints:

- `1 <= s.length <= 10^5`
- `1 <= x, y <= 10^4`
- `s` consists of lowercase English letters.

## Solution

```go
func maximumGain(s string, x int, y int) int {
	removeCombo := func(str string, a, b byte, score int) (int, string) {
		stack := []byte{}
		gain := 0
		for i := 0; i < len(str); i++ {
			if len(stack) > 0 && stack[len(stack)-1] == a && str[i] == b {
				stack = stack[:len(stack)-1]
				gain += score
			} else {
				stack = append(stack, str[i])
			}
		}
		return gain, string(stack)
	}

	var totalGain int
	if x > y {
		gain1, remain := removeCombo(s, 'a', 'b', x)
		gain1Extra, _ := removeCombo(remain, 'b', 'a', y)
		totalGain = gain1 + gain1Extra
	} else {
		gain2, remain2 := removeCombo(s, 'b', 'a', y)
		gain2Extra, _ := removeCombo(remain2, 'a', 'b', x)
		totalGain = gain2 + gain2Extra
	}
	return totalGain
}
```
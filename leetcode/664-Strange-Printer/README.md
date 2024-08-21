# [664. Strange Printer](https://leetcode.com/problems/strange-printer/)

## Problem

There is a strange printer with the following two special properties:

- The printer can only print a sequence of **the same character** each time.
- At each turn, the printer can print new characters starting from and ending at any place and will cover the original existing characters.

Given a string `s`, return the minimum number of turns the printer needed to print it.

 

Example 1:

```
Input: s = "aaabbb"
Output: 2
Explanation: Print "aaa" first and then print "bbb".
```

Example 2:

```
Input: s = "aba"
Output: 2
Explanation: Print "aaa" first and then print "b" from the second place of the string, which will cover the existing character 'a'.
``` 

Constraints:

- `1 <= s.length <= 100`
- `s` consists of lowercase English letters.

## Solution

```go
func strangePrinter(s string) int {
	n := len(s)
	turns := make([][]int, n)
	for i := range turns {
		turns[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		turns[i][i] = 1
	}

	for length := 2; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1

			if length == 2 {
				if s[i] == s[j] {
					turns[i][j] = 1
				} else {
					turns[i][j] = 2
				}
				continue
			}

			turns[i][j] = math.MaxInt32
			for k := i; k < j; k++ {
				turns[i][j] = min(turns[i][j], turns[i][k]+turns[k+1][j])
			}
			if s[i] == s[j] {
				turns[i][j]--
			}
		}
	}

	return turns[0][n-1]
}
```
# [344. Reverse String](https://leetcode.com/problems/reverse-string/)

## Problem

Write a function that reverses a string. The input string is given as an array of characters `s`.

You must do this by modifying the input array in-place with `O(1)` extra memory.

Example 1:

```
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
```

Example 2:

```
Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
```

Constraints:

- `1 <= s.length <= 105`
- `s[i]` is a printable ascii character.

## Solution

```go
func reverseString(s []byte)  {
    for i:=0; i<len(s)/2; i++ {
        s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
    }
}
```
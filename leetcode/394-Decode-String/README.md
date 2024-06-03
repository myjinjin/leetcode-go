# [394. Decode String](https://leetcode.com/problems/decode-string/)

## Problem

Given an encoded string, return its decoded string.

The encoding rule is: `k[encoded_string]`, where the `encoded_string` inside the square brackets is being repeated exactly `k` times. Note that `k` is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, `k`. For example, there will not be input like `3a` or `2[4]`.

The test cases are generated so that the length of the output will never exceed `10^5`.

 

Example 1:

```
Input: s = "3[a]2[bc]"
Output: "aaabcbc"
```

Example 2:

```
Input: s = "3[a2[c]]"
Output: "accaccacc"
```

Example 3:

```
Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"
```

Constraints:

- `1 <= s.length <= 30`
- `s` consists of lowercase English letters, digits, and square brackets `'[]'`.
- `s` is guaranteed to be a valid input.
- All the integers in `s` are in the range `[1, 300]`.

## Solution

```go
func decodeString(s string) string {
    stack := []string{}
    currNum := 0
    currStr := ""

    for _, c := range s {
        if c == '[' {
            stack = append(stack, currStr)
            stack = append(stack, strconv.Itoa(currNum))
            currStr = ""
            currNum = 0
        } else if c == ']' {
            num, _ := strconv.Atoi(stack[len(stack)-1])
            stack = stack[:len(stack)-1]
            prevStr := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            currStr = prevStr + strings.Repeat(currStr, num)
        } else if c>='0' && c<='9' {
            currNum = currNum*10 + int(c-'0')
        } else {
            currStr += string(c)
        }
    }

    return currStr
}
```
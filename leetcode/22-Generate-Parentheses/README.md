# [22. Generate Parentheses](https://leetcode.com/problems/generate-parentheses/)

## Problem

Given `n` pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

 

Example 1:

```
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
```

Example 2:

```
Input: n = 1
Output: ["()"]
``` 

Constraints:

- `1 <= n <= 8`

## Solution

```go
func generateParenthesis(n int) []string {
	result := []string{}
	var dfs func(curr string, open, close int, res *[]string)
	dfs = func(curr string, open, close int, res *[]string) {
		if len(curr) == n*2 {
			*res = append(*res, curr)
			return
		}

		if open == n {
			dfs(curr+")", open, close+1, res)
		} else if open > close {
			dfs(curr+")", open, close+1, res)
			dfs(curr+"(", open+1, close, res)
		} else {
			dfs(curr+"(", open+1, close, res)
		}
	}

	dfs("", 0, 0, &result)
	return result
}
```
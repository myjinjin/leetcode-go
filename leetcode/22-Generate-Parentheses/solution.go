package leetcode

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

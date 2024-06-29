package leetcode

func getAncestors(n int, edges [][]int) [][]int {
	result := make([][]int, n)
	child := make(map[int][]int)

	for _, edge := range edges {
		from, to := edge[0], edge[1]
		child[from] = append(child[from], to)
	}

	var dfs func(ancestor, curr int)
	dfs = func(ancestor, curr int) {
		for _, ch := range child[curr] {
			if !contains(result[ch], ancestor) {
				result[ch] = append(result[ch], ancestor)
				dfs(ancestor, ch)
			}
		}
	}

	for i := 0; i < n; i++ {
		dfs(i, i)
	}

	return result
}

func contains(arr []int, num int) bool {
	for _, n := range arr {
		if n == num {
			return true
		}
	}
	return false
}

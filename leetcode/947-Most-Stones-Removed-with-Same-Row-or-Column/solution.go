package leetcode

func removeStones(stones [][]int) int {
	n := len(stones)
	visited := make(map[int]bool)

	var dfs func(i int) int
	dfs = func(i int) int {
		visited[i] = true
		count := 0
		for j := 0; j < n; j++ {
			if !visited[j] && (stones[j][0] == stones[i][0] || stones[j][1] == stones[i][1]) {
				count += dfs(j) + 1
			}
		}
		return count
	}

	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			count += dfs(i)
		}
	}
	return count
}

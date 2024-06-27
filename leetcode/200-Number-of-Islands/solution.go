package leetcode

func numIslands(grid [][]byte) int {
	islands := 0
	m := len(grid)
	n := len(grid[0])

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] != '1' {
			return
		}

		grid[row][col] = '0'
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				islands++
			}
		}
	}

	return islands
}

package leetcode

func regionsBySlashes(grid []string) int {
	n := len(grid)
	walls := make([][]bool, 4*n)
	for i := range walls {
		walls[i] = make([]bool, 4*n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '/' {
				walls[4*i][4*j+3] = true
				walls[4*i+1][4*j+2] = true
				walls[4*i+2][4*j+1] = true
				walls[4*i+3][4*j] = true
			} else if grid[i][j] == '\\' {
				walls[4*i][4*j] = true
				walls[4*i+1][4*j+1] = true
				walls[4*i+2][4*j+2] = true
				walls[4*i+3][4*j+3] = true
			}
		}
	}

	visited := make([][]bool, 4*n)
	for i := range visited {
		visited[i] = make([]bool, 4*n)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= 4*n || j < 0 || j >= 4*n || walls[i][j] || visited[i][j] {
			return
		}
		visited[i][j] = true
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	regions := 0
	for i := 0; i < 4*n; i++ {
		for j := 0; j < 4*n; j++ {
			if !walls[i][j] && !visited[i][j] {
				dfs(i, j)
				regions++
			}
		}
	}
	return regions
}

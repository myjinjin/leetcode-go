package leetcode

func findMaxFish(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}

		fish := grid[i][j]
		grid[i][j] = 0

		fish += dfs(i+1, j)
		fish += dfs(i-1, j)
		fish += dfs(i, j+1)
		fish += dfs(i, j-1)
		return fish
	}

	maxFish := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				fish := dfs(i, j)
				if maxFish < fish {
					maxFish = fish
				}
			}
		}
	}
	return maxFish
}

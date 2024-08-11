package leetcode

func minDays(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	if isDisconnected(grid) {
		return 0
	}

	for days := 1; days <= m*n; days++ {
		if canDisconnect(grid, days) {
			return days
		}
	}

	return m * n
}

func canDisconnect(grid [][]int, days int) bool {
	m, n := len(grid), len(grid[0])
	cells := make([][2]int, 0, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				cells = append(cells, [2]int{i, j})
			}
		}
	}

	return tryRemoveCells(grid, cells, 0, days, make([][2]int, 0, days))
}

func tryRemoveCells(grid [][]int, cells [][2]int, start, days int, removed [][2]int) bool {
	if len(removed) == days {
		copyGrid := copyGrid(grid)
		for _, cell := range removed {
			copyGrid[cell[0]][cell[1]] = 0
		}
		return isDisconnected(copyGrid)
	}

	for i := start; i < len(cells); i++ {
		removed = append(removed, cells[i])
		if tryRemoveCells(grid, cells, i+1, days, removed) {
			return true
		}
		removed = removed[:len(removed)-1]
	}

	return false
}

func dfs(grid [][]int, i, j int, visited [][]bool) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || visited[i][j] || grid[i][j] == 0 {
		return
	}

	visited[i][j] = true

	dfs(grid, i+1, j, visited)
	dfs(grid, i-1, j, visited)
	dfs(grid, i, j+1, visited)
	dfs(grid, i, j-1, visited)
}

func countIslands(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	islands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				dfs(grid, i, j, visited)
				islands++
			}
		}
	}
	return islands
}

func isDisconnected(grid [][]int) bool {
	island := countIslands(grid)
	return island != 1
}

func copyGrid(grid [][]int) [][]int {
	copied := make([][]int, len(grid))
	for i := range grid {
		copied[i] = make([]int, len(grid[i]))
		copy(copied[i], grid[i])
	}
	return copied
}

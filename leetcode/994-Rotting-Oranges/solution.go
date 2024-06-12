package leetcode

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := [][]int{}
	emptyCount := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j, 0})
			} else if grid[i][j] == 0 {
				emptyCount++
			}
		}
	}
	rottenCount := len(queue)
	freshCount := m*n - rottenCount - emptyCount
	if freshCount == 0 {
		return 0
	}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			currRotten := queue[0]
			queue = queue[1:]
			row, col := currRotten[0], currRotten[1]
			path := currRotten[2]

			for _, d := range directions {
				newRow := row + d[0]
				newCol := col + d[1]

				if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && !visited[newRow][newCol] && grid[newRow][newCol] == 1 {
					queue = append(queue, []int{newRow, newCol, path + 1})
					visited[newRow][newCol] = true
					rottenCount++
					if rottenCount == m*n-emptyCount {
						return path + 1
					}
				}
			}
		}
	}

	return -1
}

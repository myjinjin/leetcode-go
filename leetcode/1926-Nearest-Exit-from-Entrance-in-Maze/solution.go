package leetcode

func nearestExit(maze [][]byte, entrance []int) int {
	m := len(maze)
	n := len(maze[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[entrance[0]][entrance[1]] = true

	queue := [][]int{{entrance[0], entrance[1], 0}}
	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		row, col, path := curr[0], curr[1], curr[2]
		if (row == 0 || row == m-1 || col == 0 || col == n-1) && !(row == entrance[0] && col == entrance[1]) {
			return path
		}

		for _, d := range directions {
			newRow := row + d[0]
			newCol := col + d[1]
			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && !visited[newRow][newCol] && maze[newRow][newCol] != '+' {
				queue = append(queue, []int{newRow, newCol, path + 1})
				visited[newRow][newCol] = true
			}
		}
	}

	return -1
}

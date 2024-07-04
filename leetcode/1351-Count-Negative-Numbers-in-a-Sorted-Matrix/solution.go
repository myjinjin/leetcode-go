package leetcode

func countNegatives(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	row := 0
	col := n - 1

	count := 0
	for row < m && col >= 0 {
		if grid[row][col] < 0 {
			count += m - row
			col--
		} else {
			row++
		}
	}
	return count
}

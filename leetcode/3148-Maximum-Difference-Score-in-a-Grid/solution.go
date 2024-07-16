package leetcode

import "math"

func maxScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var maxScoreVal = math.MinInt32
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			minVal := math.MaxInt32
			if i == 0 && j == 0 {
				continue
			}
			if i != 0 {
				minVal = min(minVal, grid[i-1][j])
			}
			if j != 0 {
				minVal = min(minVal, grid[i][j-1])
			}

			maxScoreVal = max(maxScoreVal, grid[i][j]-minVal)
			grid[i][j] = min(grid[i][j], minVal)
		}
	}
	return maxScoreVal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

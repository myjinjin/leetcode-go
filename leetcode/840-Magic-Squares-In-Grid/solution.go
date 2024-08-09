package leetcode

func numMagicSquaresInside(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var result int

	isMagic := func(i, j int) bool {
		exist := make(map[int]bool)
		for r := i; r < i+3; r++ {
			for c := j; c < j+3; c++ {
				if grid[r][c] < 1 || grid[r][c] > 9 {
					return false
				}
				if exist[grid[r][c]] {
					return false
				}
				exist[grid[r][c]] = true
			}
		}

		for r := i; r < i+3; r++ {
			if grid[r][j]+grid[r][j+1]+grid[r][j+2] != 15 {
				return false
			}
		}

		for c := j; c < j+3; c++ {
			if grid[i][c]+grid[i+1][c]+grid[i+2][c] != 15 {
				return false
			}
		}

		if grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] != 15 {
			return false
		}
		if grid[i][j+2]+grid[i+1][j+1]+grid[i+2][j] != 15 {
			return false
		}

		return true
	}

	for i := 0; i <= m-3; i++ {
		for j := 0; j <= n-3; j++ {
			if isMagic(i, j) {
				result++
			}
		}
	}
	return result
}

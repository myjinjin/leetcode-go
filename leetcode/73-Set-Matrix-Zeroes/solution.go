package leetcode

func setZeroes(matrix [][]int) {
	zeroIndex := make([][2]int, 0)

	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeroIndex = append(zeroIndex, [2]int{i, j})
			}
		}
	}

	for _, indexes := range zeroIndex {
		r := indexes[0]
		c := indexes[1]

		for j := 0; j < n; j++ {
			matrix[r][j] = 0
		}
		for i := 0; i < m; i++ {
			matrix[i][c] = 0
		}
	}
}

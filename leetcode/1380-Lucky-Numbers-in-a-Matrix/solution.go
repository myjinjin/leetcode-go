package leetcode

import "math"

func luckyNumbers(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	rowMin := make([]int, m)
	colMax := make([]int, n)

	for i := range rowMin {
		rowMin[i] = math.MaxInt32
		for j := 0; j < n; j++ {
			if matrix[i][j] < rowMin[i] {
				rowMin[i] = matrix[i][j]
			}
		}
	}

	for j := range colMax {
		colMax[j] = math.MinInt32
		for i := 0; i < m; i++ {
			if matrix[i][j] > colMax[j] {
				colMax[j] = matrix[i][j]
			}
		}
	}

	result := make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == rowMin[i] && matrix[i][j] == colMax[j] {
				result = append(result, matrix[i][j])
			}
		}
	}

	return result
}

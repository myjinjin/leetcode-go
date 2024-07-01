package leetcode

func spiralOrder(matrix [][]int) []int {
	result := []int{}

	rowBegin := 0
	colBegin := 0
	rowEnd := len(matrix) - 1
	colEnd := len(matrix[0]) - 1

	for rowBegin <= rowEnd && colBegin <= colEnd {
		for j := colBegin; j <= colEnd; j++ {
			result = append(result, matrix[rowBegin][j])
		}
		rowBegin++

		for i := rowBegin; i <= rowEnd; i++ {
			result = append(result, matrix[i][colEnd])
		}
		colEnd--

		if rowBegin <= rowEnd {
			for j := colEnd; j >= colBegin; j-- {
				result = append(result, matrix[rowEnd][j])
			}
		}
		rowEnd--

		if colBegin <= colEnd {
			for i := rowEnd; i >= rowBegin; i-- {
				result = append(result, matrix[i][colBegin])
			}
		}
		colBegin++
	}

	return result
}

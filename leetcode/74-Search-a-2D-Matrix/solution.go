package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	left, right := 0, m*n-1
	for left < right {
		mid := (left + right) / 2

		row := mid / n
		col := mid % n
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	row := left / n
	col := left % n

	return matrix[row][col] == target
}

package leetcode

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	arr := make([][]int, m)
	for i := range arr {
		arr[i] = make([]int, n)
	}

	index := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr[i][j] = original[index]
			index++
		}
	}

	return arr
}

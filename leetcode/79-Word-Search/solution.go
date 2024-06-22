package leetcode

func exist(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])
	directions := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	var dfs func(i, j int, index int) bool
	dfs = func(i, j int, index int) bool {
		if index == len(word) {
			return true
		}

		if i < 0 || i >= row || j < 0 || j >= col || board[i][j] != word[index] {
			return false
		}

		temp := board[i][j]
		board[i][j] = '.'

		for _, d := range directions {
			r, c := i+d[0], j+d[1]
			if dfs(r, c, index+1) {
				return true
			}
		}

		board[i][j] = temp
		return false
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

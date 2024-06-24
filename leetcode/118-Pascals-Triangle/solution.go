package leetcode

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	dp[0] = []int{1}
	if numRows == 1 {
		return dp
	}
	dp[1] = []int{1, 1}
	if numRows == 2 {
		return dp
	}

	for i := 2; i < numRows; i++ {
		dp[i] = make([]int, i+1)
		dp[i][0] = 1
		dp[i][i] = 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
	return dp
}

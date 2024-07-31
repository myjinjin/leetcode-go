package leetcode

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 0; i < n; i++ {
		book := books[i]
		width, height := book[0], book[1]
		dp[i+1] = dp[i] + height

		for j := i - 1; j >= 0 && width+books[j][0] <= shelfWidth; j-- {
			height = max(height, books[j][1])
			width += books[j][0]
			dp[i+1] = min(dp[i+1], dp[j]+height)
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

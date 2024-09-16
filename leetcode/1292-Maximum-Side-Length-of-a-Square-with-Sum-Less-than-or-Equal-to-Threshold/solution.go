package leetcode

func maxSideLength(mat [][]int, threshold int) int {
	m := len(mat)
	n := len(mat[0])

	prefixSum := make([][]int, m+1)
	for i := range prefixSum {
		prefixSum[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			prefixSum[i][j] = prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1] + mat[i-1][j-1]
		}
	}

	left, right := 0, min(m, n)
	result := 0

	for left <= right {
		mid := left + (right-left)/2
		if isValid(prefixSum, threshold, mid) {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func isValid(prefixSum [][]int, threshold int, length int) bool {
	for i := length; i < len(prefixSum); i++ {
		for j := length; j < len(prefixSum[0]); j++ {
			sum := prefixSum[i][j] - prefixSum[i-length][j] - prefixSum[i][j-length] + prefixSum[i-length][j-length]
			if sum <= threshold {
				return true
			}
		}
	}
	return false
}

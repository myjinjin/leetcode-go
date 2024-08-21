package leetcode

import "math"

func strangePrinter(s string) int {
	n := len(s)
	turns := make([][]int, n)
	for i := range turns {
		turns[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		turns[i][i] = 1
	}

	for length := 2; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1

			if length == 2 {
				if s[i] == s[j] {
					turns[i][j] = 1
				} else {
					turns[i][j] = 2
				}
				continue
			}

			turns[i][j] = math.MaxInt32
			for k := i; k < j; k++ {
				turns[i][j] = min(turns[i][j], turns[i][k]+turns[k+1][j])
			}
			if s[i] == s[j] {
				turns[i][j]--
			}
		}
	}

	return turns[0][n-1]
}

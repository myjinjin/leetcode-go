package leetcode

func stoneGameII(piles []int) int {
	n := len(piles)
	suffixSums := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffixSums[i] = suffixSums[i+1] + piles[i]
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}

	var maxGet func(int, int) int
	maxGet = func(m, curr int) int {
		if curr >= n {
			return 0
		}

		if curr+2*m >= n {
			return suffixSums[curr]
		}

		if memo[curr][m] != 0 {
			return memo[curr][m]
		}

		stone := 0
		for x := 1; x <= 2*m; x++ {
			currentStones := suffixSums[curr] - maxGet(max(m, x), curr+x)
			stone = max(stone, currentStones)
		}

		memo[curr][m] = stone
		return stone
	}

	return maxGet(1, 0)
}

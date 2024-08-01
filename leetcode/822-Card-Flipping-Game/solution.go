package leetcode

import "math"

func flipgame(fronts []int, backs []int) int {
	badNumber := make(map[int]bool)
	n := len(fronts)
	for i := 0; i < n; i++ {
		if fronts[i] == backs[i] {
			badNumber[fronts[i]] = true
		}
	}

	goodNumber := math.MaxInt32

	for i := 0; i < n; i++ {
		if _, ok := badNumber[fronts[i]]; !ok {
			goodNumber = min(goodNumber, fronts[i])
		}
		if _, ok := badNumber[backs[i]]; !ok {
			goodNumber = min(goodNumber, backs[i])
		}
	}

	if goodNumber == math.MaxInt32 {
		return 0
	}
	return goodNumber
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

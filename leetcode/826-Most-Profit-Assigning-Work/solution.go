package leetcode

import (
	"sort"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	difficultyWithProfits := make([]difficultyWithProfit, len(difficulty))
	for i := 0; i < len(difficulty); i++ {
		difficultyWithProfits[i] = difficultyWithProfit{
			difficulty: difficulty[i],
			profit:     profit[i],
		}
	}

	sort.Slice(difficultyWithProfits, func(i, j int) bool {
		return difficultyWithProfits[i].difficulty < difficultyWithProfits[j].difficulty
	})

	sort.Ints(worker)

	totalProfit := 0
	maxProfit := 0
	j := 0
	for i := 0; i < len(worker); i++ {
		for j < len(difficultyWithProfits) && difficultyWithProfits[j].difficulty <= worker[i] {
			maxProfit = max(maxProfit, difficultyWithProfits[j].profit)
			j++
		}
		totalProfit += maxProfit
	}

	return totalProfit
}

type difficultyWithProfit struct {
	difficulty int
	profit     int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

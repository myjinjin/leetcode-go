package leetcode

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)

	nextNumMustBe := 0
	minIncrement := 0

	for _, num := range nums {
		if nextNumMustBe < num {
			nextNumMustBe = num
		}
		minIncrement += nextNumMustBe - num
		nextNumMustBe++
	}

	return minIncrement
}

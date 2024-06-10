package leetcode

import "sort"

func heightChecker(heights []int) int {
	notMatchCount := 0

	expected := make([]int, len(heights))
	copy(expected, heights)

	sort.Ints(expected)

	for i := 0; i < len(heights); i++ {
		if expected[i] != heights[i] {
			notMatchCount++
		}
	}

	return notMatchCount
}

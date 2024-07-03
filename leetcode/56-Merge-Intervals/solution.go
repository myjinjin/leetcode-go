package leetcode

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	merged := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prevBegin := intervals[0][0]
	prevEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		currBegin := intervals[i][0]
		currEnd := intervals[i][1]
		if currBegin <= prevEnd {
			prevEnd = max(prevEnd, currEnd)
		} else {
			merged = append(merged, []int{prevBegin, prevEnd})
			prevBegin = currBegin
			prevEnd = currEnd
		}
	}

	merged = append(merged, []int{prevBegin, prevEnd})
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

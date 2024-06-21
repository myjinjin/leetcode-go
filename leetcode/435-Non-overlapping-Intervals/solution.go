package leetcode

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	result := 0
	n := len(intervals)
	prevIdx := 0
	tmpPrevInterval := 0
	currIdx := 1
	for prevIdx < n-1 && currIdx < n {
		prev := intervals[prevIdx]
		curr := intervals[currIdx]
		if curr[0] < prev[1] {
			result++
			currIdx++
			tmpPrevInterval++
		} else {
			currIdx++
			prevIdx += tmpPrevInterval + 1
			tmpPrevInterval = 0
		}
	}
	return result
}

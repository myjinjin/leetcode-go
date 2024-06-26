package leetcode

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 1 {
		return 1
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	n := len(points)
	arrowShots := 1
	lastPos := points[0][1]
	for i := 1; i < n; i++ {
		if points[i][0] > lastPos {
			arrowShots++
			lastPos = points[i][1]
		}
	}
	return arrowShots
}

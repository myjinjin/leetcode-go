package leetcode

import (
	"strconv"
	"strings"
)

func equalPairs(grid [][]int) int {
	n := len(grid)
	rowStrings := make(map[string]int)

	var sb strings.Builder
	for row := 0; row < n; row++ {
		sb.Reset()
		for col := 0; col < n; col++ {
			if col > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(strconv.Itoa(grid[row][col]))
		}
		rowStrings[sb.String()]++
	}

	count := 0
	for col := 0; col < n; col++ {
		sb.Reset()
		for row := 0; row < n; row++ {
			if row > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(strconv.Itoa(grid[row][col]))
		}
		if v, ok := rowStrings[sb.String()]; ok {
			count += v
		}
	}
	return count
}

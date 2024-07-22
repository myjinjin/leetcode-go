package leetcode

import "sort"

func sortPeople(names []string, heights []int) []string {
	orderedHeight := make([][2]int, len(heights))
	for i := 0; i < len(heights); i++ {
		orderedHeight[i] = [2]int{i, heights[i]}
	}

	sort.Slice(orderedHeight, func(i, j int) bool {
		return orderedHeight[i][1] > orderedHeight[j][1]
	})

	result := make([]string, len(names))
	for i := 0; i < len(result); i++ {
		result[i] = names[orderedHeight[i][0]]
	}
	return result
}

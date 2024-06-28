package leetcode

import "sort"

func maximumImportance(n int, roads [][]int) int64 {
	degree := make([]int, n)
	for _, road := range roads {
		degree[road[0]]++
		degree[road[1]]++
	}

	cities := make([]int, n)
	for i := range cities {
		cities[i] = i
	}

	sort.Slice(cities, func(i, j int) bool {
		return degree[cities[i]] < degree[cities[j]]
	})

	importance := make([]int, n)
	for i, city := range cities {
		importance[city] = i + 1
	}

	var totalImportance int64
	for _, road := range roads {
		totalImportance += int64(importance[road[0]] + importance[road[1]])
	}
	return totalImportance
}

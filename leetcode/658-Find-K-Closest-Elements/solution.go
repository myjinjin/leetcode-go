package leetcode

import "sort"

func findClosestElements(arr []int, k int, x int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if abs(x, arr[i]) == abs(x, arr[j]) {
			return arr[i] < arr[j]
		}
		return abs(x, arr[i]) < abs(x, arr[j])
	})

	result := arr[:k]
	sort.Ints(result)
	return result
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}

package leetcode

import "sort"

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	left, right := 1, position[len(position)-1]-position[0]
	result := -1
	for left <= right {
		mid := (left + right) / 2
		lastPosition, balls := position[0], 1
		for i := 1; i < len(position); i++ {
			if position[i]-lastPosition >= mid {
				lastPosition = position[i]
				balls++
			}
		}
		if balls >= m {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

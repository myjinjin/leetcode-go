package leetcode

import "sort"

func dividePlayers(skill []int) int64 {
	n := len(skill)
	sum := 0
	for i := 0; i < n; i++ {
		sum += skill[i]
	}
	if sum%(n/2) != 0 {
		return -1
	}

	sumTarget := sum / (n / 2)

	sort.Ints(skill)
	left, right := 0, n-1

	result := int64(0)
	for left < right {
		if skill[left]+skill[right] != sumTarget {
			return -1
		}
		result += int64(skill[left]) * int64(skill[right])
		left++
		right--
	}

	return result
}

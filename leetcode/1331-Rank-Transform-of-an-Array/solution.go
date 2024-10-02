package leetcode

import "sort"

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	rankMap := make(map[int]int)
	rank := 1
	for _, num := range sortedArr {
		if _, exist := rankMap[num]; !exist {
			rankMap[num] = rank
			rank++
		}
	}

	for i, num := range arr {
		arr[i] = rankMap[num]
	}

	return arr
}

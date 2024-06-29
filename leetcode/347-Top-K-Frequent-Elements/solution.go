package leetcode

import "sort"

func topKFrequent(nums []int, k int) []int {
	result := make([]int, k)
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	numCount := make([][2]int, 0, len(numMap))
	for num, count := range numMap {
		numCount = append(numCount, [2]int{num, count})
	}

	sort.Slice(numCount, func(i, j int) bool {
		return numCount[i][1] > numCount[j][1]
	})

	for i := 0; i < k; i++ {
		result[i] = numCount[i][0]
	}
	return result
}

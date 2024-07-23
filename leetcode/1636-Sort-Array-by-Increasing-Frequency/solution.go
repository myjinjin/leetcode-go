package leetcode

import "sort"

func frequencySort(nums []int) []int {
	numCount := make(map[int]int)
	for i := range nums {
		numCount[nums[i]]++
	}

	numsWithCount := make([][2]int, 0, len(numCount))
	for k, v := range numCount {
		numsWithCount = append(numsWithCount, [2]int{k, v})
	}

	sort.Slice(numsWithCount, func(i, j int) bool {
		if numsWithCount[i][1] < numsWithCount[j][1] {
			return true
		}
		if numsWithCount[i][1] > numsWithCount[j][1] {
			return false
		}
		return numsWithCount[i][0] > numsWithCount[j][0]
	})

	result := make([]int, 0, len(nums))
	for _, v := range numsWithCount {
		value := v[0]
		count := v[1]
		for i := 0; i < count; i++ {
			result = append(result, value)
		}
	}
	return result
}

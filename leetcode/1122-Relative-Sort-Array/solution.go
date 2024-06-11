package leetcode

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	result := make([]int, 0, len(arr1))

	arr1Counter := make(map[int]int)
	for _, n := range arr1 {
		arr1Counter[n]++
	}

	for _, n := range arr2 {
		if v, ok := arr1Counter[n]; ok {
			for i := 0; i < v; i++ {
				result = append(result, n)
			}
			delete(arr1Counter, n)
		}
	}

	remainNums := make([]int, 0)
	for num, count := range arr1Counter {
		for i := 0; i < count; i++ {
			remainNums = append(remainNums, num)
		}
	}

	sort.Ints(remainNums)

	result = append(result, remainNums...)
	return result
}

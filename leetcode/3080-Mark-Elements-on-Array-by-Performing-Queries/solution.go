package leetcode

import "sort"

func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	n := len(nums)
	numsWithIndex := make([][2]int, n)
	for i, num := range nums {
		numsWithIndex[i] = [2]int{num, i}
	}

	sort.Slice(numsWithIndex, func(i, j int) bool {
		if numsWithIndex[i][0] != numsWithIndex[j][0] {
			return numsWithIndex[i][0] < numsWithIndex[j][0]
		}
		return numsWithIndex[i][1] < numsWithIndex[j][1]
	})

	unmarked := make(map[int]bool)
	sum := int64(0)
	for i, num := range nums {
		unmarked[i] = true
		sum += int64(num)
	}

	result := make([]int64, len(queries))
	unmarkedSorted := 0

	for i, query := range queries {
		index, k := query[0], query[1]

		if unmarked[index] {
			unmarked[index] = false
			sum -= int64(nums[index])
		}

		for unmarkedSorted < n && k > 0 {
			idx := numsWithIndex[unmarkedSorted][1]
			if unmarked[idx] {
				unmarked[idx] = false
				sum -= int64(nums[idx])
				k--
			}
			unmarkedSorted++
		}

		result[i] = sum
	}

	return result
}

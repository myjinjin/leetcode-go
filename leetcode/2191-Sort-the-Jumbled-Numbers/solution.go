package leetcode

import "sort"

func sortJumbled(mapping []int, nums []int) []int {
	type pair struct {
		original int
		mapped   int
		index    int
	}

	getMappedValue := func(num int) int {
		if num == 0 {
			return mapping[0]
		}
		var mapped int
		for i, multiplier := 1, 1; num > 0; i, num = i*10, num/10 {
			digit := num % 10
			mapped += mapping[digit] * multiplier
			multiplier *= 10
		}
		return mapped
	}

	pairs := make([]pair, len(nums))
	for i, num := range nums {
		pairs[i] = pair{num, getMappedValue(num), i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].mapped == pairs[j].mapped {
			return pairs[i].index < pairs[j].index
		}
		return pairs[i].mapped < pairs[j].mapped
	})

	result := make([]int, len(nums))
	for i, p := range pairs {
		result[i] = p.original
	}

	return result
}

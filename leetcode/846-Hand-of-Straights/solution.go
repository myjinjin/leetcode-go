package leetcode

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	numbers := make(map[int]int)
	for _, n := range hand {
		numbers[n]++
	}

	keys := []int{}
	for k := range numbers {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, key := range keys {
		count := numbers[key]
		if count > 0 {
			numbers[key] -= count
			for i := 1; i < groupSize; i++ {
				if numbers[i+key] >= count {
					numbers[i+key] -= count
				} else {
					return false
				}
			}
		}
	}

	return true
}

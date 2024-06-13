package leetcode

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	successful := []int{}
	sort.Ints(potions)

	for _, spell := range spells {
		left, right := 0, len(potions)

		for left < right {
			mid := (left + right) / 2
			if int64(potions[mid])*int64(spell) < success {
				left = mid + 1
			} else {
				right = mid
			}
		}

		successful = append(successful, len(potions)-left)
	}
	return successful
}

package leetcode

import "sort"

func minimumPushes(word string) int {
	freq := make(map[rune]int)
	for _, r := range word {
		freq[r]++
	}

	runes := make([]rune, 0, len(freq))
	for r := range freq {
		runes = append(runes, r)
	}

	sort.Slice(runes, func(i, j int) bool {
		return freq[runes[i]] > freq[runes[j]]
	})

	pushes := 0
	for i, r := range runes {
		pushes += freq[r] * ((i / 8) + 1)
	}

	return pushes
}

package leetcode

import "strings"

func findWordsContaining(words []string, x byte) []int {
	result := make([]int, 0, len(words))

	for i, word := range words {
		if strings.ContainsAny(word, string(x)) {
			result = append(result, i)
		}
	}

	return result
}

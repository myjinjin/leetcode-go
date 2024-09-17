package leetcode

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	words := make(map[string]int)

	ones := strings.Split(s1, " ")
	twos := strings.Split(s2, " ")

	for _, word := range ones {
		words[word]++
	}
	for _, word := range twos {
		words[word]++
	}

	result := []string{}
	for k, v := range words {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}

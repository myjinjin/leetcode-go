package leetcode

import "strings"

func commonChars(words []string) []string {
	wordCount := make([]int, 26)
	for _, c := range words[0] {
		wordCount[int(c-'a')]++
	}

	for i := 1; i < len(words); i++ {
		word := words[i]
		tmpCount := make([]int, 26)
		for _, c := range word {
			tmpCount[int(c-'a')]++
		}
		for i := 0; i < len(wordCount); i++ {
			wordCount[i] = min(wordCount[i], tmpCount[i])
		}
	}

	var sb strings.Builder
	for i := 0; i < len(wordCount); i++ {
		count := wordCount[i]
		for count > 0 {
			sb.WriteByte(byte(i + int('a')))
			count--
		}
	}
	return strings.Split(sb.String(), "")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

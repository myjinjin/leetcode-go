package leetcode

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words1 := strings.Split(sentence1, " ")
	words2 := strings.Split(sentence2, " ")

	if len(words1) < len(words2) {
		words1, words2 = words2, words1
	}

	for len(words1) > 0 && len(words2) > 0 && words1[0] == words2[0] {
		words1 = words1[1:]
		words2 = words2[1:]
	}

	for len(words1) > 0 && len(words2) > 0 && words1[len(words1)-1] == words2[len(words2)-1] {
		words1 = words1[:len(words1)-1]
		words2 = words2[:len(words2)-1]
	}

	return len(words2) == 0
}

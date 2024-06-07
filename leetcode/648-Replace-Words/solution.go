package leetcode

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	dict := make(map[string]bool)
	for _, word := range dictionary {
		dict[word] = true
	}

	var sb strings.Builder
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if i != 0 {
			sb.WriteString(" ")
		}
		searchWord := make([]byte, 0, len(word))
		for i := 0; i < len(word); i++ {
			searchWord = append(searchWord, word[i])
			if _, ok := dict[string(searchWord)]; ok {
				break
			}
		}
		sb.Write(searchWord)
	}

	return sb.String()
}

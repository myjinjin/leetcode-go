package leetcode

func countConsistentStrings(allowed string, words []string) int {
	allowedChars := make(map[rune]bool)
	for _, c := range allowed {
		allowedChars[c] = true
	}

	count := 0
	for _, word := range words {
		isConsistent := true
		for _, c := range word {
			if _, exist := allowedChars[c]; !exist {
				isConsistent = false
				break
			}
		}
		if isConsistent {
			count++
		}
	}

	return count
}

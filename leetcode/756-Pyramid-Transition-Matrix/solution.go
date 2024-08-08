package leetcode

func pyramidTransition(bottom string, allowed []string) bool {
	validPatterns := make(map[string][]byte)
	for _, s := range allowed {
		key := s[:2]
		validPatterns[key] = append(validPatterns[key], s[2])
	}

	type PyramidState struct {
		currRow           string
		nextRowInProgress string
	}

	cachedResults := make(map[PyramidState]bool)

	var buildNextRow func(currRow string, nextRowInProgress []byte) bool
	buildNextRow = func(currRow string, nextRowInProgress []byte) bool {
		if len(currRow) == 1 {
			return true
		}
		if len(nextRowInProgress)+1 == len(currRow) {
			return buildNextRow(string(nextRowInProgress), []byte{})
		}

		state := PyramidState{currRow, string(nextRowInProgress)}
		if result, exists := cachedResults[state]; exists {
			return result
		}

		nextBlockIndex := len(nextRowInProgress)
		for _, topBlock := range validPatterns[currRow[nextBlockIndex:nextBlockIndex+2]] {
			if buildNextRow(currRow, append(nextRowInProgress, topBlock)) {
				cachedResults[state] = true
				return true
			}
		}

		cachedResults[state] = false
		return false
	}

	return buildNextRow(bottom, []byte{})
}

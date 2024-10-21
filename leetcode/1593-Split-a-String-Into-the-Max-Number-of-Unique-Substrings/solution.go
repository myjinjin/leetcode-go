package leetcode

func maxUniqueSplit(s string) int {
	usedSubstrings := make(map[string]bool)

	var backtrack func(start int) (count int)
	backtrack = func(start int) (count int) {
		if start == len(s) {
			return 0
		}

		maxCount := 0

		for end := start + 1; end <= len(s); end++ {
			substring := s[start:end]
			if !usedSubstrings[substring] {
				usedSubstrings[substring] = true
				count := 1 + backtrack(end)
				maxCount = max(maxCount, count)
				usedSubstrings[substring] = false
			}
		}

		return maxCount
	}

	return backtrack(0)
}

package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	chars := make(map[byte]int)
	result := 0
	start := 0

	for end := 0; end < len(s); end++ {
		if idx, found := chars[s[end]]; found && idx >= start {
			start = idx + 1
		}
		chars[s[end]] = end
		result = max(result, end-start+1)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

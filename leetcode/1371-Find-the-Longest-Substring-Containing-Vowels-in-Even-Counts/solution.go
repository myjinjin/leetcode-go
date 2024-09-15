package leetcode

func findTheLongestSubstring(s string) int {
	vowels := map[byte]int{'a': 0, 'e': 1, 'i': 2, 'o': 3, 'u': 4}
	seen := make(map[int]int)
	seen[0] = -1
	result, status := 0, 0

	for i := 0; i < len(s); i++ {
		if bit, exists := vowels[s[i]]; exists {
			status ^= 1 << bit
		}
		if firstIndex, exists := seen[status]; exists {
			result = max(result, i-firstIndex)
		} else {
			seen[status] = i
		}
	}

	return result
}

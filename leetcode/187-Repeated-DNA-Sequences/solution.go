package leetcode

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}

	seen := make(map[string]int)
	result := make(map[string]bool)

	for i := 0; i <= len(s)-10; i++ {
		substr := s[i : i+10]
		seen[substr]++
		if seen[substr] > 1 {
			result[substr] = true
		}
	}

	var repeated []string
	for seq := range result {
		repeated = append(repeated, seq)
	}
	return repeated
}

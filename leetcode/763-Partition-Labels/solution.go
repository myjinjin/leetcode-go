package leetcode

func partitionLabels(s string) []int {
	lastIndex := make(map[rune]int)
	for i, ch := range s {
		lastIndex[ch] = i
	}

	result := []int{}
	start, end := 0, 0
	for i, ch := range s {
		if lastIndex[ch] > end {
			end = lastIndex[ch]
		}
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}

	return result
}

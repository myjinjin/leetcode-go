package leetcode

import "strings"

func minLength(s string) int {
	for strings.Count(s, "AB") != 0 || strings.Count(s, "CD") != 0 {
		s = strings.ReplaceAll(s, "AB", "")
		s = strings.ReplaceAll(s, "CD", "")
	}

	return len(s)
}

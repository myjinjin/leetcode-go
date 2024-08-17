package leetcode

import "strings"

func makeStringsEqual(s string, target string) bool {
	sCount := strings.Count(s, "1")
	tCount := strings.Count(target, "1")
	if tCount > 0 && sCount == 0 {
		return false
	}

	if tCount == 0 && sCount > 0 {
		return false
	}

	return true
}

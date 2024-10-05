package leetcode

import "reflect"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Count := make(map[rune]int)
	for _, r := range s1 {
		s1Count[r]++
	}

	windowCount := make(map[rune]int)
	for i := 0; i < len(s1); i++ {
		windowCount[rune(s2[i])]++
	}

	if reflect.DeepEqual(s1Count, windowCount) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		windowCount[rune(s2[i-len(s1)])]--
		if windowCount[rune(s2[i-len(s1)])] == 0 {
			delete(windowCount, rune(s2[i-len(s1)]))
		}
		windowCount[rune(s2[i])]++

		if reflect.DeepEqual(s1Count, windowCount) {
			return true
		}
	}
	return false
}

package leetcode

import "reflect"

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	chars := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		chars[p[i]]++
	}
	result := []int{}
	windowSize := len(p)
	window := make(map[byte]int)
	for i := 0; i < windowSize; i++ {
		window[s[i]]++
	}
	if reflect.DeepEqual(window, chars) {
		result = append(result, 0)
	}

	for i := 1; i < len(s)-windowSize+1; i++ {
		if _, ok := window[s[i-1]]; ok {
			window[s[i-1]]--
			if window[s[i-1]] == 0 {
				delete(window, s[i-1])
			}
		}
		window[s[i+windowSize-1]]++
		if reflect.DeepEqual(window, chars) {
			result = append(result, i)
		}
	}

	return result
}

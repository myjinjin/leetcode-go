package leetcode

import "unicode"

func camelMatch(queries []string, pattern string) []bool {
	result := make([]bool, len(queries))

	for i, query := range queries {
		result[i] = matchPattern(query, pattern)
	}

	return result
}

func matchPattern(query, pattern string) bool {
	j := 0
	for _, c := range query {
		if j < len(pattern) && c == rune(pattern[j]) {
			j++
		} else if unicode.IsUpper(c) {
			return false
		}
	}

	return j == len(pattern)
}

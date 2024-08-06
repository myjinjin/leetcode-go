package leetcode

import "math"

func hasAllCodes(s string, k int) bool {
	if len(s) < k {
		return false
	}

	substrs := make(map[string]bool)

	for start := 0; start < len(s)-k+1; start++ {
		curr := s[start : start+k]

		if _, exist := substrs[curr]; !exist {
			substrs[curr] = true
		}
	}

	return len(substrs) == int(math.Pow(2, float64(k)))
}

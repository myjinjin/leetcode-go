package leetcode

func minimumChairs(s string) int {
	n := len(s)
	waiting := 0
	maxWaiting := 0
	for i := 0; i < n; i++ {
		if s[i] == 'E' {
			waiting++
		} else if s[i] == 'L' {
			waiting--
		}
		maxWaiting = max(maxWaiting, waiting)
	}

	return maxWaiting
}

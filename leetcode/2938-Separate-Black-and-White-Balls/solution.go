package leetcode

func minimumSteps(s string) int64 {
	var steps int64
	left := 0
	for right, c := range s {
		if c == '0' {
			steps += int64(right - left)
			left++
		}
	}

	return steps
}

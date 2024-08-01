package leetcode

import "strconv"

func countSeniors(details []string) int {
	senior := 0

	for _, d := range details {
		ageStr := d[11:13]
		age, _ := strconv.Atoi(ageStr)
		if age > 60 {
			senior++
		}
	}

	return senior
}

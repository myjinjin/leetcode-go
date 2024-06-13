package leetcode

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	count := 0
	for i := range students {
		if students[i] == seats[i] {
			continue
		} else {
			count += abs(students[i] - seats[i])
		}
	}
	return count
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

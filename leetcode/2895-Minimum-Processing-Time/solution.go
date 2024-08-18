package leetcode

import (
	"sort"
)

func minProcessingTime(processorTime []int, tasks []int) int {
	sort.Ints(processorTime)
	sort.Sort(sort.Reverse(sort.IntSlice(tasks)))

	time := 0
	i := 0
	for len(tasks) > 0 && i < len(processorTime) {
		bound := min(len(tasks), 4)
		m := max(tasks[0:bound]...)
		tasks = tasks[bound:]
		time = max(time, m+processorTime[i])
		i++
	}

	return time
}

func max(nums ...int) int {
	m := 0
	for _, n := range nums {
		if m < n {
			m = n
		}
	}
	return m
}

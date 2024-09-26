package leetcode

import "sort"

type MyCalendar struct {
	events [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start int, end int) bool {
	for _, event := range c.events {
		if start < event[1] && end > event[0] {
			return false
		}
	}
	c.events = append(c.events, [2]int{start, end})
	sort.Slice(c.events, func(i, j int) bool {
		return c.events[i][0] < c.events[j][0]
	})
	return true
}

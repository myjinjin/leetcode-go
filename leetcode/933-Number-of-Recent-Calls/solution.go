package leetcode

import "container/list"

type RecentCounter struct {
	queue *list.List
}

func Constructor() RecentCounter {
	return RecentCounter{queue: list.New()}
}

func (c *RecentCounter) Ping(t int) int {
	for c.queue.Len() > 0 && t-c.queue.Front().Value.(int) > 3000 {
		c.queue.Remove(c.queue.Front())
	}
	c.queue.PushBack(t)
	return c.queue.Len()
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

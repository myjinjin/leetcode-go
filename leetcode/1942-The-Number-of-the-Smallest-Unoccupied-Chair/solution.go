package leetcode

import (
	"container/heap"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Event struct {
	time      int
	id        int
	isArrival bool
}

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	events := make([]Event, 0, 2*n)

	for i, time := range times {
		events = append(events, Event{time[0], i, true})
		events = append(events, Event{time[1], i, false})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return !events[i].isArrival && events[j].isArrival
		}
		return events[i].time < events[j].time
	})

	availableChairs := &IntHeap{}
	for i := 0; i < n; i++ {
		heap.Push(availableChairs, i)
	}

	occupiedChairs := make(map[int]int)

	for _, event := range events {
		if event.isArrival {
			chair := heap.Pop(availableChairs).(int)
			if event.id == targetFriend {
				return chair
			}
			occupiedChairs[event.id] = chair
		} else {
			chair := occupiedChairs[event.id]
			heap.Push(availableChairs, chair)
			delete(occupiedChairs, event.id)
		}
	}

	return heap.Pop(availableChairs).(int)
}

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

func minGroups(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	pq := &IntHeap{}
	heap.Init(pq)

	for _, interval := range intervals {
		left, right := interval[0], interval[1]
		if pq.Len() > 0 && (*pq)[0] < left {
			heap.Pop(pq)
		}
		heap.Push(pq, right)
	}

	return pq.Len()
}

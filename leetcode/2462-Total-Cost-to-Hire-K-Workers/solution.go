package leetcode

import (
	"container/heap"
)

func totalCost(costs []int, k int, candidates int) int64 {
	ci := []costWithIndex{}
	for i := 0; i < len(costs); i++ {
		ci = append(ci, costWithIndex{cost: costs[i], index: i})
	}

	totalCost := 0
	leftMinHeap := &MinHeap{}
	rightMinHeap := &MinHeap{}

	leftIndex := 0
	rightIndex := len(ci) - 1

	for i := 0; i < k; i++ {
		for leftIndex <= rightIndex && leftMinHeap.Len() < candidates {
			heap.Push(leftMinHeap, ci[leftIndex])
			leftIndex++
		}
		for leftIndex <= rightIndex && rightMinHeap.Len() < candidates {
			heap.Push(rightMinHeap, ci[rightIndex])
			rightIndex--
		}

		if leftMinHeap.Len() == 0 {
			totalCost += heap.Pop(rightMinHeap).(costWithIndex).cost
		} else if rightMinHeap.Len() == 0 {
			totalCost += heap.Pop(leftMinHeap).(costWithIndex).cost
		} else {
			left := heap.Pop(leftMinHeap).(costWithIndex)
			right := heap.Pop(rightMinHeap).(costWithIndex)

			if left.cost < right.cost || (left.cost == right.cost && left.index < right.index) {
				totalCost += left.cost
				heap.Push(rightMinHeap, right)
			} else {
				totalCost += right.cost
				heap.Push(leftMinHeap, left)
			}
		}
	}

	return int64(totalCost)
}

type costWithIndex struct {
	cost  int
	index int
}

type MinHeap []costWithIndex

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	if h[i].cost == h[j].cost {
		return h[i].index < h[j].index
	}
	return h[i].cost < h[j].cost
}

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(costWithIndex))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

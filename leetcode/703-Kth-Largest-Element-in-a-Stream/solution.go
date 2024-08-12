package leetcode

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type KthLargest struct {
	h *IntHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	for _, num := range nums {
		if h.Len() < k {
			heap.Push(h, num)
		} else if num > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, num)
		}
	}
	heap.Init(h)
	return KthLargest{h: h, k: k}
}

func (k *KthLargest) Add(val int) int {
	h := k.h
	if h.Len() < k.k {
		heap.Push(h, val)
	} else if val > (*h)[0] {
		heap.Pop(h)
		heap.Push(h, val)
	}
	return (*h)[0]
}

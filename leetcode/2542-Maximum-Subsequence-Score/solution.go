package leetcode

import (
	"container/heap"
	"sort"
)

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	pairs := [][]int{}
	for i := 0; i < len(nums1); i++ {
		pairs = append(pairs, []int{nums1[i], nums2[i]})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	minHeap := &MinHeap{}
	nums1Sum := 0
	for i := 0; i < k; i++ {
		heap.Push(minHeap, pairs[i])
		nums1Sum += pairs[i][0]
	}

	maxScore := nums1Sum * pairs[k-1][1]

	for i := k; i < len(pairs); i++ {
		smallestPair := heap.Pop(minHeap).([]int)
		nums1Sum += pairs[i][0] - smallestPair[0]

		heap.Push(minHeap, pairs[i])

		maxScore = max(maxScore, nums1Sum*pairs[i][1])
	}

	return int64(maxScore)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type MinHeap [][]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

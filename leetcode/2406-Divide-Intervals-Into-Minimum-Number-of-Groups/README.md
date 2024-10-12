# [2406. Divide Intervals Into Minimum Number of Groups](https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/)

## Problem

You are given a 2D integer array `intervals` where `intervals[i] = [lefti, righti]` represents the inclusive interval `[lefti, righti]`.

You have to divide the intervals into one or more **groups** such that each interval is in **exactly** one group, and no two intervals that are in the same group **intersect** each other.

Return the **minimum** number of groups you need to make.

Two intervals intersect if there is at least one common number between them. For example, the intervals `[1, 5]` and `[5, 8]` intersect.


Example 1:

```
Input: intervals = [[5,10],[6,8],[1,5],[2,3],[1,10]]
Output: 3
Explanation: We can divide the intervals into the following groups:
- Group 1: [1, 5], [6, 8].
- Group 2: [2, 3], [5, 10].
- Group 3: [1, 10].
It can be proven that it is not possible to divide the intervals into fewer than 3 groups.
```

Example 2:

```
Input: intervals = [[1,3],[5,6],[8,10],[11,13]]
Output: 1
Explanation: None of the intervals overlap, so we can put all of them in one group.
```

Constraints:

- `1 <= intervals.length <= 10^5`
- `intervals[i].length == 2`
- `1 <= lefti <= righti <= 10^6`


## Solution

```go
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
```
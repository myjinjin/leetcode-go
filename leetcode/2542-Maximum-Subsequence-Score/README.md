# [2542. Maximum Subsequence Score](https://leetcode.com/problems/maximum-subsequence-score/)

## Problem

You are given two 0-indexed integer arrays `nums1` and `nums2` of equal length `n` and a positive integer `k`. You must choose a subsequence of indices from `nums1` of length `k`.

For chosen indices `i0`, `i1`, ..., `ik - 1`, your score is defined as:

- The sum of the selected elements from `nums1` multiplied with the minimum of the selected elements from `nums2`.
- It can defined simply as: `(nums1[i0] + nums1[i1] +...+ nums1[ik - 1]) * min(nums2[i0] , nums2[i1], ... ,nums2[ik - 1])`.

Return the maximum possible score.

A subsequence of indices of an array is a set that can be derived from the set `{0, 1, ..., n-1}` by deleting some or no elements.

Example 1:

```
Input: nums1 = [1,3,3,2], nums2 = [2,1,3,4], k = 3
Output: 12
Explanation: 
The four possible subsequence scores are:
- We choose the indices 0, 1, and 2 with score = (1+3+3) * min(2,1,3) = 7.
- We choose the indices 0, 1, and 3 with score = (1+3+2) * min(2,1,4) = 6. 
- We choose the indices 0, 2, and 3 with score = (1+3+2) * min(2,3,4) = 12. 
- We choose the indices 1, 2, and 3 with score = (3+3+2) * min(1,3,4) = 8.
Therefore, we return the max score, which is 12.
```

Example 2:

```
Input: nums1 = [4,2,3,1,1], nums2 = [7,5,10,9,6], k = 1
Output: 30
Explanation: 
Choosing index 2 is optimal: nums1[2] * nums2[2] = 3 * 10 = 30 is the maximum possible score.
``` 

Constraints:

- `n == nums1.length == nums2.length`
- `1 <= n <= 10^5`
- `0 <= nums1[i], nums2[j] <= 10^5`
- `1 <= k <= n`


## Solution

```go
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
```
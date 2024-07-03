# [56. Merge Intervals](https://leetcode.com/problems/merge-intervals/)

## Problem

Given an array of `intervals` where `intervals[i] = [starti, endi]`, merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.


Example 1:

```
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
```

Example 2:

```
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
```

Constraints:

- `1 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= starti <= endi <= 10^4`


## Solution

```go
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	merged := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prevBegin := intervals[0][0]
	prevEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		currBegin := intervals[i][0]
		currEnd := intervals[i][1]
		if currBegin <= prevEnd {
			prevEnd = max(prevEnd, currEnd)
		} else {
			merged = append(merged, []int{prevBegin, prevEnd})
			prevBegin = currBegin
			prevEnd = currEnd
		}
	}

	merged = append(merged, []int{prevBegin, prevEnd})
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
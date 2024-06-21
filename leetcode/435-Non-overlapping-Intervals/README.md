# [435. Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/)

## Problem

Given an array of intervals `intervals` where `intervals[i] = [start_i, end_i]`, return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

 

Example 1:

```
Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.
```

Example 2:

```
Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.
```

Example 3:

```
Input: intervals = [[1,2],[2,3]]
Output: 0
Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
```

Constraints:

- `1 <= intervals.length <= 10^5`
- `intervals[i].length == 2`
- `-5 * 10^4 <= start_i < end_i <= 5 * 10^4`

## Solution

```go
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	result := 0
	n := len(intervals)
	prevIdx := 0
	tmpPrevInterval := 0
	currIdx := 1
	for prevIdx < n-1 && currIdx < n {
		prev := intervals[prevIdx]
		curr := intervals[currIdx]
		if curr[0] < prev[1] {
			result++
			currIdx++
			tmpPrevInterval++
		} else {
			currIdx++
			prevIdx += tmpPrevInterval + 1
			tmpPrevInterval = 0
		}
	}
	return result
}
```
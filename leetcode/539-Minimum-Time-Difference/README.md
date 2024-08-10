# [539. Minimum Time Difference](https://leetcode.com/problems/minimum-time-difference/)

## Problem

Given a list of 24-hour clock time points in **"HH:MM"** format, return the minimum **minutes** difference between any two time-points in the list.


Example 1:

```
Input: timePoints = ["23:59","00:00"]
Output: 1
```

Example 2:

```
Input: timePoints = ["00:00","23:59","00:00"]
Output: 0
``` 

Constraints:

- `2 <= timePoints.length <= 2 * 10^4`
- `timePoints[i]` is in the format **"HH:MM"**.



## Solution

```go
func findMinDifference(timePoints []string) int {
	timePointInts := make([]int, len(timePoints))
	for i, tp := range timePoints {
		timePointInts[i] = toMinutes(tp)
	}

	sort.Ints(timePointInts)

	diff := timePointInts[len(timePointInts)-1] - timePointInts[0]
	if diff > 720 {
		diff = 1440 - diff
	}

	for i := 1; i < len(timePointInts); i++ {
		newDiff := abs(timePointInts[i] - timePointInts[i-1])
		if newDiff > 720 {
			newDiff = 1440 - newDiff
		}
		diff = min(diff, newDiff)
	}
	return diff
}

func toMinutes(timePoint string) int {
	splited := strings.Split(timePoint, ":")
	hStr, mStr := splited[0], splited[1]
	h, _ := strconv.Atoi(hStr)
	m, _ := strconv.Atoi(mStr)
	return h*60 + m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
```
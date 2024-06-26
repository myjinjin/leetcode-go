# [763. Partition Labels](https://leetcode.com/problems/partition-labels/)

## Problem

You are given a string `s`. We want to partition the string into as many parts as possible so that each letter appears in at most one part.

Note that the partition is done so that after concatenating all the parts in order, the resultant string should be `s`.

Return a list of integers representing the size of these parts.


Example 1:

```
Input: s = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.
```

Example 2:

```
Input: s = "eccbbbbdec"
Output: [10]
``` 

Constraints:

- `1 <= s.length <= 500`
- `s` consists of lowercase English letters.

## Solution

```go
func partitionLabels(s string) []int {
	lastIndex := make(map[rune]int)
	for i, ch := range s {
		lastIndex[ch] = i
	}

	result := []int{}
	start, end := 0, 0
	for i, ch := range s {
		if lastIndex[ch] > end {
			end = lastIndex[ch]
		}
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}

	return result
}
```
# [2418. Sort the People](https://leetcode.com/problems/sort-the-people/)

## Problem

You are given an array of strings `names`, and an array `heights` that consists of distinct positive integers. Both arrays are of length `n`.

For each index `i`, `names[i]` and `heights[i]` denote the name and height of the `ith` person.

Return `names` sorted in **descending** order by the people's heights.


Example 1:

```
Input: names = ["Mary","John","Emma"], heights = [180,165,170]
Output: ["Mary","Emma","John"]
Explanation: Mary is the tallest, followed by Emma and John.
```

Example 2:

```
Input: names = ["Alice","Bob","Bob"], heights = [155,185,150]
Output: ["Bob","Alice","Bob"]
Explanation: The first Bob is the tallest, followed by Alice and the second Bob.
``` 

Constraints:

- `n == names.length == heights.length`
- `1 <= n <= 10^3`
- `1 <= names[i].length <= 20`
- `1 <= heights[i] <= 10^5`
- `names[i]` consists of lower and upper case English letters.
- All the values of `heights` are distinct.

## Solution

```go
func sortPeople(names []string, heights []int) []string {
	orderedHeight := make([][2]int, len(heights))
	for i := 0; i < len(heights); i++ {
		orderedHeight[i] = [2]int{i, heights[i]}
	}

	sort.Slice(orderedHeight, func(i, j int) bool {
		return orderedHeight[i][1] > orderedHeight[j][1]
	})

	result := make([]string, len(names))
	for i := 0; i < len(result); i++ {
		result[i] = names[orderedHeight[i][0]]
	}
	return result
}
```
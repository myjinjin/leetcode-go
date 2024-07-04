# [744. Find Smallest Letter Greater Than Target](https://leetcode.com/problems/find-smallest-letter-greater-than-target/)

## Problem

You are given an array of characters `letters` that is sorted in non-decreasing order, and a character `target`. There are at least two different characters in `letters`.

Return the smallest character in `letters` that is lexicographically greater than `target`. If such a character does not exist, return the first character in `letters`.

Example 1:

```
Input: letters = ["c","f","j"], target = "a"
Output: "c"
Explanation: The smallest character that is lexicographically greater than 'a' in letters is 'c'.
```

Example 2:

```
Input: letters = ["c","f","j"], target = "c"
Output: "f"
Explanation: The smallest character that is lexicographically greater than 'c' in letters is 'f'.
```

Example 3:

```
Input: letters = ["x","x","y","y"], target = "z"
Output: "x"
Explanation: There are no characters in letters that is lexicographically greater than 'z' so we return letters[0].
``` 

Constraints:

- `2 <= letters.length <= 10^4`
- `letters[i]` is a lowercase English letter.
- `letters` is sorted in non-decreasing order.
- `letters` contains at least two different characters.
- `target` is a lowercase English letter.

## Solution

```go
func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	left := 0
	right := n - 1

	for left <= right {
		mid := (left + right) / 2
		curr := letters[mid]
		if curr > target {
			if mid-1 >= 0 && letters[mid-1] <= target {
				return curr
			}
			right = mid - 1
		} else if curr < target {
			left = mid + 1
		} else {
			if mid+1 < n && letters[mid+1] > target {
				return letters[mid+1]
			}
			left = mid + 1
		}
	}
	return letters[0]
}
```
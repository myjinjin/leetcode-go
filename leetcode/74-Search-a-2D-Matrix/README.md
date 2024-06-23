# [74. Search a 2D Matrix](https://leetcode.com/problems/search-a-2d-matrix/)

## Problem

You are given an `m x n` integer matrix `matrix` with the following two properties:

- Each row is sorted in non-decreasing order.
- The first integer of each row is greater than the last integer of the previous row.

Given an integer `target`, return `true` if `target` is in `matrix` or `false` otherwise.

You must write a solution in `O(log(m * n))` time complexity.


Example 1:

![alt text](image.png)

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
```

Example 2:

![alt text](image-1.png)

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
```

Constraints:

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 100`
- `-10^4 <= matrix[i][j], target <= 10^4`

## Solution

```go
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	left, right := 0, m*n-1
	for left < right {
		mid := (left + right) / 2

		row := mid / n
		col := mid % n
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	row := left / n
	col := left % n

	return matrix[row][col] == target
}
```
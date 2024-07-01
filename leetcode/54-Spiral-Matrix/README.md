# [54. Spiral Matrix](https://leetcode.com/problems/spiral-matrix/)

## Problem

Given an `m x n` matrix, return all elements of the `matrix` in spiral order.

Example 1:

![alt text](image.png)

```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
```

Example 2:

![alt text](image-1.png)

```
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
``` 

Constraints:

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 10`
- `-100 <= matrix[i][j] <= 100`

## Solution

```go
func spiralOrder(matrix [][]int) []int {
	result := []int{}

	rowBegin := 0
	colBegin := 0
	rowEnd := len(matrix) - 1
	colEnd := len(matrix[0]) - 1

	for rowBegin <= rowEnd && colBegin <= colEnd {
		for j := colBegin; j <= colEnd; j++ {
			result = append(result, matrix[rowBegin][j])
		}
		rowBegin++

		for i := rowBegin; i <= rowEnd; i++ {
			result = append(result, matrix[i][colEnd])
		}
		colEnd--

		if rowBegin <= rowEnd {
			for j := colEnd; j >= colBegin; j-- {
				result = append(result, matrix[rowEnd][j])
			}
		}
		rowEnd--

		if colBegin <= colEnd {
			for i := rowEnd; i >= rowBegin; i-- {
				result = append(result, matrix[i][colBegin])
			}
		}
		colBegin++
	}

	return result
}
```
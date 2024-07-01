# [73. Set Matrix Zeroes](https://leetcode.com/problems/set-matrix-zeroes/)

## Problem

Given an `m x n` integer matrix `matrix`, if an element is `0`, set its entire row and column to `0`'s.

You must do it **in place**.


Example 1:

![alt text](image.png)

```
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]
```

Example 2:

![alt text](image-1.png)

```
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
``` 

Constraints:

- `m == matrix.length`
- `n == matrix[0].length`
- `1 <= m, n <= 200`
- `-2^31 <= matrix[i][j] <= 2^31 - 1`
 

Follow up:

A straightforward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?


## Solution

```go
func setZeroes(matrix [][]int) {
	zeroIndex := make([][2]int, 0)

	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeroIndex = append(zeroIndex, [2]int{i, j})
			}
		}
	}

	for _, indexes := range zeroIndex {
		r := indexes[0]
		c := indexes[1]

		for j := 0; j < n; j++ {
			matrix[r][j] = 0
		}
		for i := 0; i < m; i++ {
			matrix[i][c] = 0
		}
	}
}
```
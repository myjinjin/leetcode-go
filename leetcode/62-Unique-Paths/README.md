# [62. Unique Paths](https://leetcode.com/problems/unique-paths/)

## Problem

There is a robot on an `m x n` grid. The robot is initially located at the top-left corner (i.e., `grid[0][0]`). The robot tries to move to the bottom-right corner (i.e., `grid[m - 1][n - 1]`). The robot can only move either down or right at any point in time.

Given the two integers `m` and `n`, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to `2 * 10^9`.

 

Example 1:

![alt text](image.png)

```
Input: m = 3, n = 7
Output: 28
```

Example 2:

```
Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down
```

Constraints:

- `1 <= m, n <= 100`

## Solution

```go
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[row][col] = dp[row-1][col] + dp[row][col-1]
		}
	}

	return dp[m-1][n-1]
}
```

```go
import "math/big"

func uniquePaths(m int, n int) int {
	if m+n-1 <= 1 {
		return 1
	}
	factorialDp := make([]*big.Int, m+n-1)
	factorialDp[0] = big.NewInt(1)
	factorialDp[1] = big.NewInt(1)
	for i := 2; i < m+n-1; i++ {
		factorialDp[i] = big.NewInt(0).Mul(factorialDp[i-1], big.NewInt(int64(i)))
	}

	result := big.NewInt(0).Div(factorialDp[m+n-2], big.NewInt(0).Mul(factorialDp[m-1], factorialDp[n-1]))
	return int(result.Int64())
}
```
# [386. Lexicographical Numbers](https://leetcode.com/problems/lexicographical-numbers/)

## Problem

Given an integer `n`, return all the numbers in the range `[1, n]` sorted in lexicographical order.

You must write an algorithm that runs in `O(n)` time and uses `O(1)` extra space.
 

Example 1:

```
Input: n = 13
Output: [1,10,11,12,13,2,3,4,5,6,7,8,9]
```

Example 2:

```
Input: n = 2
Output: [1,2]
```

Constraints:

- `1 <= n <= 5 * 10^4`

## Solution

```go
func lexicalOrder(n int) []int {
    result := make([]int, n)
    curr := 1
    for i := 0; i < n; i++ {
        result[i] = curr
        if curr*10 <= n {
            curr *= 10
        } else {
            for curr%10 == 9 || curr+1 > n {
                curr /= 10
            }
            curr++
        }
    }
    return result
}
```
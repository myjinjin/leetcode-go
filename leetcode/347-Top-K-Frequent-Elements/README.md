# [347. Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/)

## Problem

Given an integer array `nums` and an integer `k`, return the `k` most frequent elements. You may return the answer in any order.

 

Example 1:

```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

Example 2:

```
Input: nums = [1], k = 1
Output: [1]
``` 

Constraints:

- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`
- `k` is in the range `[1, the number of unique elements in the array]`.
- It is guaranteed that the answer is unique.


## Solution

```go
func topKFrequent(nums []int, k int) []int {
	result := make([]int, k)
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	numCount := make([][2]int, 0, len(numMap))
	for num, count := range numMap {
		numCount = append(numCount, [2]int{num, count})
	}

	sort.Slice(numCount, func(i, j int) bool {
		return numCount[i][1] > numCount[j][1]
	})

	for i := 0; i < k; i++ {
		result[i] = numCount[i][0]
	}
	return result
}
```
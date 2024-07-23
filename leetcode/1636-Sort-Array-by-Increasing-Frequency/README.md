# [1636. Sort Array by Increasing Frequency](https://leetcode.com/problems/sort-array-by-increasing-frequency/)

## Problem

Given an array of integers `nums`, sort the array in **increasing** order based on the frequency of the values. If multiple values have the same frequency, sort them in **decreasing** order.

Return the sorted array.

 

Example 1:

```
Input: nums = [1,1,2,2,2,3]
Output: [3,1,1,2,2,2]
Explanation: '3' has a frequency of 1, '1' has a frequency of 2, and '2' has a frequency of 3.
```

Example 2:

```
Input: nums = [2,3,1,3,2]
Output: [1,3,3,2,2]
Explanation: '2' and '3' both have a frequency of 2, so they are sorted in decreasing order.
```

Example 3:

```
Input: nums = [-1,1,-6,4,5,-6,1,4,1]
Output: [5,-1,4,4,-6,-6,1,1,1]
``` 

Constraints:

- `1 <= nums.length <= 100`
- `-100 <= nums[i] <= 100`


## Solution

```go
func frequencySort(nums []int) []int {
	numCount := make(map[int]int)
	for i := range nums {
		numCount[nums[i]]++
	}

	numsWithCount := make([][2]int, 0, len(numCount))
	for k, v := range numCount {
		numsWithCount = append(numsWithCount, [2]int{k, v})
	}

	sort.Slice(numsWithCount, func(i, j int) bool {
		if numsWithCount[i][1] < numsWithCount[j][1] {
			return true
		}
		if numsWithCount[i][1] > numsWithCount[j][1] {
			return false
		}
		return numsWithCount[i][0] > numsWithCount[j][0]
	})

	result := make([]int, 0, len(nums))
	for _, v := range numsWithCount {
		value := v[0]
		count := v[1]
		for i := 0; i < count; i++ {
			result = append(result, value)
		}
	}
	return result
}
```
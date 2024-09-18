# [179. Largest Number](https://leetcode.com/problems/largest-number/)

## Problem

Given a list of non-negative integers `nums`, arrange them such that they form the largest number and return it.

Since the result may be very large, so you need to return a string instead of an integer.

Example 1:

```
Input: nums = [10,2]
Output: "210"
```

Example 2:

```
Input: nums = [3,30,34,5,9]
Output: "9534330"
``` 

Constraints:

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 10^9`


## Solution

```go
func largestNumber(nums []int) string {
	strs := make([]string, len(nums))

	allZero := true
	for i, num := range nums {
		if num != 0 {
			allZero = false
		}
		strs[i] = strconv.Itoa(num)
	}

	if allZero {
		return "0"
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	return strings.Join(strs, "")
}
```
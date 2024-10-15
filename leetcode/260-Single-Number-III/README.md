# [260. Single Number III](https://leetcode.com/problems/single-number-iii/)

## Problem

Given an integer array nums, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once. You can return the answer in any order.

You must write an algorithm that runs in linear runtime complexity and uses only constant extra space.

 

Example 1:

Input: nums = [1,2,1,3,2,5]
Output: [3,5]
Explanation:  [5, 3] is also a valid answer.
Example 2:

Input: nums = [-1,0]
Output: [-1,0]
Example 3:

Input: nums = [0,1]
Output: [1,0]
 

Constraints:

- `2 <= nums.length <= 3 * 10^4`
- `-2^31 <= nums[i] <= 2^31 - 1`
Each integer in nums will appear twice, only two integers will appear once.

## Solution

```go
func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	rightMost := 1
	for xor&rightMost == 0 {
		rightMost <<= 1
	}

	num1, num2 := 0, 0
	for _, num := range nums {
		if num&rightMost != 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}

	return []int{num1, num2}
}
```
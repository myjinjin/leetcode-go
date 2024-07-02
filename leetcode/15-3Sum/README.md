# [15. 3Sum](https://leetcode.com/problems/3sum/)

## Problem

Given an integer array nums, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

Notice that the solution set must not contain duplicate triplets.

 

Example 1:

```
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation: 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
```

Example 2:

```
Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
```

Example 3:

```
Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
``` 

Constraints:

- `3 <= nums.length <= 3000`
- `-10^5 <= nums[i] <= 10^5`

## Solution

```go
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for first := 0; first < len(nums)-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second := first + 1
		third := len(nums) - 1

		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			if sum == 0 {
				result = append(result, []int{nums[first], nums[second], nums[third]})
				second++
				third--
				for second < third && nums[second] == nums[second-1] {
					second++
				}
				for second < third && nums[third] == nums[third+1] {
					third--
				}
			}
			if sum > 0 {
				third--
			} else if sum < 0 {
				second++
			}
		}
	}
	return result
}
```
package leetcode

import "sort"

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

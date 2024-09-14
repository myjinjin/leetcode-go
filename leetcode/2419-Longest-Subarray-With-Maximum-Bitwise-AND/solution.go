package leetcode

func longestSubarray(nums []int) int {
	longest := 0
	length := 0
	maxNum := 0

	for _, num := range nums {
		if maxNum < num {
			maxNum = num
			length = 1
			longest = 1
		} else if maxNum == num {
			length++
			longest = max(longest, length)
		} else {
			length = 0
		}
	}

	return longest
}

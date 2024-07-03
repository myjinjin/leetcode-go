package leetcode

func findDuplicate(nums []int) int {
	var result int
	numMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := numMap[nums[i]]; ok {
			result = nums[i]
			break
		}
		numMap[nums[i]] = true
	}
	return result
}

package leetcode

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	sumIndex := make(map[int]int)
	sumIndex[0] = -1
	sum := 0

	for i := 0; i < n; i++ {
		sum += nums[i]
		sum %= k
		if prev, ok := sumIndex[sum]; ok {
			if i-prev >= 2 {
				return true
			}
		} else {
			sumIndex[sum] = i
		}
	}

	return false
}

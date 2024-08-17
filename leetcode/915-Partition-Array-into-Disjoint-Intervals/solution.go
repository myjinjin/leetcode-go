package leetcode

func partitionDisjoint(nums []int) int {
	n := len(nums)
	leftMax := nums[0]
	maxSoFar := nums[0]
	partitionIndex := 0

	for i := 1; i < n; i++ {
		if nums[i] < leftMax {
			partitionIndex = i
			leftMax = maxSoFar
		} else if nums[i] > maxSoFar {
			maxSoFar = nums[i]
		}

	}

	return partitionIndex + 1
}

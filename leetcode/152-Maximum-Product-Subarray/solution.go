package leetcode

func maxProduct(nums []int) int {
	n := len(nums)

	maxProduct := nums[0]
	minProduct := nums[0]
	result := maxProduct

	for i := 1; i < n; i++ {
		curr := nums[i]
		if curr < 0 {
			minProduct, maxProduct = maxProduct, minProduct
		}
		minProduct = min(curr, curr*minProduct)
		maxProduct = max(curr, curr*maxProduct)

		result = max(result, maxProduct)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

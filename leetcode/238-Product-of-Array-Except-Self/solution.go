package leetcode

func productExceptSelf(nums []int) []int {
	n := len(nums)

	product := make([]int, n)
	product[0] = 1
	for i := 1; i < n; i++ {
		product[i] = product[i-1] * nums[i-1]
	}

	right := 1
	for i := n - 2; i >= 0; i-- {
		right *= nums[i+1]
		product[i] = product[i] * right
	}

	return product
}

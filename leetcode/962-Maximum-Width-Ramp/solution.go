package leetcode

func maxWidthRamp(nums []int) int {
	n := len(nums)
	stack := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if len(stack) == 0 || nums[stack[len(stack)-1]] > nums[i] {
			stack = append(stack, i)
		}
	}

	width := 0

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			width = max(width, i-stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			break
		}
	}

	return width
}

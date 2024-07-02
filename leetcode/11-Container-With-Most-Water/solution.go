package leetcode

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	area := 0
	for left < right {
		leftHeight := height[left]
		rightHeight := height[right]
		minHeight := min(leftHeight, rightHeight)
		width := right - left
		area = max(area, minHeight*width)
		if leftHeight < rightHeight {
			left++
		} else {
			right--
		}
	}
	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

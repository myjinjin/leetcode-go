package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHelper(node *TreeNode, minVal, maxVal int) bool {
	if node == nil {
		return true
	}

	if node.Val >= maxVal || node.Val <= minVal {
		return false
	}
	return isValidBSTHelper(node.Left, minVal, node.Val) && isValidBSTHelper(node.Right, node.Val, maxVal)
}

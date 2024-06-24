package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	height(root, &diameter)
	return diameter
}

func height(node *TreeNode, diameter *int) int {
	if node == nil {
		return 0
	}

	leftHeight := height(node.Left, diameter)
	rightHeight := height(node.Right, diameter)

	*diameter = max(*diameter, leftHeight+rightHeight)
	return max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

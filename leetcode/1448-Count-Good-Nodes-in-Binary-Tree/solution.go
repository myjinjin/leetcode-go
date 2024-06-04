package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return goodNodeHelper(root, -10000)
}

func goodNodeHelper(node *TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val >= maxVal {
		count++
		maxVal = node.Val
	}

	count += goodNodeHelper(node.Left, maxVal)
	count += goodNodeHelper(node.Right, maxVal)

	return count
}

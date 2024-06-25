package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	var rightToRootToLeft func(node *TreeNode)
	rightToRootToLeft = func(node *TreeNode) {
		if node == nil {
			return
		}
		rightToRootToLeft(node.Right)
		sum += node.Val
		node.Val = sum
		rightToRootToLeft(node.Left)
	}

	rightToRootToLeft(root)
	return root
}

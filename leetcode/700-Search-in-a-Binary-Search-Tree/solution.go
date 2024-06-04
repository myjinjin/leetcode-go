package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	var traverse func(node *TreeNode, val int) *TreeNode
	traverse = func(node *TreeNode, val int) *TreeNode {
		if node == nil {
			return node
		}
		if node.Val == val {
			return node
		}
		if node.Val > val {
			return traverse(node.Left, val)
		}
		if node.Val < val {
			return traverse(node.Right, val)
		}
		return nil
	}
	return traverse(root, val)
}

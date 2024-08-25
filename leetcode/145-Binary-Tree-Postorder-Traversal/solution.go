package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	path := []int{}
	helper(root, &path)
	return path
}

func helper(node *TreeNode, path *[]int) {
	if node == nil {
		return
	}

	helper(node.Left, path)
	helper(node.Right, path)
	*path = append(*path, node.Val)
}

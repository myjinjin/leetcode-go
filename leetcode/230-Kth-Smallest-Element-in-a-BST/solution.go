package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	result := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			inorder(node.Left)
		}
		result = append(result, node.Val)
		if node.Right != nil {
			inorder(node.Right)
		}
	}
	inorder(root)
	return result[k-1]
}

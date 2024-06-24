package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	stack := []*TreeNode{}
	stack = preorder(root, stack)

	for i := len(stack) - 1; i >= 0; i-- {
		stack[i].Left = nil
		if i < len(stack)-1 {
			stack[i].Right = stack[i+1]
		}
	}
}

func preorder(root *TreeNode, stack []*TreeNode) []*TreeNode {
	if root == nil {
		return stack
	}

	stack = append(stack, root)

	if root.Left != nil {
		stack = preorder(root.Left, stack)
	}
	if root.Right != nil {
		stack = preorder(root.Right, stack)
	}

	return stack
}

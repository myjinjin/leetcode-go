package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	inorderRootIndex := 0
	for inorderRootIndex < len(inorder) && inorder[inorderRootIndex] != rootVal {
		inorderRootIndex++
	}

	leftInorder := inorder[:inorderRootIndex]
	rightInorder := inorder[inorderRootIndex+1:]

	leftPreorder := preorder[1 : inorderRootIndex+1]
	rightPreorder := preorder[inorderRootIndex+1:]

	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)

	return root
}

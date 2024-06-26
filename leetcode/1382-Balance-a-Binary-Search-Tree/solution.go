package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	var sortedArr []*TreeNode

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		sortedArr = append(sortedArr, node)
		inorder(node.Right)
	}

	var sortedArrToBST func(start, end int) *TreeNode
	sortedArrToBST = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) / 2
		root := sortedArr[mid]
		root.Left = sortedArrToBST(start, mid-1)
		root.Right = sortedArrToBST(mid+1, end)
		return root
	}

	inorder(root)
	return sortedArrToBST(0, len(sortedArr)-1)
}

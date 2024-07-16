package leetcode

import "bytes"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	lowestCommonAncestor := findLowestCommonAncestor(root, startValue, destValue)

	pathToStart := make([]byte, 0)
	pathToDest := make([]byte, 0)

	findPath(lowestCommonAncestor, &pathToStart, startValue)
	findPath(lowestCommonAncestor, &pathToDest, destValue)

	upwardPath := bytes.Repeat([]byte{'U'}, len(pathToStart))
	return string(upwardPath) + string(pathToDest)
}

func findLowestCommonAncestor(root *TreeNode, startValue, destValue int) *TreeNode {
	if root == nil || root.Val == startValue || root.Val == destValue {
		return root
	}

	left := findLowestCommonAncestor(root.Left, startValue, destValue)
	right := findLowestCommonAncestor(root.Right, startValue, destValue)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}

func findPath(node *TreeNode, path *[]byte, target int) bool {
	if node == nil {
		return false
	}

	if node.Val == target {
		return true
	}

	*path = append(*path, 'L')
	if findPath(node.Left, path, target) {
		return true
	}

	*path = (*path)[:len(*path)-1]

	*path = append(*path, 'R')
	if findPath(node.Right, path, target) {
		return true
	}

	*path = (*path)[:len(*path)-1]

	return false
}

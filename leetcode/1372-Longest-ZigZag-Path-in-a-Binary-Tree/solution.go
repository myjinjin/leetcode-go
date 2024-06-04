package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	leftCount := dfs(root.Left, true, 0)
	rightCount := dfs(root.Right, false, 0)
	return max(leftCount, rightCount)
}

func dfs(node *TreeNode, isLeft bool, count int) int {
	if node == nil {
		return count
	}

	if isLeft {
		leftCount := dfs(node.Left, true, 0)
		rightCount := dfs(node.Right, false, count+1)
		return max(leftCount, rightCount)
	}
	leftCount := dfs(node.Left, true, count+1)
	rightCount := dfs(node.Right, false, 0)
	return max(leftCount, rightCount)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

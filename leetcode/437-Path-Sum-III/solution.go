package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return dfs(root, 0, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func dfs(node *TreeNode, currSum int, targetSum int) int {
	if node == nil {
		return 0
	}
	count := 0
	currSum += node.Val
	if currSum == targetSum {
		count++
	}

	count += dfs(node.Left, currSum, targetSum)
	count += dfs(node.Right, currSum, targetSum)
	return count
}

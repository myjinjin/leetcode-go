package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	result := []*TreeNode{}
	deleteSet := make(map[int]bool)
	for _, val := range toDelete {
		deleteSet[val] = true
	}

	var dfs func(node *TreeNode, isRoot bool) *TreeNode
	dfs = func(node *TreeNode, isRoot bool) *TreeNode {
		if node == nil {
			return nil
		}

		toBeDeleted := deleteSet[node.Val]

		if isRoot && !toBeDeleted {
			result = append(result, node)
		}

		if toBeDeleted {
			node.Left = dfs(node.Left, true)
			node.Right = dfs(node.Right, true)
			return nil
		}

		node.Left = dfs(node.Left, false)
		node.Right = dfs(node.Right, false)
		return node
	}

	dfs(root, true)

	return result
}

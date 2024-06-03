package leetcode

import "reflect"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs func(root *TreeNode, result *[]int)
	dfs = func(root *TreeNode, result *[]int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			*result = append(*result, root.Val)
		}

		dfs(root.Left, result)
		dfs(root.Right, result)
	}

	oneResult := make([]int, 0)
	twoResult := make([]int, 0)
	dfs(root1, &oneResult)
	dfs(root2, &twoResult)

	return reflect.DeepEqual(oneResult, twoResult)
}

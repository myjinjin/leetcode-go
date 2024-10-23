package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	level := make(map[int][][]*TreeNode)

	var dfs func(node, parent *TreeNode, depth int)
	dfs = func(node, parent *TreeNode, depth int) {
		if node == nil {
			return
		}

		level[depth] = append(level[depth], []*TreeNode{node, parent})

		dfs(node.Left, node, depth+1)
		dfs(node.Right, node, depth+1)
	}

	dfs(root, nil, 0)

	for depth, nodes := range level {
		parentSum := make(map[*TreeNode]int)
		levelSum := 0
		for _, p := range nodes {
			node, parent := p[0], p[1]
			parentSum[parent] += node.Val
			levelSum += node.Val
		}

		for _, p := range nodes {
			node, parent := p[0], p[1]
			if depth == 0 {
				node.Val = 0
			} else {
				node.Val = levelSum - parentSum[parent]
			}
		}
	}

	return root
}

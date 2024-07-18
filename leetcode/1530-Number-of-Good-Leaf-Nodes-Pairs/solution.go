package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	goodPairCount := 0

	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{}
		}

		if node.Left == nil && node.Right == nil {
			return []int{0}
		}

		leftLeafDistances := dfs(node.Left)
		rightLeafDistances := dfs(node.Right)

		for _, leftDist := range leftLeafDistances {
			for _, rightDist := range rightLeafDistances {
				if leftDist+rightDist+2 <= distance {
					goodPairCount++
				}
			}
		}

		currNodeDistances := []int{}
		for _, dist := range leftLeafDistances {
			newDist := dist + 1
			if newDist < distance {
				currNodeDistances = append(currNodeDistances, newDist)
			}
		}

		for _, dist := range rightLeafDistances {
			newDist := dist + 1
			if newDist < distance {
				currNodeDistances = append(currNodeDistances, newDist)
			}
		}

		return currNodeDistances
	}

	dfs(root)
	return goodPairCount
}

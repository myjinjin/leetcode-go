package leetcode

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	queue := []*TreeNode{root}
	var levelSums []int64

	for len(queue) > 0 {
		levelSize := len(queue)
		var levelSum int64

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += int64(node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		levelSums = append(levelSums, levelSum)
	}

	if k > len(levelSums) {
		return -1
	}

	sort.Slice(levelSums, func(i, j int) bool {
		return levelSums[i] > levelSums[j]
	})

	return levelSums[k-1]
}

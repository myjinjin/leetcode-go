package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	queue := []*TreeNode{root}

	maxSum := -math.MaxInt
	maxSumLevel := 0
	level := 0
	for len(queue) > 0 {
		level++
		size := len(queue)
		currSum := 0
		for i := 0; i < size; i++ {
			curr := queue[0]
			currSum += curr.Val
			queue = queue[1:]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		if currSum > maxSum {
			maxSum = currSum
			maxSumLevel = level
		}
	}

	return maxSumLevel
}

package leetcode

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil && root.Right == nil || root.Right != nil && root.Left == nil {
		return false
	}

	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		size := len(queue)
		temp := []int{}

		for size > 0 {
			curr := queue[0]
			queue = queue[1:]
			if curr == nil {
				temp = append(temp, math.MaxInt32)
			} else {
				temp = append(temp, curr.Val)
				queue = append(queue, curr.Left, curr.Right)
			}
			size--
		}
		for i := 0; i < len(temp)/2; i++ {
			if temp[i] != temp[len(temp)-1-i] {
				return false
			}
		}
	}
	return true
}

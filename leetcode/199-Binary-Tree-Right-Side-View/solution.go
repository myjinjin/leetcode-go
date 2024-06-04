package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	result := []int{}

	for len(queue) != 0 {
		size := len(queue)

		var prev *TreeNode

		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}

			prev = current
		}

		result = append(result, prev.Val)
	}

	return result
}

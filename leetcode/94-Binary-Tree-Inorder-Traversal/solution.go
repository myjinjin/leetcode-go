package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- Recursive ---
// func inorderTraversal(root *TreeNode) []int {
// 	result := []int{}
// 	var traverse func(node *TreeNode)
// 	traverse = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}

// 		traverse(node.Left)
// 		result = append(result, node.Val)
// 		traverse(node.Right)
// 	}

// 	traverse(root)

// 	return result
// }

// --- Iterative ---
func inorderTraversal(root *TreeNode) []int {
	result := []int{}

	queue := []*TreeNode{}
	curr := root
	for curr != nil || len(queue) > 0 {
		for curr != nil {
			queue = append(queue, curr)
			curr = curr.Left
		}
		curr = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		result = append(result, curr.Val)
		curr = curr.Right
	}

	return result
}

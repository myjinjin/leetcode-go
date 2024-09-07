package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	n := len(nums)

	var dfs func(node *TreeNode, index int) bool
	dfs = func(node *TreeNode, index int) bool {
		if index == n {
			return true
		}
		if node == nil {
			return false
		}
		if node.Val == nums[index] {
			return dfs(node.Left, index+1) || dfs(node.Right, index+1)
		}
		return false
	}

	var traverse func(node *TreeNode) bool
	traverse = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if dfs(node, 0) {
			return true
		}
		return traverse(node.Left) || traverse(node.Right)
	}

	return traverse(root)
}

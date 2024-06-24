package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	n := len(nums)
	rootIndex := n / 2
	rootVal := nums[rootIndex]
	root := &TreeNode{Val: rootVal}
	left := sortedArrayToBST(nums[:rootIndex])
	right := sortedArrayToBST(nums[rootIndex+1:])
	root.Left = left
	root.Right = right
	return root
}

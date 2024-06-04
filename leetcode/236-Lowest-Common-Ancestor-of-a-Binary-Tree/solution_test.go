package leetcode

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	testCases := []struct {
		root     *TreeNode
		p        int
		q        int
		expected int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 6},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 8},
				},
			},
			p:        5,
			q:        1,
			expected: 3,
		},
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 6},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 8},
				},
			},
			p:        5,
			q:        4,
			expected: 5,
		},
		{
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			p:        1,
			q:        2,
			expected: 1,
		},
	}

	for i, tc := range testCases {
		pNode := findNode(tc.root, tc.p)
		qNode := findNode(tc.root, tc.q)

		result := lowestCommonAncestor(tc.root, pNode, qNode)
		if result.Val != tc.expected {
			t.Errorf("Test case %d failed. Expected: %v, Got: %v", i+1, tc.expected, result.Val)
		}
	}
}

func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	left := findNode(root.Left, val)
	if left != nil {
		return left
	}

	right := findNode(root.Right, val)
	if right != nil {
		return right
	}

	return nil
}

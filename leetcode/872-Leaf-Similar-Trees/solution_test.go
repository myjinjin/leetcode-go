package leetcode

import "testing"

func TestLeafSimilar(t *testing.T) {
	testCases := []struct {
		root1    *TreeNode
		root2    *TreeNode
		expected bool
	}{
		{
			root1: &TreeNode{
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
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 8},
				},
			},
			root2: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
				Right: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 9},
						Right: &TreeNode{Val: 8},
					},
				},
			},
			expected: true,
		},
		{
			root1: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			root2: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 2},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		result := leafSimilar(tc.root1, tc.root2)
		if result != tc.expected {
			t.Errorf("Input: root1 = %v, root2 = %v, Expected: %v, but got %v", tc.root1, tc.root2, tc.expected, result)
		}
	}
}

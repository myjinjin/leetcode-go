package leetcode

import (
	"testing"
)

func TestPathSum(t *testing.T) {
	tests := []struct {
		input     *TreeNode
		targetSum int
		expected  int
	}{
		{
			input: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: -2,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: -3,
					Right: &TreeNode{
						Val: 11,
					},
				},
			},
			targetSum: 8,
			expected:  3,
		},
		{
			input: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			targetSum: 22,
			expected:  3,
		},
	}

	for _, tc := range tests {
		result := pathSum(tc.input, tc.targetSum)
		if result != tc.expected {
			t.Errorf("Expected %d, but got %d", tc.expected, result)
		}
	}
}

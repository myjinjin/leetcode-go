package leetcode

import (
	"fmt"
	"testing"
)

func TestGoodNodes(t *testing.T) {
	tests := []struct {
		input    *TreeNode
		expected int
	}{
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			expected: 4,
		},
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			expected: 3,
		},
		{
			input: &TreeNode{
				Val: 1,
			},
			expected: 1,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("Test Case %d ", i+1), func(t *testing.T) {
			result := goodNodes(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
